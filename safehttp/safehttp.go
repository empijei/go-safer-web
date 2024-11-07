// Package safehttp provides framework-agnostic utilities to add security layers to go HTTP servers.
//
// Most go web frameworks export a "root" type that implements http.Handler, for example:
// * gin.Engine
// * echo.Echo
// * httprouter.Router
// * buffalo.App
// In case you're using or plan to use such libraries, please wrap the root handler
// with the security middleware.
// If, instead, you're using vanilla go net/http this should be done on the http.Router
// or as the root handler on the http.Server.
package safehttp

type ConfigValue int

const (
	Enabled ConfigValue = iota
	Disabled
	ReportOnly
)

type CSRFConfigValue struct {
	ConfigValue
	// UseFormData makes the middleware parse forms and check for the token in
	// the request body instead of the request header.
	UseFormData bool
}

type MethodsConfigValue struct {
	ConfigValue
	AllowedMethods []string
}

type HostCheckConfigValue struct {
	ConfigValue
	ExpectedHosts []string
}

type Config struct {
	FramingProtection       ConfigValue
	CrossOriginOpenerPolicy ConfigValue
	StrictCSP               ConfigValue
	TrustedTypes            ConfigValue
	ResourceIsolation       ConfigValue
	StrictContentType       ConfigValue // TODO client and server
	DisableXSSFiltering     ConfigValue
	HSTS                    ConfigValue
	StrictCookies           ConfigValue

	CSRFProtection   CSRFConfigValue
	MethodsAllowList MethodsConfigValue
	HostCheck        HostCheckConfigValue
}
