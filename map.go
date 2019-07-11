package utility

import (
	"errors"
	"reflect"
)

//MapKeys 返回map中的keys
func MapKeys(maps interface{}, filter ...func(key interface{}) bool) (interface{}, error) {
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
	if 0 < len(filter) {
		for _, v := range keys {
			if !filter[0](v.Interface()) {
				newSlice = reflect.Append(newSlice, v)
			}
		}
	} else {
		for _, v := range keys {
			newSlice = reflect.Append(newSlice, v)
		}
	}

	return newSlice.Interface(), nil
}
