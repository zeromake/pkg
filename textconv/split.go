package textconv

// split
// "$1\000$2"
//[]*regexp.Regexp{
//	regexp.MustCompile("([a-z0-9])([A-Z])"),
//	regexp.MustCompile("([A-Z])([A-Z][a-z])"),
//}
// strip
// "\000"
//[]*regexp.Regexp{
//	regexp.MustCompile("[^a-zA-Z0-9]+"),
//}
//
//func replace(input string, regexps []*regexp.Regexp, value string) string {
//	var v = input
//	for _, re := range regexps {
//		v = re.ReplaceAllString(v, value)
//	}
//	return v
//}

// SplitString 分割字符串 0-9a-zA-Z 的词语
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
	v := input[length]
	if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
		length++
		offset++
	}
	if offset > 0 {
		result = append(result, input[length-offset:length])
	}
	return result
}
