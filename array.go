package utility

import (
	"errors"
	"math"
	"reflect"
)

//判断 val 是否在 array 里
// 返回 exists（bool）  index（int）
func InSlice(val interface{}, array interface{}) (exists bool, index int) {
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

func SliceUnique(s []string) []string {
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
func SliceFilter(array interface{}, filter func(value interface{}) bool) interface{} {
	tmpType := reflect.TypeOf(array)
	if tmpType.Kind() != reflect.Slice {
		return array
	}

	s := reflect.ValueOf(array)
	newSlice := reflect.MakeSlice(tmpType, 0, s.Cap())
	for i := 0; i < s.Len(); i++ {
		if !filter(s.Index(i).Interface()) {
			newSlice = reflect.Append(newSlice, s.Index(i))
		}
	}
	return newSlice.Interface()
}

func SlicePop(arrayPoint interface{}) {
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
}

//将数组开头的单元移出数组
func SliceShift(arrayPoint interface{}) {
	tmpType := reflect.TypeOf(arrayPoint)
	if tmpType.Kind() != reflect.Ptr {
		panic("需要传递数组指针")
	}
	v := reflect.ValueOf(arrayPoint)
	elem := v.Elem()
	if elem.Kind() != reflect.Slice {
		panic("需要传递数组指针")
	}
	elem.Set(reflect.AppendSlice(elem.Slice(1, elem.Len()), elem.Slice(1, 1)))
}

// SliceIntersect 数组交集
func SliceIntersect(a1, a2 interface{}) interface{} {
	t1 := reflect.TypeOf(a1)
	if t1.Kind() != reflect.Slice {
		panic("a1 is not a slice")
	}
	t2 := reflect.TypeOf(a2)
	if t2.Kind() != reflect.Slice {
		panic("a2 is not a slice")
	}

	v1 := reflect.ValueOf(a1)
	v2 := reflect.ValueOf(a2)
	if v1.Type().String() != v2.Type().String() {
		panic("a1 and a2 must be the same type of slice")
	}
	tmp := make(map[interface{}]interface{})
	for i := 0; i < v1.Len(); i++ {
		tmp[v1.Index(i).Interface()] = struct{}{}
	}
	newSlice := reflect.MakeSlice(v1.Type(), 0, v1.Len())
	for i := 0; i < v2.Len(); i++ {
		if _, ok := tmp[v2.Index(i).Interface()]; ok {
			newSlice = reflect.Append(newSlice, v2.Index(i))
		}
	}
	return newSlice.Interface()
}

func SliceDiff(a1, a2 interface{}) interface{} {
	t1 := reflect.TypeOf(a1)
	if t1.Kind() != reflect.Slice {
		panic("a1 is not a slice")
	}
	t2 := reflect.TypeOf(a2)
	if t2.Kind() != reflect.Slice {
		panic("a2 is not a slice")
	}

	v1 := reflect.ValueOf(a1)
	v2 := reflect.ValueOf(a2)
	if v1.Type().String() != v2.Type().String() {
		panic("a1 and a2 must be the same type of slice")
	}
	tmp := make(map[interface{}]interface{})
	for i := 0; i < v2.Len(); i++ {
		tmp[v2.Index(i).Interface()] = struct{}{}
	}
	newSlice := reflect.MakeSlice(v1.Type(), 0, v1.Len())
	for i := 0; i < v1.Len(); i++ {
		if _, ok := tmp[v1.Index(i).Interface()]; !ok {
			newSlice = reflect.Append(newSlice, v1.Index(i))
		}
	}
	return newSlice.Interface()
}

//返回数组中指定的一列
func SliceColumn(array interface{}, columnKey string) (interface{}, error) {
	t1 := reflect.TypeOf(array)
	if t1.Kind() != reflect.Slice || t1.Elem().Kind() != reflect.Map {
		return nil, errors.New("array is not a slice")
	}
	vArray := reflect.ValueOf(array)
	vMap := vArray.Index(0)
	vkey := reflect.ValueOf(columnKey)
	tmp1 := vMap.MapIndex(vkey)
	if !tmp1.IsValid() { //columnKey is not exists
		return nil, errors.New(columnKey + " is not exists")
	}
	if tmp1.Kind() == reflect.Interface {
		tmp1 = tmp1.Elem()
	}
	newSlice := reflect.MakeSlice(reflect.SliceOf(tmp1.Type()), 0, vArray.Len())
	newSlice = reflect.Append(newSlice, tmp1)
	for i := 1; i < vArray.Len(); i++ {
		tmp1 := vArray.Index(i).MapIndex(vkey)
		if !tmp1.IsValid() { //columnKey is not exists
			return nil, errors.New(columnKey + " is not exists")
		}
		if tmp1.Kind() == reflect.Interface {
			tmp1 = tmp1.Elem()
		}
		newSlice = reflect.Append(newSlice, tmp1)
	}
	return newSlice.Interface(), nil
}

//计算数组中所有值的乘积
//目前支持 int16，int32， int64， float32，float64
// 返回值类型 为 int64 或float64
func SliceProduct(array interface{}) interface{} {
	figure := func(v1, v2 reflect.Value, kind reflect.Kind) interface{} {
		switch kind {
		case reflect.Int64:
			return v1.Int() * v2.Int()
		case reflect.Float64:
			return v1.Float() * v2.Float()
		}
		return v1.Interface()
	}
	return sliceFigure(array, 1, figure)
}

// 计算数组中所有值的和
// 目前支持 int16，int32， int64， float32，float64
// 返回值类型 为 int64 或float64
func SliceSum(array interface{}) interface{} {
	figure := func(v1, v2 reflect.Value, kind reflect.Kind) interface{} {
		switch kind {
		case reflect.Int64:
			return v1.Int() + v2.Int()
		case reflect.Float64:
			return v1.Float() + v2.Float()
		}
		return v1.Interface()
	}
	return sliceFigure(array, 0, figure)
}

func sliceFigure(array interface{}, initValue int, figure func(v1, v2 reflect.Value, kind reflect.Kind) interface{}) interface{} {
	t1 := reflect.TypeOf(array)
	if t1.Kind() != reflect.Slice {
		return reflect.New(t1).Interface()
	}
	v1 := reflect.ValueOf(array)
	switch v1.Type().Elem().Kind() {
	case reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64,
		reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vSum := reflect.New(reflect.TypeOf(int64(initValue))).Elem()
		vSum.Set(reflect.ValueOf(int64(initValue)))
		for i := 0; i < v1.Len(); i++ {
			vSum.Set(reflect.ValueOf(figure(vSum, v1.Index(i), reflect.Int64).(int64)))
		}
		return vSum.Int()
	case reflect.Float32, reflect.Float64:
		vSum := reflect.New(reflect.TypeOf(float64(initValue))).Elem()
		vSum.Set(reflect.ValueOf(float64(initValue)))
		for i := 0; i < v1.Len(); i++ {
			vSum.Set(reflect.ValueOf(figure(vSum, v1.Index(i), reflect.Float64).(float64)))
		}
		return vSum.Float()
	}
	return reflect.New(t1).Interface()
}

//拆分数组
func SliceChunk(array interface{}, size int) interface{} {
	t1 := reflect.TypeOf(array)
	if t1.Kind() != reflect.Slice {
		return t1
	}
	v1 := reflect.ValueOf(array)
	chunkSize := int(math.Ceil(float64(v1.Len()) / float64(size)))
	if 1 == chunkSize {
		return array
	}
	tempSlice := reflect.MakeSlice(reflect.SliceOf(t1), 0, chunkSize)
	for i := 0; i < chunkSize; i++ {
		end := (i + 1) * size
		if end >= v1.Len() {
			end = v1.Len()
		}
		newSlice := reflect.MakeSlice(t1, 0, size)
		newSlice = reflect.AppendSlice(newSlice.Slice(0, newSlice.Len()), v1.Slice(i*size, end))

		tempSlice = reflect.Append(tempSlice, newSlice)
	}
	return tempSlice.Interface()
}

func SliceWalk(arrayPoint interface{}, call func(value interface{}, index int) interface{}) bool {
	t1 := reflect.TypeOf(arrayPoint)
	if t1.Kind() != reflect.Ptr || t1.Elem().Kind() != reflect.Slice {
		return false
	}
	v1 := reflect.ValueOf(arrayPoint).Elem()
	for i := 0; i < v1.Len(); i++ {
		v1.Index(i).Set(reflect.ValueOf(call(v1.Index(i).Interface(), i)))
	}
	return true

}
