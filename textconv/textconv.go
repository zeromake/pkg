package textconv


// CamelCase convert to `camelCase`
func CamelCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = ""
	opt.Transform = CamelCaseTransform
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// CapitalCase convert to `Capital Case`
func CapitalCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = " "
	opt.Transform = CapitalCaseTransform
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// ConstantCase convert to `CONSTANT_CASE`
func ConstantCase(input string, options... Option) string {
	opt := DefaultOptions
	opt.Delimiter = "_"
	opt.Transform = ConstantCaseTransform
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// DotCase convert to `dot.case`
func DotCase(input string, options... Option) string {
	opt := DefaultOptions
	opt.Delimiter = "."
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// HeaderCase convert to `Header-Case`
func HeaderCase(input string, options... Option) string {
	opt := DefaultOptions
	opt.Delimiter = "-"
	opt.Transform = CapitalCaseTransform
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// ParamCase convert to `param-case`
func ParamCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = "-"
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// PathCase convert to `path/case`
func PathCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = "/"
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// PascalCase convert to `PascalCase`
func PascalCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = ""
	opt.Transform = PascalCaseTransform
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// SnakeCase convert to `snake_case`
func SnakeCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Delimiter = "_"
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

// SentenceCase convert to `Sentence case`
func SentenceCase(input string, options ...Option) string {
	opt := DefaultOptions
	opt.Transform = SentenceCaseTransform
	buildOptions(&opt, options)
	return NoCase(input, opt)
}

