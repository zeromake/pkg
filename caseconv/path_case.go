package caseconv


func PathCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = "/"
	buildOptions(&opt, options)
	return Case(input, opt)
}
