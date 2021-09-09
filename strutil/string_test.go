package strutil

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

var testString = "11111"

func TestStringToBytes(t *testing.T) {
	p1 := (*reflect.StringHeader)(unsafe.Pointer(&testString))
	bb := StringToBytes(testString)
	p2 := (*reflect.SliceHeader)(unsafe.Pointer(&bb))
	assert.Equal(t, p1.Data, p2.Data)
	assert.Equal(t, p1.Len, p2.Len)
}

var testBytes = []byte{
	'1',
	'1',
	'1',
	'1',
}

func TestBytesToString(t *testing.T) {
	p1 := (*reflect.SliceHeader)(unsafe.Pointer(&testBytes))
	ss := BytesToString(testBytes)
	p2 := (*reflect.StringHeader)(unsafe.Pointer(&ss))
	assert.Equal(t, p1.Data, p2.Data)
	assert.Equal(t, p1.Len, p2.Len)
}

func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StringToBytes(testString)
	}
}

// func BenchmarkStringToBytes2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_ = StringToBytes2(testString)
// 	}
// }
