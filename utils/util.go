package utils

import "net/url"

func IsValidURL(raw string) bool {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}

	// require http or https
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	// must have a host
	if u.Host == "" {
		return false
	}

	return true
}
