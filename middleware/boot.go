package middleware

import "net/http"

func Apply(h http.HandlerFunc, m ...func(http.Handler) http.Handler) http.Handler {
	var res http.Handler = h
	for _, middleware := range m {
		res = middleware(res)
	}
	return res
}
