package minifiers

import "strings"

func JS(source string) (string, error) {
	result := strings.TrimSpace(source)
	result = strings.ReplaceAll(result, "\n", "")
	result = strings.ReplaceAll(result, "\r", "")
	return result, nil
}
