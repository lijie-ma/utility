package utility

import (
	"reflect"
)

//判断 val 是否在 array 里
// 返回 exists（bool）  index（int）
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func ArrayUnique(s []string) []string {
	ret := s[:0]
	assist := map[string]struct{}{}
	for _, v := range s {
		if _, ok := assist[v]; !ok {
			assist[v] = struct{}{}
			ret = append(ret, v)
		}
	}
	return ret
}

//数组过滤，如果传递的不是数组，则返回原输入
func ArrayFilter(array interface{}, filter func(value interface{}) bool) interface{} {
	tmpType := reflect.TypeOf(array)
	if tmpType.Kind() != reflect.Slice {
		return array
	}

	s := reflect.ValueOf(array)
	newSlice := reflect.MakeSlice(tmpType, 0, s.Cap())
	for i:=0; i < s.Len(); i++ {
		if !filter(s.Index(i).Interface()) {
			newSlice = reflect.Append(newSlice, s.Index(i))
		}
	}
	return newSlice
}

func ArrayPop(arrayPoint interface{}) interface{} {
	tmpType := reflect.TypeOf(arrayPoint)
	if tmpType.Kind() != reflect.Ptr {
		panic("需要传递数组指针")
	}
	v := reflect.ValueOf(arrayPoint)
	elem := v.Elem()
	if elem.Kind() != reflect.Slice {
		panic("需要传递数组指针")
	}
	elem.Set(reflect.AppendSlice(elem.Slice(0, elem.Len()-1), elem.Slice(0, 0)))
	return elem
}