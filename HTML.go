package minifiers

func HTML(source string) (string, error) {
	compXML, err := XML(source)
	if err != nil {
		return source, err
	}

	return compXML, nil
}
