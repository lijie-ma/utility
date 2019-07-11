package utility

import (
	"errors"
	"reflect"
)

//MapKeys 返回map中的keys
func MapKeys(maps interface{}) (interface{}, error) {
	tmap := reflect.TypeOf(maps)
	if tmap.Kind() != reflect.Map {
		return nil, errors.New("maps is not a map")
	}
	vMap := reflect.ValueOf(maps)
	keys := vMap.MapKeys()
	if 0 == len(keys) {
		return nil, errors.New("maps's lenght is zero")
	}
	newSlice := reflect.MakeSlice(reflect.SliceOf(keys[0].Type()), 0, len(keys))
	for _, v := range keys {
		newSlice = reflect.Append(newSlice, v)
	}

	return newSlice.Interface(), nil
}
