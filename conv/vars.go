package conv

import (
	"bytes"
	"strings"

	"github.com/zeromake/pkg/utils"
)

// var pmap sync.Map
// var umap sync.Map
var commonInitialismsReplacer *strings.Replacer
var reCommonInitialismsReplacer *strings.Replacer
var commonInitialisms = []string{
	"API",
	"ASCII",
	"CPU",
	"CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML",
	"HTTPS",
	"HTTP",
	"ID",
	"IP",
	"JSON",
	"LHS",
	"QPS",
	"RAM",
	"RHS",
	"RPC",
	"SLA",
	"SMTP",
	"SSH",
	"TLS",
	"TTL",
	"UUID",
	"UID",
	"UI",
	"URI",
	"URL",
	"UTF8",
	"VM",
	"XML",
	"XSRF",
	"XSS",
}

func init() {
	var (
		replacerLen                    = len(commonInitialisms)
		commonInitialismsForReplacer   = make([]string, replacerLen*2)
		reCommonInitialismsForReplacer = make([]string, replacerLen*2)
	)
	for i := 0; i < replacerLen; i += 2 {
		initialism := commonInitialisms[i]
		lowerInitialism := strings.ToLower(initialism)
		initialismTitle := strings.Title(lowerInitialism)
		commonInitialismsForReplacer[i] = initialism
		commonInitialismsForReplacer[i+1] = initialismTitle
		reCommonInitialismsForReplacer[i] = lowerInitialism
		reCommonInitialismsForReplacer[i+1] = initialism
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
	reCommonInitialismsReplacer = strings.NewReplacer(reCommonInitialismsForReplacer...)
}

// // CacheUnderlineToPascalCase 下滑线转帕斯卡带全局缓存
// func CacheUnderlineToPascalCase(name string) string {
// 	if v, ok := umap.Load(name); ok {
// 		return v.(string)
// 	}
// 	s := UnderlineToPascalCase(name)
// 	if s != "" {
// 		umap.Store(name, s)
// 	}
// 	return s
// }

// // CachePascalCaseToUnderline 帕斯卡转下滑线带全局缓存
// func CachePascalCaseToUnderline(name string) string {
// 	if v, ok := umap.Load(name); ok {
// 		return v.(string)
// 	}
// 	s := PascalCaseToUnderline(name)
// 	if s != "" {
// 		umap.Store(name, s)
// 	}
// 	return s
// }

// UnderlineToPascalCase 下划线命名转帕斯卡命名
func UnderlineToPascalCase(name string) string {
	if name == "" {
		return ""
	}
	var (
		value    = reCommonInitialismsReplacer.Replace(name)
		bb       = utils.StringToBytes(value)
		buf      = make([]byte, len(bb))
		isTitle  = true
		isOneKey = false
	)
	index := 0
	for i, v := range bb {
		if v == '_' {
			isTitle = true
			if i > 0 {
				b := buf[i-1]
				isOneKey = b >= 'A' && b <= 'Z'
			}
			continue
		}
		if isTitle {
			if v >= 'a' && v <= 'z' {
				v -= 32
			}
			isTitle = false
		} else if isOneKey && v >= 'A' && v <= 'Z' {
			v += 32
		}
		buf[index] = v
		index++
	}
	s := utils.BytesToString(buf[:index])
	return s
}

// PascalCaseToUnderline 帕斯卡命名转下划线命名
func PascalCaseToUnderline(name string) string {
	if name == "" {
		return ""
	}

	var (
		value                                    = commonInitialismsReplacer.Replace(name)
		buf                                      = bytes.NewBufferString("")
		lastCase, currCase, nextCase, nextNumber bool
		bb                                       = utils.StringToBytes(value)
		vLen                                     = len(bb) - 1
	)

	for i := 0; i < vLen; i++ {
		v := bb[i]
		nextV := bb[i+1]
		nextCase = nextV >= 'A' && nextV <= 'Z'
		nextNumber = nextV >= '0' && nextV <= '9'

		if i > 0 {
			if currCase {
				if lastCase && (nextCase || nextNumber) {
					buf.WriteByte(v)
				} else {
					if bb[i-1] != '_' && nextV != '_' {
						buf.WriteByte('_')
					}
					buf.WriteByte(v)
				}
			} else {
				buf.WriteByte(v)
				if i == vLen-1 && (nextCase && !nextNumber) {
					buf.WriteByte('_')
				}
			}
		} else {
			currCase = true
			buf.WriteByte(v)
		}
		lastCase = currCase
		currCase = nextCase
	}

	buf.WriteByte(bb[vLen])
	s := strings.ToLower(utils.BytesToString(buf.Bytes()))
	return s
}
