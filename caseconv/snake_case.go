package caseconv

func SnakeCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = "_"
	buildOptions(&opt, options)
	return Case(input, opt)
}
