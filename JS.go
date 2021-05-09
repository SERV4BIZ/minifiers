package minifiers

import (
	"strings"
)

func JS(source string) (string, error) {
	result := strings.TrimSpace(source)

	MODE_NORMAL := 0
	MODE_SINGLE_COMMENT := 1
	MODE_BLOCK_COMMENT := 2
	MODE_SINGLE_QUOTE := 3
	MODE_DOUBLE_QUOTE := 4
	MODE_BACKTICKS := 5

	intMode := MODE_NORMAL

	builder := strings.Builder{}
	runes := []rune(result)
	length := len(runes)

	// Compress processing
	for i := 0; i < length; i++ {
		// Single comment
		if intMode == MODE_NORMAL {
			if runes[i] == '/' {
				if i-1 >= 0 {
					if runes[i-1] != '\\' {
						if i+1 < length {
							if runes[i+1] == '/' {
								intMode = MODE_SINGLE_COMMENT
							}
						}
					}
				} else {
					if i+1 < length {
						if runes[i+1] == '/' {
							intMode = MODE_SINGLE_COMMENT
						}
					}
				}
			}
		} else if intMode == MODE_SINGLE_COMMENT {
			if runes[i] == '\n' {
				intMode = MODE_NORMAL
				i++
			}
		}
		// End of Single comment

		// Block comment
		if intMode == MODE_NORMAL {
			if runes[i] == '/' {
				if i-1 >= 0 {
					if runes[i-1] != '\\' {
						if i+1 < length {
							if runes[i+1] == '*' {
								intMode = MODE_BLOCK_COMMENT
							}
						}
					}
				} else {
					if i+1 < length {
						if runes[i+1] == '*' {
							intMode = MODE_BLOCK_COMMENT
						}
					}
				}
			}
		} else if intMode == MODE_BLOCK_COMMENT {
			if runes[i] == '/' {
				if i-1 >= 0 {
					if runes[i-1] == '*' {
						intMode = MODE_NORMAL
						i++
					}
				}
			}
		}
		// End of block comment

		// String single quote
		if intMode == MODE_NORMAL {
			if runes[i] == '\'' {
				if i-1 >= 0 {
					if runes[i-1] != '\\' {
						intMode = MODE_SINGLE_QUOTE
					}
				} else {
					intMode = MODE_SINGLE_QUOTE
				}
			}
		} else if intMode == MODE_SINGLE_QUOTE {
			if runes[i] == '\'' {
				if i-1 >= 0 {
					if runes[i-1] != '\\' {
						intMode = MODE_NORMAL
					}
				}
			}
		}
		// End of single quote

		// String double quote
		if intMode == MODE_NORMAL {
			if runes[i] == '"' {
				if i-1 >= 0 {
					if runes[i-1] != '\\' {
						intMode = MODE_DOUBLE_QUOTE
					}
				} else {
					intMode = MODE_DOUBLE_QUOTE
				}
			}
		} else if intMode == MODE_DOUBLE_QUOTE {
			if runes[i] == '"' {
				if i-1 >= 0 {
					if runes[i-1] != '\\' {
						intMode = MODE_NORMAL
					}
				}
			}
		}
		// End of double quote

		// String backticks
		if intMode == MODE_NORMAL {
			if runes[i] == '`' {
				if i-1 >= 0 {
					if runes[i-1] != '\\' {
						intMode = MODE_BACKTICKS
					}
				} else {
					intMode = MODE_BACKTICKS
				}
			}
		} else if intMode == MODE_BACKTICKS {
			if runes[i] == '`' {
				if i-1 >= 0 {
					if runes[i-1] != '\\' {
						intMode = MODE_NORMAL
					}
				}
			}
		}
		// End of backticks

		if intMode != MODE_SINGLE_COMMENT && intMode != MODE_BLOCK_COMMENT {
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

	for strings.Contains(result, " {") {
		result = strings.ReplaceAll(result, " {", "{")
	}

	for strings.Contains(result, "{ ") {
		result = strings.ReplaceAll(result, "{ ", "{")
	}

	for strings.Contains(result, " }") {
		result = strings.ReplaceAll(result, " }", "}")
	}

	for strings.Contains(result, "} ") {
		result = strings.ReplaceAll(result, "} ", "}")
	}

	for strings.Contains(result, " (") {
		result = strings.ReplaceAll(result, " (", "(")
	}

	for strings.Contains(result, "( ") {
		result = strings.ReplaceAll(result, "( ", "(")
	}

	for strings.Contains(result, " )") {
		result = strings.ReplaceAll(result, " )", ")")
	}

	for strings.Contains(result, ") ") {
		result = strings.ReplaceAll(result, ") ", ")")
	}

	for strings.Contains(result, " [") {
		result = strings.ReplaceAll(result, " [", "[")
	}

	for strings.Contains(result, "[ ") {
		result = strings.ReplaceAll(result, "[ ", "[")
	}

	for strings.Contains(result, " ]") {
		result = strings.ReplaceAll(result, " ]", "]")
	}

	for strings.Contains(result, "] ") {
		result = strings.ReplaceAll(result, "] ", "]")
	}

	for strings.Contains(result, " :") {
		result = strings.ReplaceAll(result, " :", ":")
	}

	for strings.Contains(result, ": ") {
		result = strings.ReplaceAll(result, ": ", ":")
	}

	for strings.Contains(result, " ;") {
		result = strings.ReplaceAll(result, " ;", ";")
	}

	for strings.Contains(result, "; ") {
		result = strings.ReplaceAll(result, "; ", ";")
	}

	// ,
	for strings.Contains(result, " ,") {
		result = strings.ReplaceAll(result, " ,", ",")
	}
	for strings.Contains(result, ", ") {
		result = strings.ReplaceAll(result, ", ", ",")
	}

	// Arithmetic Operators
	// +
	for strings.Contains(result, " +") {
		result = strings.ReplaceAll(result, " +", "+")
	}
	for strings.Contains(result, "+ ") {
		result = strings.ReplaceAll(result, "+ ", "+")
	}

	// -
	for strings.Contains(result, " -") {
		result = strings.ReplaceAll(result, " -", "-")
	}
	for strings.Contains(result, "- ") {
		result = strings.ReplaceAll(result, "- ", "-")
	}

	// *
	for strings.Contains(result, " *") {
		result = strings.ReplaceAll(result, " *", "*")
	}
	for strings.Contains(result, "* ") {
		result = strings.ReplaceAll(result, "* ", "*")
	}

	// div
	for strings.Contains(result, " /") {
		result = strings.ReplaceAll(result, " /", "/")
	}
	for strings.Contains(result, "/ ") {
		result = strings.ReplaceAll(result, "/ ", "/")
	}

	// %
	for strings.Contains(result, " %") {
		result = strings.ReplaceAll(result, " %", "%")
	}
	for strings.Contains(result, "% ") {
		result = strings.ReplaceAll(result, "% ", "%")
	}

	// ++
	for strings.Contains(result, " ++") {
		result = strings.ReplaceAll(result, " ++", "++")
	}
	for strings.Contains(result, "++ ") {
		result = strings.ReplaceAll(result, "++ ", "++")
	}

	// --
	for strings.Contains(result, " --") {
		result = strings.ReplaceAll(result, " --", "--")
	}
	for strings.Contains(result, "-- ") {
		result = strings.ReplaceAll(result, "-- ", "--")
	}

	// Comparison Operators
	// ==
	for strings.Contains(result, " ==") {
		result = strings.ReplaceAll(result, " ==", "==")
	}
	for strings.Contains(result, "== ") {
		result = strings.ReplaceAll(result, "== ", "==")
	}

	// ===
	for strings.Contains(result, " ===") {
		result = strings.ReplaceAll(result, " ===", "===")
	}
	for strings.Contains(result, "=== ") {
		result = strings.ReplaceAll(result, "=== ", "===")
	}

	// !=
	for strings.Contains(result, " !=") {
		result = strings.ReplaceAll(result, " !=", "!=")
	}
	for strings.Contains(result, "!= ") {
		result = strings.ReplaceAll(result, "!= ", "!=")
	}

	// >
	for strings.Contains(result, " >") {
		result = strings.ReplaceAll(result, " >", ">")
	}
	for strings.Contains(result, "> ") {
		result = strings.ReplaceAll(result, "> ", ">")
	}

	// <
	for strings.Contains(result, " <") {
		result = strings.ReplaceAll(result, " <", "<")
	}
	for strings.Contains(result, "< ") {
		result = strings.ReplaceAll(result, "< ", "<")
	}

	// >=
	for strings.Contains(result, " >=") {
		result = strings.ReplaceAll(result, " >=", ">=")
	}
	for strings.Contains(result, ">= ") {
		result = strings.ReplaceAll(result, ">= ", ">=")
	}

	// <=
	for strings.Contains(result, " <=") {
		result = strings.ReplaceAll(result, " <=", "<=")
	}
	for strings.Contains(result, "<= ") {
		result = strings.ReplaceAll(result, "<= ", "<=")
	}

	// Logical Operators
	// and &&
	for strings.Contains(result, " &&") {
		result = strings.ReplaceAll(result, " &&", "&&")
	}
	for strings.Contains(result, "&& ") {
		result = strings.ReplaceAll(result, "&& ", "&&")
	}

	// or ||
	for strings.Contains(result, " ||") {
		result = strings.ReplaceAll(result, " ||", "||")
	}
	for strings.Contains(result, "|| ") {
		result = strings.ReplaceAll(result, "|| ", "||")
	}

	// Not !
	for strings.Contains(result, " !") {
		result = strings.ReplaceAll(result, " !", "!")
	}
	for strings.Contains(result, "! ") {
		result = strings.ReplaceAll(result, "! ", "!")
	}

	// Assignment Operators
	// =
	for strings.Contains(result, " =") {
		result = strings.ReplaceAll(result, " =", "=")
	}
	for strings.Contains(result, "= ") {
		result = strings.ReplaceAll(result, "= ", "=")
	}

	// +=
	for strings.Contains(result, " +=") {
		result = strings.ReplaceAll(result, " +=", "+=")
	}
	for strings.Contains(result, "+= ") {
		result = strings.ReplaceAll(result, "+= ", "+=")
	}

	// -=
	for strings.Contains(result, " -=") {
		result = strings.ReplaceAll(result, " -=", "-=")
	}
	for strings.Contains(result, "-= ") {
		result = strings.ReplaceAll(result, "-= ", "-=")
	}

	// *=
	for strings.Contains(result, " *=") {
		result = strings.ReplaceAll(result, " *=", "*=")
	}
	for strings.Contains(result, "*= ") {
		result = strings.ReplaceAll(result, "*= ", "*=")
	}

	// /=
	for strings.Contains(result, " /=") {
		result = strings.ReplaceAll(result, " /=", "/=")
	}
	for strings.Contains(result, "/= ") {
		result = strings.ReplaceAll(result, "/= ", "/=")
	}

	// %=
	for strings.Contains(result, " %=") {
		result = strings.ReplaceAll(result, " %=", "%=")
	}
	for strings.Contains(result, "%= ") {
		result = strings.ReplaceAll(result, "%= ", "%=")
	}

	// Ternary Operator
	// ?
	for strings.Contains(result, " ?") {
		result = strings.ReplaceAll(result, " ?", "?")
	}
	for strings.Contains(result, "? ") {
		result = strings.ReplaceAll(result, "? ", "?")
	}

	return result, nil
}
