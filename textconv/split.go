package textconv

import (
	"regexp"
	"strings"
)

var (
	splitRepl = "$1\000$2"
	stripRepl = "\000"
)

func replace(input string, regexps []*regexp.Regexp, value string) string {
	var v = input
	for _, re := range regexps {
		v = re.ReplaceAllString(v, value)
	}
	return v
}

var splitRegexp = []*regexp.Regexp{
	regexp.MustCompile("([a-z0-9])([A-Z])"),
	regexp.MustCompile("([A-Z])([A-Z][a-z])"),
}
var stripRegexp = []*regexp.Regexp{
	regexp.MustCompile("[^a-zA-Z0-9]+"),
}

func RegexpSplitString(input string) []string {
	var split = replace(
		input,
		splitRegexp,
		splitRepl,
	)
	var result = replace(
		split,
		stripRegexp,
		stripRepl,
	)
	arr := strings.Split(result, stripRepl)
	var ret []string
	for _, a := range arr {
		if len(a) == 0 {
			continue
		}
		ret = append(ret, a)
	}
	return ret
}

// SplitString 分割字符串 0-9a-zA-Z 的词语
func SplitString(input string) []string {
	var (
		// 切割后的结果字符串数组
		result []string
		// 当前字符为大写
		currCase bool
		// 当前字符为数字
		currNumber bool
		// 下一个字符为大写
		nextCase bool
		// 上一个字符为大写
		prevCase bool
		// 上一个字符为数字
		prevNumber bool
		// 需要分割字符串的长度
		offset int
		// 输入字符串长度 -1
		length = len(input) - 1
	)
	for i := 0; i < length; i++ {
		v := input[i]
		nextV := input[i+1]
		currCase = v >= 'A' && v <= 'Z'
		currNumber = !currCase && v >= '0' && v <= '9'
		nextCase = nextV >= 'A' && nextV <= 'Z'
		if currCase || currNumber || (v >= 'a' && v <= 'z') {
			// 当前字符为字母或数字
			if offset > 0 {
				// 有需要切割的字符串
				if ((!prevCase || prevNumber) && currCase) || (prevCase && currCase && !nextCase) {
					// 满足 [a-z0-9][A-Z] 或 [A-Z][A-Z][a-z] 情况
					result = append(result, input[i-offset:i])
					offset = 0
				}
			}
			offset++
		} else if offset > 0 {
			// 碰上不为字母或数字
			result = append(result, input[i-offset:i])
			offset = 0
		}
		prevCase = currCase
		prevNumber = currNumber
	}
	v := input[length]
	if nextCase || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
		length++
		offset++
	}
	if offset > 0 {
		result = append(result, input[length-offset:length])
	}
	return result
}
