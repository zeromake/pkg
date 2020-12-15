package reflect

import (
	"fmt"
	"reflect"
)

// CopyToStruct 结构体拷贝
func CopyToStruct(dst interface{}, src interface{}) error {
	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.Indirect(reflect.ValueOf(src))
	if dstValue.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must is a ptr")
	}
	dstValue = dstValue.Elem()
	dstType := dstValue.Type()
	for i := 0; i < dstValue.NumField(); i++ {
		field := dstValue.Field(i)
		if field.CanSet() {
			fieldName := dstType.Field(i)
			s := srcValue.FieldByName(fieldName.Name)
			if s.Kind() == field.Kind() {
				field.Set(s)
			}
		}
	}
	return nil
}
