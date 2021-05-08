package strutil

import (
	"unsafe"
)

// BytesToString 字节数组转字符串
func BytesToString(b []byte) string {
	// bytes 转字符串可以直接复用指针
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes 字符串转字节数组
//func StringToBytes(s string) []byte {
//	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
//	bh := reflect.SliceHeader{
//		Data: sh.Data,
//		Len:  sh.Len,
//		Cap:  sh.Len,
//	}
//	return *(*[]byte)(unsafe.Pointer(&bh))
//}

// StringToBytes 字符串转字节数组
func StringToBytes(s string) []byte {
	sh := (*[2]uintptr)(unsafe.Pointer(&s))
	// 使用 []uintptr 比 reflect.SliceHeader: 0.3166 ns/op -> 0.3088 ns/op
	bh := [3]uintptr{
		sh[0],
		sh[1],
		sh[1],
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
