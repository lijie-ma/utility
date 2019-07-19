package utility

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func Substr(s string, start int, length ...int) string {
	fetchLen := len(s)
	if len(length) > 0 {
		fetchLen = length[0]
	}
	runeSlice := []rune(s)
	lenSlice := len(runeSlice)

	if start < 0 {
		start = lenSlice + start
	}
	if start > lenSlice {
		start = lenSlice
	}
	end := start + fetchLen
	if end > lenSlice {
		end = lenSlice
	}
	if fetchLen < 0 {
		end = lenSlice + fetchLen
	}
	if start > end {
		start, end = end, start
	}
	return string(runeSlice[start:end])
}

//HttpBuildQuery 生成 URL-encode 之后的请求字符串
func HttpBuildQuery(params map[string]interface{}) string {
	var urlValue = url.Values{}
	for k, v := range params {
		switch v.(type) {
		case float32:
			urlValue.Add(k, fmt.Sprintf("%f", v.(float32)))
		case []float32:
			for _, sv := range v.([]float32) {
				urlValue.Add(k+`[]`, fmt.Sprintf("%f", sv))
			}
		case float64:
			urlValue.Add(k, fmt.Sprintf("%f", v.(float64)))
		case []float64:
			for _, sv := range v.([]float64) {
				urlValue.Add(k+`[]`, fmt.Sprintf("%f", sv))
			}
		case int:
			urlValue.Add(k, strconv.Itoa(v.(int)))
		case []int:
			for _, sv := range v.([]int) {
				urlValue.Add(k+`[]`, strconv.Itoa(sv))
			}
		case int64:
			urlValue.Add(k, strconv.FormatInt(v.(int64), 10))
		case []int64:
			for _, sv := range v.([]int64) {
				urlValue.Add(k+`[]`, strconv.FormatInt(sv, 10))
			}
		case string:
			urlValue.Add(k, v.(string))
		case []string:
			for _, sv := range v.([]string) {
				urlValue.Add(k+`[]`, sv)
			}
		default:
			panic(k + " invalid type " + reflect.TypeOf(v).Kind().String())
		}
	}
	return urlValue.Encode()
}

func Ucfirst(s string) string {
	for _, v := range s {
		if 'a' <= v && 'z' >= v {
			return string(v-32) + s[1:]
		}
		return s
	}
	return s
}

func Lcfirst(s string) string {
	for _, v := range s {
		if 'A' <= v && 'Z' >= v {
			return string(v+32) + s[1:]
		}
		return s
	}
	return s
}
