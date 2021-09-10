package patch

import (
	"errors"
	"github.com/zeromake/pkg/patch/types"
	"reflect"
	"strings"
)

func ModifyPatchSlice(
	rValue reflect.Value,
	path string,
	v interface{},
	opt *types.Option,
) (result reflect.Value, err error) {
	rType := rValue.Type()
	key, nextPath := types.Split(path)
	elementType := rType.Elem()
	index := int(types.Number(key).Int64())
	has := index < rValue.Len()
	var old reflect.Value
	if has {
		old = rValue.Index(index)
	}
	if nextPath != "" {
		var element reflect.Value
		if !has {
			if elementType.Kind() == reflect.Ptr {
				element = reflect.New(elementType.Elem())
			} else {
				element = reflect.New(elementType).Elem()
			}
			old = element
		}
		if elementType.Kind() == reflect.Slice {
			result, err = ModifyPatchSlice(old, nextPath, v, opt)
			if err != nil {
				return
			}
			old.Set(result)
		} else {
			err = ModifyPatch(old, nextPath, v, opt)
			if err != nil {
				return
			}
		}
		if !has {
			rValue = reflect.Append(rValue, old)
		}
		return rValue, nil
	}
	if opt.Replace {
		if !has {
			if elementType.Kind() == reflect.Ptr {
				old = reflect.New(elementType.Elem())
			} else {
				old = reflect.New(elementType).Elem()
			}
			rValue = reflect.Append(rValue, old)
		}
		err = SetValue(v, old)
		if err != nil {
			return
		}
		return
	}
	if elementType.Kind() == reflect.Ptr {
		old = reflect.New(elementType.Elem())
	} else {
		old = reflect.New(elementType).Elem()
	}
	err = SetValue(v, old)
	if err != nil {
		return
	}
	rValue = reflect.Append(rValue, old)
	return rValue, nil
}

func IncrPatch(rValue reflect.Value, path string, incr int64) (err error) {
	if incr == 0 {
		return
	}
	pp, ok := rValue.Interface().(types.Patch)
	if ok {
		p := &types.Item{
			Op:    "incr",
			Path:  path,
			Value: incr,
		}
		return pp.Patch(p)
	}
	rType := rValue.Type()
	key, nextPath := types.Split(path)
	switch rType.Kind() {
	case reflect.Map:
		keyType := rType.Key()
		elementType := rType.Elem()
		if rValue.IsZero() {
			rValue.Set(reflect.MakeMap(rType))
		}
		keyRvalue := reflect.New(keyType).Elem()
		err = SetValue(key, keyRvalue)
		if err != nil {
			return
		}
		mapElement := rValue.MapIndex(keyRvalue)
		has := mapElement.Kind() != reflect.Invalid
		if nextPath != "" {
			if !has {
				if elementType.Kind() == reflect.Ptr {
					mapElement = reflect.New(elementType.Elem())
				} else {
					mapElement = reflect.New(elementType).Elem()
				}
				rValue.SetMapIndex(keyRvalue, mapElement)
			}
			return IncrPatch(mapElement, nextPath, incr)
		}
		if has {
			err = IncrNumber(mapElement, incr)
			if err != nil {
				return
			}
		} else {
			mapElement = reflect.New(elementType).Elem()
			err = IncrNumber(mapElement, incr)
			if err != nil {
				return
			}
			rValue.SetMapIndex(keyRvalue, mapElement)
		}
	case reflect.Slice:
		index := int(types.Number(key).Int64())
		has := index < rValue.Len()
		var old reflect.Value
		if !has {
			return
		}
		old = rValue.Index(index)
		if nextPath != "" {
			return IncrPatch(old, nextPath, incr)
		}
		return IncrNumber(old, incr)
	case reflect.Ptr:
		if rValue.IsZero() {
			rValue.Set(reflect.New(rType.Elem()))
		}
		return IncrPatch(rValue.Elem(), path, incr)
	case reflect.Struct:
		var index = 0
		for ; index < rType.NumField(); index++ {
			field := rType.Field(index)
			tag := field.Tag.Get("json")
			if tag == "-" {
				continue
			}
			tag = strings.SplitN(tag, ",", 2)[0]
			if tag == key {
				break
			}
		}
		if index >= rType.NumField() {
			return
		}
		vv := rValue.Field(index)
		has := !vv.IsZero()
		if nextPath != "" {
			var nn reflect.Value
			var elemType reflect.Type
			if vv.Kind() == reflect.Ptr {
				elemType = vv.Type().Elem()
				if !has {
					nn = reflect.New(elemType)
				}
			} else {
				elemType = vv.Type()
				if !has {
					nn = reflect.New(elemType).Elem()
				}
			}
			if !has {
				vv.Set(nn)
			}
			return IncrPatch(vv, nextPath, incr)
		}
		return IncrNumber(vv, incr)
	}
	return
}

