package minifiers

import "fmt"

func CSS(source string) (string, error) {
	fmt.Println("CSS")
	result := BASE(source)
	return result, nil
}
