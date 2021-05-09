package minifiers

import (
	"errors"
	"strings"

	"github.com/SERV4BIZ/gfp/jsons"
)

func JSON(source string) (string, error) {
	result := strings.TrimSpace(source)
	runes := []rune(result)

	if runes[0] == '{' {
		jsoResult, err := jsons.ObjectString(result)
		if err != nil {
			return "", err
		}
		return jsoResult.ToString(), nil
	} else if runes[0] == '[' {
		jsaResult, err := jsons.ArrayString(result)
		if err != nil {
			return "", err
		}
		return jsaResult.ToString(), nil
	}
	return "", errors.New("source format is not support")
}
