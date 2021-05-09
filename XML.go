package minifiers

import "strings"

func XML(source string) (string, error) {
	result := strings.TrimSpace(source)

	MODE_NORMAL := 0
	MODE_BLOCK_COMMENT := 1

	intMode := MODE_NORMAL

	builder := strings.Builder{}
	runes := []rune(result)
	length := len(runes)
	for i := 0; i < length; i++ {
		// Block comment
		if intMode == MODE_NORMAL {
			if runes[i] == '<' {
				if i+3 < length {
					if runes[i+1] == '!' && runes[i+2] == '-' && runes[i+3] == '-' {
						intMode = MODE_BLOCK_COMMENT
					}
				}
			}
		} else if intMode == MODE_BLOCK_COMMENT {
			if runes[i] == '>' {
				if i-2 >= 0 {
					if runes[i-1] == '-' && runes[i-2] == '-' {
						intMode = MODE_NORMAL
						i++
					}
				}
			}
		}
		// End of block comment

		if intMode != MODE_BLOCK_COMMENT {
			if i < length {
				builder.WriteRune(runes[i])
			}
		}
	}

	result = strings.ReplaceAll(builder.String(), "\n", "")
	result = strings.ReplaceAll(result, "\r", "")
	result = strings.ReplaceAll(result, "\t", "")

	// Make all white space to short
	for strings.Contains(result, "  ") {
		result = strings.ReplaceAll(result, "  ", " ")
	}

	for strings.Contains(result, " =") {
		result = strings.ReplaceAll(result, " =", "=")
	}
	for strings.Contains(result, "= ") {
		result = strings.ReplaceAll(result, "= ", "=")
	}

	for strings.Contains(result, " <") {
		result = strings.ReplaceAll(result, " <", "<")
	}
	for strings.Contains(result, "< ") {
		result = strings.ReplaceAll(result, "< ", "<")
	}

	for strings.Contains(result, " >") {
		result = strings.ReplaceAll(result, " >", ">")
	}
	for strings.Contains(result, "> ") {
		result = strings.ReplaceAll(result, "> ", ">")
	}

	return result, nil
}
