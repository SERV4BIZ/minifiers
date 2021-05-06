package minifiers

func JSON(source string) (string, error) {
	result := BASE(source)
	return result, nil
}
