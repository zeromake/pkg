package textconv

func SnakeCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = "_"
	buildOptions(&opt, options)
	return NoCase(input, opt)
}
