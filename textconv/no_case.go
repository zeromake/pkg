package textconv

import (
	"strings"
)

var (
	DefaultOptions = Options{
		Split:     SplitString,
		Delimiter: " ",
		Transform: NoCaseTransform,
	}
)

type Option func(option *Options)

func buildOptions(option *Options, options []Option) {
	for _, o := range options {
		o(option)
	}
}

type Options struct {
	Split     func(input string) []string
	Delimiter string
	Transform func(word string, index int) string
}

func NoCase(input string, options Options) string {
	if len(input) == 0 {
		return input
	}
	var result []string
	result = options.Split(input)
	for i, a := range result {
		if options.Transform != nil {
			a = options.Transform(a, i)
		}
		result[i] = a
	}
	return strings.Join(result, options.Delimiter)
}
