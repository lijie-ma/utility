# utility
常用函数工具包

## 安装
```
go get -u github.com/lijie-ma/utility
```

## demo

```golang
func init() {
	SetTimeZone(`UTC`) // 默认 Asia/chongqing
}

func main() {
	today := utility.Date()
	after24Hours, _ := utility.FutureDateFromDay(today, 24)

	dateTime := utilit.Date(YYYY_MM_DD_H_I_S) // 2019-07-01 12:12:11

	after5Min, _ := utility.FutureDateTimeFromDay(dateTime, 5*60)

	fmt.Println(after24Hours, after5Min)

	//日期比较
	switch utility.Compare(today, after24Hours) {
	case 0:
		fmt.Println("eq")
	case -1:
		fmt.Println("lt")
	case 1:
		fmt.Println("gt")
	case -2:
		fmt.Println("error")
	}
}


```