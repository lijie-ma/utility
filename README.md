# utility
date util

## 安装
```
go get -u github.com/lijie-ma/utility
```

## demo

```golang
#当前日期
today := utility.Date() // 2019-07-02

#以给定时间计算未来的时间 (小时)
tomorrow, err := utility.FutureDateFromDay(today, 24)

yestoday, err := utility.FutureDateTimeFromDay(Date(YYYY_MM_DD_H_I_S), -24*3600)

// 将日期转化为时间戳 （秒）
unix := utility.Date2Unix(today)
// 每月的最后一天
days := utility.MonthLastDay(today)

#日期比较
switch Compare("2019-07-01", "2019-07-01") {
    case 0:
        fmt.Println("eq")
    case -1:
        fmt.Println("lt")
    case 1:
        fmt.Println("gt")
    case -2:
        fmt.Println("error")
}

```