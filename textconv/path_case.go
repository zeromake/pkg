package textconv

func PathCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = "/"
	buildOptions(&opt, options)
	return NoCase(input, opt)
}
