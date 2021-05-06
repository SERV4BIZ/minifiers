package minifiers

func HTML(source string) (string, error) {
	result := BASE(source)
	return result, nil
}
