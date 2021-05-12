package minifiers

func HTML(source string) (string, error) {
	compJS, err := JS(source)
	if err != nil {
		return source, err
	}

	compCSS, err := CSS(compJS)
	if err != nil {
		return source, err
	}

	compXML, err := XML(compCSS)
	if err != nil {
		return source, err
	}

	return compXML, nil
}
