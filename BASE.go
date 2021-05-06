package minifiers

import (
	"strings"
)

func BASE(source string) string {
	result := strings.TrimSpace(source)
	result = strings.ReplaceAll(result, "\n", "")
	result = strings.ReplaceAll(result, "\r", "")
	result = strings.ReplaceAll(result, "\t", "")
	return result
}
