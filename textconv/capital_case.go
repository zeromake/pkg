package textconv

func CapitalCaseTransform(input string, _ int) string {
	return PascalCaseTransformMerge(input)
}

func CapitalCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = " "
	opt.Transform = CapitalCaseTransform
	buildOptions(&opt, options)
	return NoCase(input, opt)
}
