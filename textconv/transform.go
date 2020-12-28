package textconv

import (
	"strings"
)

func CamelCaseTransform(input string, index int) string {
	if index == 0 {
		return strings.ToLower(input)
	}
	return PascalCaseTransform(input, index)
}

func CapitalCaseTransform(input string, _ int) string {
	return PascalCaseTransformMerge(input)
}

func ConstantCaseTransform(input string, _ int) string {
	return strings.ToUpper(input)
}

func NoCaseTransform(input string, _ int) string {
	return strings.ToLower(input)
}

func PascalCaseTransform(input string, index int) string {
	var firstChar = input[0]
	if index > 0 && firstChar >= '0' && firstChar <= '9' {
		return "_" + strings.ToLower(input)
	}
	return PascalCaseTransformMerge(input)
}

func PascalCaseTransformMerge(s string) string {
	hasUpper := false
	length := len(s)
	for i := 1; i < length; i++ {
		c := s[i]
		hasUpper = hasUpper || ('A' <= c && c <= 'Z')
	}
	firstChar := s[0]
	firstIsLower := 'a' <= firstChar && firstChar <= 'z'
	if !hasUpper && !firstIsLower {
		return s
	}
	var b strings.Builder
	b.Grow(length)
	if firstIsLower {
		firstChar -= 32
	}
	b.WriteByte(firstChar)
	if hasUpper {
		for i := 1; i < length; i++ {
			c := s[i]
			if 'A' <= c && c <= 'Z' {
				c += 32
			}
			b.WriteByte(c)
		}
	} else {
		b.WriteString(s[1:])
	}
	return b.String()
}

func SentenceCaseTransform(input string, index int) string {
	if index == 0 {
		return PascalCaseTransformMerge(input)
	}
	return strings.ToLower(input)
}

var commonMap = map[string]struct{}{
	"Api":   {},
	"Ascii": {},
	"Cpu":   {},
	"Css":   {},
	"Dns":   {},
	"Eof":   {},
	"Guid":  {},
	"Html":  {},
	"Https": {},
	"Http":  {},
	"Id":    {},
	"Ip":    {},
	"Json":  {},
	"Lhs":   {},
	"Qps":   {},
	"Ram":   {},
	"Rhs":   {},
	"Rpc":   {},
	"Sla":   {},
	"Smtp":  {},
	"Ssh":   {},
	"Tls":   {},
	"Ttl":   {},
	"Uuid":  {},
	"Uid":   {},
	"Ui":    {},
	"Uri":   {},
	"Url":   {},
	"Utf8":  {},
	"Vm":    {},
	"Xml":   {},
	"Xsrf":  {},
	"Xss":   {},
}

func FieldCaseTransform(transform func(string, int) string, input string, index int, prevUpper bool) (string, bool) {
	input = transform(input, index)
	if _, ok := commonMap[input]; ok && !prevUpper {
		return strings.ToUpper(input), true
	}
	return input, false
}
