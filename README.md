# utility
date util

## 安装
```
go get -u github.com/lijie-ma/utility
```

## demo

```golang
today := utility.Date() // 2019-07-02

tomorrow, err := utility.FutureDateFromDay(today, 24)

yestoday, err := utility.FutureDateFromDay(today, -24)

// 将日期转化为时间戳 （秒）
unix := utility.Date2Unix(today)

days := utility.MonthLastDay(today)

```