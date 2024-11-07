# go-safer-web

An easy to adopt security framework for Go web applications.

This module tries to take the idea behind google/go-safeweb and apply it in the
least intrusive way possible.

It fully relies on standard library types wherever possible (e.g. it doesn't
forbid http.Handler) and created tiny wrappers when needed (e.g. safesql.String).
