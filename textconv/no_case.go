package textconv

import (
	"regexp"
	"strings"
)

var (
	splitRepl = "$1\000$2"
	stripRepl = "\000"
)

var (
	DefaultOptions = Options{
		//SplitRegexp: []*regexp.Regexp{
		//	regexp.MustCompile("([a-z0-9])([A-Z])"),
		//	regexp.MustCompile("([A-Z])([A-Z][a-z])"),
		//},
		//StripRegexp: []*regexp.Regexp{
		//	regexp.MustCompile("[^a-zA-Z0-9]+"),
		//},
		Delimiter: " ",
		Transform: func(i string, _ int) string {
			return strings.ToLower(i)
		},
	}
)

type Option func(option *Options)

func buildOptions(option *Options, options []Option) {
	for _, o := range options {
		o(option)
	}
}

type Options struct {
	SplitRegexp []*regexp.Regexp
	StripRegexp []*regexp.Regexp
	Delimiter   string
	Transform   func(word string, index int) string
}

func (o *Options) Clone() *Options {
	return &Options{
		Delimiter: o.Delimiter,
		Transform: o.Transform,
	}
}

func replace(input string, regexps []*regexp.Regexp, value string) string {
	var v = input
	for _, re := range regexps {
		v = re.ReplaceAllString(v, value)
	}
	return v
}

//func EqualRegexp(a, b []*regexp.Regexp) bool {
//	if &a == &b {
//		return true
//	}
//	if len(a) != len(b) {
//		return false
//	}
//	n := len(a)
//	flag := true
//	for i := 0; i < n; i++ {
//		flag = flag || a[i] == b[i]
//	}
//	return flag
//}

func NoCase(input string, options Options) string {
	var result []string
	if options.SplitRegexp == nil && options.StripRegexp == nil {
		result = SplitString(input)
		for i, a := range result {
			if options.Transform != nil {
				a = options.Transform(a, i)
			}
			result[i] = a
		}
	} else {
		var split = replace(
			input,
			options.SplitRegexp,
			splitRepl,
		)
		var ss = replace(
			split,
			options.StripRegexp,
			stripRepl,
		)
		arr := strings.Split(ss, stripRepl)
		i := 0
		for _, a := range arr {
			if len(a) == 0 {
				continue
			}
			if options.Transform != nil {
				a = options.Transform(a, i)
			}
			result = append(result, a)
			i++
		}
	}
	return strings.Join(result, options.Delimiter)
}

//func RegexpSplitString(input string) []string {
//	var split = replace(
//		input,
//		DefaultOptions.SplitRegexp,
//		splitRepl,
//	)
//	var result = replace(
//		split,
//		DefaultOptions.StripRegexp,
//		stripRepl,
//	)
//	arr := strings.Split(result, stripRepl)
//	var ret []string
//	for _, a := range arr {
//		if len(a) == 0 {
//			continue
//		}
//		ret = append(ret, a)
//	}
//	return ret
//}

//func SplitBytes(input []byte) [][]byte {
//	var (
//		result     [][]byte
//		currCase   bool
//		currNumber bool
//		nextCase   bool
//		prevCase   bool
//		prevNumber bool
//		offset     int
//		length     = len(input) - 1
//	)
//	for i := 0; i < length; i++ {
//		v := input[i]
//		nextV := input[i+1]
//		currCase = v >= 'A' && v <= 'Z'
//		currNumber = !currCase && v >= '0' && v <= '9'
//		nextCase = nextV >= 'A' && nextV <= 'Z'
//		if currCase || currNumber || (v >= 'a' && v <= 'z') {
//			if offset > 0 {
//				if ((!prevCase || prevNumber) && currCase) || (prevCase && currCase && !nextCase) {
//					result = append(result, input[i-offset:i])
//					offset = 0
//				}
//			}
//			offset++
//		} else if offset > 0 {
//			result = append(result, input[i-offset:i])
//			offset = 0
//		}
//		prevCase = currCase
//		prevNumber = currNumber
//	}
//	if offset > 0 {
//		v := input[length]
//		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
//			length++
//			offset++
//		}
//		result = append(result, input[length-offset:length])
//	}
//	return result
//}

func SplitString(input string) []string {
	var (
		result     []string
		currCase   bool
		currNumber bool
		nextCase   bool
		prevCase   bool
		prevNumber bool
		offset     int
		length     = len(input) - 1
	)
	for i := 0; i < length; i++ {
		v := input[i]
		nextV := input[i+1]
		currCase = v >= 'A' && v <= 'Z'
		currNumber = !currCase && v >= '0' && v <= '9'
		nextCase = nextV >= 'A' && nextV <= 'Z'
		if currCase || currNumber || (v >= 'a' && v <= 'z') {
			if offset > 0 {
				if ((!prevCase || prevNumber) && currCase) || (prevCase && currCase && !nextCase) {
					result = append(result, input[i-offset:i])
					offset = 0
				}
			}
			offset++
		} else if offset > 0 {
			result = append(result, input[i-offset:i])
			offset = 0
		}
		prevCase = currCase
		prevNumber = currNumber
	}
	if offset > 0 {
		v := input[length]
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			length++
			offset++
		}
		result = append(result, input[length-offset:length])
	}
	return result
}
