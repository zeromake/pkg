package textconv

import "strings"

func IsLowerCase(input string) bool {
	if len(input) == 0 {
		return false
	}
	return input == strings.ToLower(input) && input != strings.ToUpper(input)
}

func IsUpperCase(input string) bool {
	if len(input) == 0 {
		return false
	}
	return input == strings.ToUpper(input) && input != strings.ToLower(input)
}
