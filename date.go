package utility

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	YYYY_MM_DD       = `2006-01-02`
	YYYYMMDD         = `20060102`
	YYYY_MM_DD_H_I_S = `2006-01-02 15:04:05`
	TIME_LOCAL       = `Asia/Chongqing`
)

//当前时间
func Date(style ...string) string {
	defaultStyle := YYYY_MM_DD
	if 0 < len(style) {
		defaultStyle = style[0]
	}
	return LocalTime().Format(defaultStyle)
}

func DateTime() string {
	return LocalTime().Format(YYYY_MM_DD_H_I_S)
}

func getLocation() *time.Location {
	l, err := time.LoadLocation(TIME_LOCAL)
	if nil != err {
		l = time.FixedZone("CST", 8*3600)
	}
	return l
}

func LocalTime() time.Time {
	return time.Now().In(getLocation())
}

//前一天
func LastDate(style string) string {
	if 0 == len(style) {
		style = YYYY_MM_DD
	}
	now := LocalTime()
	m, _ := time.ParseDuration("-24h")
	lastDay := now.Add(m)
	return lastDay.Format(style)
}

//过往时间， 默认取 5分钟前的时间
func LastTime(style string, agoSeconds ...int) string {
	if 0 == len(style) {
		style = YYYY_MM_DD_H_I_S
	}
	s := -5 * 60
	if 0 < len(agoSeconds) {
		s = -agoSeconds[0]
	}
	now := LocalTime()
	lastDay := now.Add(time.Duration(s) * time.Second)
	return lastDay.Format(style)
}

func FutureDateFromDay(date string, hours int) (string, error) {
	if hours < 0 {
		return addTime(date, YYYY_MM_DD, hours*3600)
	}
	return addTime(date, YYYY_MM_DD, hours*3600, true)
}

func FutureDateTimeFromDay(date string, seconds int) (string, error) {
	if seconds < 0 {
		return addTime(date, YYYY_MM_DD_H_I_S, seconds)
	}
	return addTime(date, YYYY_MM_DD_H_I_S, seconds, true)
}

//计算时间差
// seconds 是相隔时间的秒数
// isFuture 标记是将来的时间，还是过去的时间， 默认过去的时间
func addTime(times string, timeStyle string, seconds int, isFuture ...bool) (string, error) {
	t, err := time.ParseInLocation(timeStyle, times, getLocation())
	if nil != err {
		return "", err
	}
	style := `-%ds`
	if 0 < len(isFuture) && isFuture[0] {
		style = `%ds`
	}
	d, _ := time.ParseDuration(fmt.Sprintf(style, seconds))
	l := t.Add(d)
	return l.Format(timeStyle), nil
}

//返回时区为 Asia/Chongqing 的unix 时间戳
func Date2Unix(date string, style string) int64 {
	t, err := time.ParseInLocation(style, date, getLocation())
	if nil != err {
		return 0
	}
	return t.Unix()
}

func Unix2Time(unixTime int64, style ...string) string {
	t := time.Unix(unixTime, 0)
	if 0 == len(style) {
		return t.Format(YYYY_MM_DD_H_I_S)
	}
	return t.Format(style[0])
}

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func Str2Date(value string) NullTime {
	t, e := time.ParseInLocation(YYYY_MM_DD, value, getLocation())
	if nil != e {
		return NullTime{Valid: false}
	}
	return NullTime{t, true}
}

func Str2Time(value string) NullTime {
	t, e := time.ParseInLocation(YYYY_MM_DD_H_I_S, value, getLocation())
	if nil != e {
		return NullTime{Valid: false}
	}
	return NullTime{t, true}
}

func atoi(arg interface{}) int {
	num, err := strconv.Atoi(arg.(string))
	if nil != err {
		return 0
	}
	return num
}

//一个月的最后一天
// month 格式 2017-09
func MonthLastDay(month string) string {
	s := strings.Split(month, "-")
	days := GetDays(atoi(s[0]), atoi(s[1]))
	return month + fmt.Sprintf("-%d", days)
}

func GetDays(year int, month int) (days int) {
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
