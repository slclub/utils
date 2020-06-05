package utils

import (
	"net/url"
	"strings"
)

// URL======================================================================
// Get string by key from url.Values.
// Parse the of url.Values to comme used type.
func GetStringFromUrl(mv url.Values, key string) (value string, ret bool) {
	if mv == nil {
		return "", false
	}
	for k, s := range mv {
		if k == key && len(s) >= 1 {
			value = s[0]
			ret = true
			return
		}
	}
	return "", false
}

// Get array string by key from url.Values
func GetArrayFromUrl(mv url.Values, key string) ([]string, bool) {
	sl := make([]string, 0)
	ret := false

	for k, s := range mv {
		if i := strings.IndexByte(k, '['); i >= 1 && k[:i] == key {
			if j := strings.IndexByte(k[i+1:], ']'); j == 0 {
				sl = append(sl, s...)
				ret = true
				return sl, ret
			}
		}
	}
	return nil, ret
}

// Get map[string]string by key from url.Values.
func GetMapFromUrl(mv url.Values, key string) (map[string]string, bool) {
	dicts := make(map[string]string)
	exist := false
	for k, v := range mv {
		if i := strings.IndexByte(k, '['); i >= 1 && k[0:i] == key {
			if j := strings.IndexByte(k[i+1:], ']'); j >= 1 {
				exist = true
				dicts[k[i+1:][:j]] = v[0]
			}
		}
	}
	return dicts, exist
}

// URL======================================================================
// Request header===========================================================
func GetPartFilterTrimOrSemicolon(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

// Request header===========================================================
