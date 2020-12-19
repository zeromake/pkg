package textconv

import "strings"

func IsLowerCase(input string) bool {
	return input == strings.ToLower(input) && input != strings.ToUpper(input)
}

func IsUpperCase(input string) bool {
	return input == strings.ToUpper(input) && input != strings.ToLower(input)
}