func ModifyPatch(rValue reflect.Value, path string, v interface{}, opt *types.Option) (err error) {
	pp, ok := rValue.Interface().(types.Patch)
	if ok {
		p := &types.Item{
			Op:    "add",
			Path:  path,
			Value: v,
		}
		if opt.Replace {
			p.Op = "replace"
		}
		return pp.Patch(p)
	}

	rType := rValue.Type()
	key, nextPath := types.Split(path)
	switch rType.Kind() {
	case reflect.Map:
		keyType := rType.Key()
		elementType := rType.Elem()
		if rValue.IsZero() {
			rValue.Set(reflect.MakeMap(rType))
		}
		keyRvalue := reflect.New(keyType).Elem()
		err = SetValue(key, keyRvalue)
		if err != nil {
			return
		}
		mapElement := rValue.MapIndex(keyRvalue)
		has := mapElement.Kind() != reflect.Invalid
		if nextPath != "" {
			if !has {
				if elementType.Kind() == reflect.Ptr {
					mapElement = reflect.New(elementType.Elem())
				} else {
					mapElement = reflect.New(elementType).Elem()
				}
				rValue.SetMapIndex(keyRvalue, mapElement)
			}
			if elementType.Kind() == reflect.Slice {
				var result reflect.Value
				result, err = ModifyPatchSlice(mapElement, nextPath, v, opt)
				if err != nil {
					return
				}
				mapElement.Set(result)
			}
			return ModifyPatch(mapElement, nextPath, v, opt)
		}
		if !opt.Replace && has {
			return
		}
		mapElement = reflect.New(elementType).Elem()
		err = SetValue(v, mapElement)
		if err != nil {
			return
		}
		rValue.SetMapIndex(keyRvalue, mapElement)
	case reflect.Struct:
		var result reflect.Value
		var index = 0
		for ; index < rType.NumField(); index++ {
			field := rType.Field(index)
			tag := field.Tag.Get("json")
			if tag == "-" {
				continue
			}
			tag = strings.SplitN(tag, ",", 2)[0]
			if tag == key {
				break
			}
		}
		if index >= rType.NumField() {
			return
		}
		vv := rValue.Field(index)
		has := !vv.IsZero()
		if nextPath != "" {
			var nn reflect.Value
			var elemType reflect.Type
			if vv.Kind() == reflect.Ptr {
				elemType = vv.Type().Elem()
				if !has {
					nn = reflect.New(elemType)
				}
			} else {
				elemType = vv.Type()
				if !has {
					nn = reflect.New(elemType).Elem()
				}
			}
			if !has {
				vv.Set(nn)
			}
			if elemType.Kind() == reflect.Slice {
				if vv.Kind() == reflect.Ptr {
					result, err = ModifyPatchSlice(vv.Elem(), nextPath, v, opt)
					if err != nil {
						return err
					}
					vv.Set(result.Addr())
					return
				} else {
					result, err = ModifyPatchSlice(vv, nextPath, v, opt)
					if err != nil {
						return err
					}
					vv.Set(result)
					return
				}
			}
			return ModifyPatch(vv, nextPath, v, opt)
		}
		if !opt.Replace && has {
			return
		}
		vv = rValue.Field(index)
		if vv.Kind() == reflect.Invalid {
			if vv.Kind() == reflect.Ptr {
				vv.Set(reflect.New(vv.Type().Elem()))
			} else {
				vv.Set(reflect.New(vv.Type()).Elem())
			}
		}
		err = SetValue(v, vv)
		return
	case reflect.Ptr:
		var result reflect.Value
		elem := rValue.Elem()
		if elem.Kind() == reflect.Slice {
			result, err = ModifyPatchSlice(elem, path, v, opt)
			if err != nil {
				return err
			}
			rValue.Set(result.Addr())
		}
		return ModifyPatch(elem, path, v, opt)
	}
	return
}

