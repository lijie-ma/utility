# utility
go版本的php函数 非大而全，只做常用的和扩展的

doc https://godoc.org/github.com/lijie-ma/utility

## Install
```
go get -u github.com/lijie-ma/utility
```

## Requirements
  Go 1.11 or above


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
	switch utility.DateCompare(today, after24Hours) {
	case 0:
		fmt.Println("eq")
	case -1:
		fmt.Println("lt")
	case 1:
		fmt.Println("gt")
	case -2:
		fmt.Println("error")
	}

	//数组操作
	s := []string{"a", "b", "c", "d", "e"}
	SlicePop(&s)
	fmt.Println(s) // [a c d e]
}


```

### time
```cassandraql
Time()
FutureDateFromDay()
MonthLastDay()
DateCompare()

```


### Slice(Slice/Map) Functions

```function
InSlice()
SliceUnique()
SliceFilter()
SlicePop()
SliceShift()
SliceIntersect()
SliceDiff()
SliceColumn()
SliceSum()
SliceProduct()
SliceChunk()
SliceWalk()

MapKeys()
```

### other

```cassandraql
Ip2Long()
Long2Ip()

Atoi()
AtoInt64()
AtoFloat64()
FloattoStr()
Base64Encode()
Base64Decode()
Md5()
```