package raw

// StringCtor is the constructor for a safesql.String only intended to be used by unchecked and legacy conversions.
// It will be assigned by the safesql package at init time.
// This is an empty interface is to avoid cyclic dependencies.
var StringCtor interface{}