func RemovePatch(rValue reflect.Value, path string) (err error) {
	pp, ok := rValue.Interface().(types.Patch)
	if ok {
		p := &types.Item{
			Op:   "remove",
			Path: path,
		}
		return pp.Patch(p)
	}
	rType := rValue.Type()
	key, nextPath := types.Split(path)
	switch rType.Kind() {
	case reflect.Ptr:
		if key == "" {
			rValue.Set(reflect.Zero(rType))
			return
		}
		return RemovePatch(rValue.Elem(), path)
	case reflect.Struct:
		var index = 0
		for ; index < rType.NumField(); index++ {
			field := rType.Field(index)
			tag := field.Tag.Get("json")
			if tag == "-" {
				continue
			}
			tag = strings.SplitN(tag, ",", 2)[0]
			if tag == key {
				break
			}
		}
		if index >= rType.NumField() {
			return
		}
		field := rValue.Field(index)
		if field.IsZero() {
			return
		}
		if nextPath != "" {
			return RemovePatch(field, nextPath)
		}
		field.Set(reflect.Zero(field.Type()))
	case reflect.Map:
		k := reflect.ValueOf(key)
		element := rValue.MapIndex(k)
		has := element.Kind() != reflect.Invalid
		if !has {
			return
		}
		if nextPath != "" {
			return RemovePatch(element, nextPath)
		}
		rValue.SetMapIndex(k, reflect.Value{})
	case reflect.Slice:
		count := rValue.Len()
		index := int(types.ToInt64(key))
		if index >= count {
			return
		}
		sliceItem := rValue.Index(index)
		if nextPath != "" {
			return RemovePatch(sliceItem, nextPath)
		}
		result := reflect.AppendSlice(rValue.Slice(0, index), rValue.Slice(index+1, count))
		rValue.Set(result)
	}
	return
}

func IncrNumber(dst reflect.Value, incr int64) (err error) {
	switch dst.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		old := dst.Int()
		old += incr
		dst.SetInt(old)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		old := dst.Uint()
		old += uint64(incr)
		dst.SetUint(old)
	case reflect.Float32, reflect.Float64:
		old := dst.Float()
		old += float64(incr)
		dst.SetFloat(old)
	case reflect.Ptr:
		dst.Set(reflect.New(dst.Type().Elem()))
		return IncrNumber(dst.Elem(), incr)
	}
	return
}

func SetValue(src interface{}, dst reflect.Value) (err error) {
	switch dst.Kind() {
	case reflect.Bool: // true, false
		value, _ := src.(bool)
		dst.SetBool(value)
	case reflect.String: // string
		ss, _ := src.(string)
		dst.SetString(ss)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst.SetInt(types.ToInt64(src))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		dst.SetUint(types.ToUint64(src))
	case reflect.Float32, reflect.Float64:
		dst.SetFloat(types.ToFloat64(src))
	case reflect.Ptr:
		return SetValue(src, dst.Elem())
	case reflect.Slice:
		rSrc := reflect.ValueOf(src)
		rSrcType := rSrc.Type()
		dstType := dst.Type()
		if rSrcType.Kind() != reflect.Slice {
			return errors.New("need slice -> slice")
		}
		result := reflect.MakeSlice(dstType, 0, rSrc.Len())
		for i := 0; i < rSrc.Len(); i++ {
			item := rSrc.Index(i)
			e := reflect.New(dstType.Elem()).Elem()
			err = SetValue(item.Interface(), e)
			if err != nil {
				return
			}
			result = reflect.Append(result, e)
		}
		dst.Set(result)
	case reflect.Struct:
		rSrc := reflect.ValueOf(src)
		rSrcType := rSrc.Type()
		if rSrcType.Kind() != reflect.Map {
			return errors.New("need map -> struct")
		}
		iter := rSrc.MapRange()
		dstType := dst.Type()
		cache := map[string]int{}
		for i := 0; i < dstType.NumField(); i++ {
			field := dstType.Field(i)
			tag := field.Tag.Get("json")
			tag = strings.SplitN(tag, ",", 2)[0]
			if tag != "" && tag != "-" {
				cache[tag] = i
			}
		}
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()
			if key.Kind() == reflect.Interface {
				key = key.Elem()
			}
			k := key.String()
			index, has := cache[k]
			if has {
				err = SetValue(value.Interface(), dst.Field(index))
				if err != nil {
					return
				}
			}
		}
	}
	return nil
}

var addOption = &types.Option{}
var replaceOption = &types.Option{
	Replace: true,
}

func ApplyPatch(value interface{}, patch *types.Item) (err error) {
	rValue := reflect.ValueOf(value)
	switch patch.Op {
	case "add":
		err = ModifyPatch(rValue, patch.Path, patch.Value, addOption)
	case "replace":
		err = ModifyPatch(rValue, patch.Path, patch.Value, replaceOption)
	case "remove":
		err = RemovePatch(rValue, patch.Path)
	case "incr":
		err = IncrPatch(rValue, patch.Path, types.ToInt64(patch.Value))
	}
	if err != nil {
		return
	}
	return
}
