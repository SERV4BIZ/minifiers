package minifiers

func JS(source string) (string, error) {
	result := BASE(source)
	return result, nil
}
