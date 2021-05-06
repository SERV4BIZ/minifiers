package minifiers

import "strings"

func BASE(source string) string {
	result := strings.TrimSpace(source)
	result = strings.ReplaceAll(result, "\n", "")
	result = strings.ReplaceAll(result, "\r", "")
	return result
}
