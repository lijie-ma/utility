//package utility 旨在提供便捷的函数操作，其中大部分的函数是源自于php函数
package utility

import (
	"fmt"
	"time"
)

const (
	//日期格式
	YYYY_MM_DD       = `2006-01-02`
	YYYYMMDD         = `20060102`
	YYYY_MM_DD_H_I_S = `2006-01-02 15:04:05`

	//默认时区
	DefalutTimeZone = `Asia/Chongqing`
)

var timeZone string

func init() {
	timeZone = DefalutTimeZone
}

//SetTimeZone 设置时区
func SetTimeZone(zone string) {
	timeZone = zone
}

func getLocation() *time.Location {
	l, err := time.LoadLocation(timeZone)
	if nil != err {
		panic(err)
	}
	return l
}

// unix 时间戳
func Time() int64 {
	return LocalTime().Unix()
}

//当前时间
func Date(style ...string) string {
	defaultStyle := YYYY_MM_DD
	if 0 < len(style) {
		defaultStyle = style[0]
	}
	return LocalTime().Format(defaultStyle)
}

//LocalTime 返回默认时区的time.Time
func LocalTime() time.Time {
	return time.Now().In(getLocation())
}

//Str2Time 字符串转换unix 时间戳
// date  格式 2019-07-01 20190701 2019-07-01 12:12:12
// 返回 0 代表失败
func Str2Time(date string) int64 {
	return Date2Unix(date)
}

//过往时间， 默认取 5分钟前的时间
func LastTime(style string, agoSeconds ...int) string {
	if 0 == len(style) {
		style = YYYY_MM_DD_H_I_S
	}
	s := -5 * 60
	if 0 < len(agoSeconds) {
		s = 0 - agoSeconds[0]
	}
	now := LocalTime()
	lastDay := now.Add(time.Duration(s) * time.Second)
	return lastDay.Format(style)
}

func FutureDateFromDay(date string, hours int) (string, error) {
	return addTime(date, YYYY_MM_DD, hours*3600)
}

func FutureDateTimeFromDay(date string, seconds int) (string, error) {
	return addTime(date, YYYY_MM_DD_H_I_S, seconds)
}

//计算时间差
// seconds 是相隔时间的秒数
// isFuture 标记是将来的时间，还是过去的时间， 默认过去的时间
func addTime(times string, timeStyle string, seconds int) (string, error) {
	t, err := parseTime(times, timeStyle)
	if nil != err {
		return "", err
	}
	l := t.Add(time.Duration(seconds) * time.Second)
	return l.Format(timeStyle), nil
}

//返回时区为 Asia/Chongqing 的unix 时间戳
func Date2Unix(date string) int64 {
	t, err := parseTime(date)
	if nil != err {
		return 0
	}
	return t.Unix()
}

//DateCompare 日期比较
// 相等返回 0， 小于 -1 大于 1
// -2 代表异常
func DateCompare(t1, t2 string, style ...string) int {
	tt1, er := parseTime(t1, style...)
	if nil != er {
		return -2
	}
	tt2, err := parseTime(t2, style...)
	if nil != err {
		return -2
	}
	if tt1.Equal(tt2) {
		return 0
	}

	if tt1.Before(tt2) {
		return -1
	}
	return 1
}

func parseTime(t1 string, style ...string) (time.Time, error) {
	s := YYYY_MM_DD
	if 0 < len(style) {
		s = style[0]
	} else {
		switch len(t1) {
		case 8:
			s = YYYYMMDD
		case 19:
			s = YYYY_MM_DD_H_I_S
		}
	}
	return time.ParseInLocation(s, t1, getLocation())
}

func Unix2Time(unixTime int64, style ...string) string {
	t := time.Unix(unixTime, 0)
	if 0 == len(style) {
		return t.Format(YYYY_MM_DD_H_I_S)
	}
	return t.Format(style[0])
}

//MonthLastDay 返回一月的最后一天的
// 返回格式 20190731 或者 2019-07-31 （依据date的传入格式）
func MonthLastDay(date string) string {
	var year, month int
	byteSlice := []byte(date)
	style := `%d%d%d`
	switch len(date) {
	case 8:
		year = Atoi(string(byteSlice[0:4]))
		month = Atoi(string(byteSlice[4:6]))
	case 10, 19:
		year = Atoi(string(byteSlice[0:4]))
		month = Atoi(string(byteSlice[5:7]))
		style = `%d-%d-%d`
	default:
		panic("date is not a valid time format")
	}
	return fmt.Sprintf(style, year, month, computeMonthDays(year, month))
}

//MonthOfDays 返回当前月份的天数
func MonthOfDays(date string) int {
	var year, month int
	byteSlice := []byte(date)
	switch len(date) {
	case 8:
		year = Atoi(string(byteSlice[0:4]))
		month = Atoi(string(byteSlice[4:6]))
	case 10, 19:
		year = Atoi(string(byteSlice[0:4]))
		month = Atoi(string(byteSlice[5:7]))
	default:
		panic("date is not a valid time format")
	}
	return computeMonthDays(year, month)
}

func computeMonthDays(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}
