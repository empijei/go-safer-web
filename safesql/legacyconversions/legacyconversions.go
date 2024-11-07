package legacyconversions

import (
	"github.com/empijei/go-safer-web/safesql"
	"github.com/empijei/go-safer-web/safesql/internal/raw"
)

var safesqlStringCtor = raw.StringCtor.(func(string) safesql.String)

// UnsafeSQLString riskily promotes the given string to a trusted string.
// Uses of this function should only be introduced to begin a migration to safesql
// and should eventually be removed.
func UnsafeSQLString(trusted string) safesql.String {
	return safesqlStringCtor(trusted)
}
