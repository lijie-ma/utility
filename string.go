package utility

import (
	"fmt"
	"net/url"
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
		case float64:
			urlValue.Add(k, fmt.Sprintf("%f", v.(float64)))
		case int:
			urlValue.Add(k, fmt.Sprintf("%d", v.(int)))
		case []int:
			for _, sv := range v.([]int) {
				urlValue.Add(k+`[]`, fmt.Sprintf("%d", sv))
			}
		case int64:
			urlValue.Add(k, fmt.Sprintf("%d", v.(int64)))
		case []int64:
			for _, sv := range v.([]int64) {
				urlValue.Add(k+`[]`, fmt.Sprintf("%d", sv))
			}
		case string:
			urlValue.Add(k, v.(string))
		case []string:
			for _, sv := range v.([]string) {
				urlValue.Add(k+`[]`, sv)
			}
		default:
			panic("invalid type")
		}
	}
	return urlValue.Encode()
}
