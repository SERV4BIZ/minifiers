package minifiers

import "strings"

func CSS(source string) (string, error) {
	result := strings.TrimSpace(source)
	result = strings.ReplaceAll(result, "\n", "")
	result = strings.ReplaceAll(result, "\r", "")
	return result, nil
}
