package caseconv

import (
	"strings"
)

func CamelCaseTransform(input string, index int) string {
	if index == 0 {
		return strings.ToLower(input)
	}
	return PascalCaseTransform(input, index)
}

func CamelCaseTransformMerge(input string, index int) string {
	if index == 0 {
		return strings.ToLower(input)
	}
	return PascalCaseTransformMerge(input)
}

func CamelCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = ""
	opt.Transform = CamelCaseTransform
	buildOptions(&opt, options)
	return Case(input, opt)
}
