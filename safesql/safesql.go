// Package safesql implements a safe version of database/sql to prevent accidental SQL injections.
//
// Usage should ideally be identical to the standard sql package with the exception that strings
// should be String instead.
//
// The simplest way to transition to this package is to turn statements like
//
//	db.Query("SELECT ...", args...)
//
// into
//
//	db.Query(safesql.New("SELECT ..."), args...)
//
// Once safesql is adopted, importing database/sql should be banned with the sole
// exception of this package.
// Types from database/sql that are inherently safe have aliases in this package to
// allow for an easier transition and smaller allowlists.
package safesql

import (
	"strconv"
	"strings"

	"github.com/empijei/go-safer-web/safesql/internal/raw"
)

func init() {
	// Initialize the bypass mechanisms for unchecked and legacy conversions.
	raw.StringCtor = func(unsafe string) String { return String{unsafe} }
}

type stringConstant string

// String wraps a string that is safe and does not contain user-controlled input.
type String struct {
	s string
}

// New constructs a String from a compile-time constant string.
// Since the stringConstant type is unexported the only way to call this function
// outside of this package is to pass a string literal or an untyped string const.
//
// Note(empijei): this can be bypassed by using generics with ~string, but that
// feels very unlikely to happen by accident and malicious programmers are not
// part of the threat model of this package.
func New(text stringConstant) String { return String{string(text)} }

// NewFromUint64 constructs a String from a uint64.
func NewFromUint64(i uint64) String { return String{strconv.FormatUint(i, 10)} }

// StringConcat concatenates the given [String]s into a trusted string.
//
// Note(empijei): this function may be abused to create arbitrary queries from
// user inputs, but malicious programmers are not part of the threat model of this package.
func StringConcat(ss ...String) String {
	return StringJoin(ss, String{})
}

// StringJoin joins the given [String]s with the given separator the same way strings.Join would.
//
// Note(empijei): this function may be abused to create arbitrary queries from
// user inputs, but malicious programmers are not part of the threat model of this package.
func StringJoin(ss []String, sep String) String {
	accum := make([]string, 0, len(ss))
	for _, s := range ss {
		accum = append(accum, s.s)
	}
	return String{strings.Join(accum, sep.s)}
}

// String returns the internal representation of the [String], safe to be used
// with a sql engine.
func (t String) String() string {
	return t.s
}

// StringSplit functions as strings.Split but for [String]s.
func StringSplit(s String, sep String) []String {
	spl := strings.Split(s.s, sep.s)
	accum := make([]String, 0, len(spl))
	for _, s := range spl {
		accum = append(accum, String{s})
	}
	return accum
}
