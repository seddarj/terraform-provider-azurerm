package helpers

import "strings"

func trimURLScheme(input string) string {
	schemes := []string{
		"https://",
		"http://",
	}
	for _, v := range schemes {
		if strings.HasPrefix(strings.ToLower(input), v) {
			return input[len(v):]
		}
	}

	return input
}
