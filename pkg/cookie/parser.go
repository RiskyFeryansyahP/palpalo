package cookie

import "strings"

func GetCookieByPrefix(cookies []string, prefix string) string {
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie, prefix) {
			return cookie
		}
	}

	return ""
}
