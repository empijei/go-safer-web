package legacyconversions

import (
	"github.com/empijei/go-safer-web/safesql"
	"github.com/empijei/go-safer-web/safesql/internal/raw"
)

var safesqlStringCtor = raw.StringCtor.(func(string) safesql.String)

// KnownSafeString riskily promotes the given string to a trusted string.
// Uses of this function should be carefully reviewed to make sure that no user input
// can ever be passed to it.
//
// Examples of safe usages are to promote strings stored in external query storages
// under the programmer control, or startup flags.
func KnownSafeString(trusted string) safesql.String {
	return safesqlStringCtor(trusted)
}
