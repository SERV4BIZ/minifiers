package minifiers

func CSS(source string) (string, error) {
	result := BASE(source)
	return result, nil
}
