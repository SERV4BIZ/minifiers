package minifiers

import (
	"fmt"
	"strings"
)

func BASE(source string) string {
	fmt.Println(source)
	result := strings.TrimSpace(source)
	result = strings.ReplaceAll(result, "\n", "")
	result = strings.ReplaceAll(result, "\r", "")
	fmt.Println(result)
	return result
}
