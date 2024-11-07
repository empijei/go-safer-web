// Package auth provides a simple implementation of an authorization mechanism.
package auth

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/empijei/go-safer-web/runtimeutils"
)

var (
	reportOnly = false
	logger     = func(s string) { log.Println(s) }
	logPrefix  = "auth: enforce: "
)

func loggerF(f string, args ...any) {
	logger(fmt.Sprintf(f, args...))
}

// UnsafelySetReportOnly completely disables auth.
// Auth errors will only be logged and no errors will ever be returned by any auth function.
// Only use during transitions to use the auth package.
func UnsafelySetReportOnly(loggerFunc func(string)) {
	if !strings.HasPrefix(runtimeutils.GetCallerName(), "main.") {
		panic("auth can only be set to ReportOnly by the main package")
	}
	logPrefix = "auth: report only: "
	reportOnly = true
	SetLogger(loggerFunc)
}
func SetLogger(loggerFunc func(string)) {
	logger = loggerFunc
}

// grantedPrivilegesKey is a context key that maps to a []string
type grantedPrivilegesKey struct{}

// Grant returns a new context with the given privileges granted to it.
// It must be called at most once per context or it will panic.
func Grant(c context.Context, privileges ...string) context.Context {
	if c.Value(grantedPrivilegesKey{}) != nil {
		panic("auth.Grant called multiple times")
	}
	return context.WithValue(c, grantedPrivilegesKey{}, privileges)
}

// checkedPrivilegesKey is a context key that maps to a []string
type checkedPrivilegesKey struct{}

// Check checks that all the given privileges have been granted for the given context.
// It returns a new context that contains information on the checked privileges.
//
// Calling Check with no privileges is valid and explicitly says that the current
// context requires no authorization to be accessed.
func Check(c context.Context, privileges ...string) (context.Context, error) {
	if privileges == nil {
		privileges = []string{}
	}

	if len(privileges) == 0 {
		// Open endpoint, return early.
		return success(c, privileges)
	}

	granted, ok := c.Value(grantedPrivilegesKey{}).([]string)
	if !ok {
		err := fmt.Errorf("%s: check failed: no privileges granted", logPrefix)
		if reportOnly {
			logger(err.Error())
			return success(c, privileges)
		}
		return fail(c, err)
	}
	for _, p := range privileges {
		if !slices.Contains(granted, p) {
			err := fmt.Errorf("%s: check failed: privilege %q requested but not granted", logPrefix, p)
			if reportOnly {
				logger(err.Error())
				continue
			}
			return fail(c, err)
		}
	}
	return success(c, privileges)
}

func success(c context.Context, privileges []string) (context.Context, error) {
	return context.WithValue(c, checkedPrivilegesKey{}, privileges), nil
}
func fail(c context.Context, err error) (context.Context, error) {
	return c, err
}

// Must checks whether a successful auth.Check has happened for the given context and
// each of the given privileges.
func Must(c context.Context, privileges ...string) error {
	checked, ok := c.Value(checkedPrivilegesKey{}).([]string)
	if !ok {
		err := fmt.Errorf("%sauth.Must: failed: an auth check was required and was not executed (caller=%q)", logPrefix, runtimeutils.GetCallerName())
		if reportOnly {
			logger(err.Error())
			return nil
		}
		return err
	}
	for _, p := range privileges {
		if !slices.Contains(checked, p) {
			err := fmt.Errorf("%sauth.Must: failed: an auth check for privilege=%q was required and was not executed (caller=%q)", logPrefix, p, runtimeutils.GetCallerName())
			if reportOnly {
				logger(err.Error())
				continue
			}
			return err
		}
	}
	return nil
}
