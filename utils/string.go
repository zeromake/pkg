package utils

import (
	"reflect"
	"unsafe"
)

// BytesToString 字节数组转字符串
func BytesToString(b []byte) string {
	// bytes 转字符串可以直接复用指针
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes 字符串转字节数组
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
