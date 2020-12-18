package textconv

import (
	"strings"
)

func PascalCaseTransform(input string, index int) string {
	var firstChar = input[0]
	if index > 0 && firstChar >= '0' && firstChar <= '9' {
		return "_" + strings.ToLower(input)
	}
	return PascalCaseTransformMerge(input)
}

func PascalCaseTransformMerge(input string) string {
	return strings.ToUpper(input[:1]) + strings.ToLower(input[1:])
}

func PascalCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = ""
	opt.Transform = PascalCaseTransform
	buildOptions(&opt, options)
	return NoCase(input, opt)
}
