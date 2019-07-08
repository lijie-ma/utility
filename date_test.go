package utility

import "testing"


func TestDate(t *testing.T) {
	d := Date()
	t.Log(d)
}

func TestDate2Unix(t *testing.T) {
	d := Date2Unix(DateTime())
	t.Log( d)
	t.Log(Unix2Time(d))
}

func TestFutureDateFromDay(t *testing.T) {
	t.Log(FutureDateFromDay(Date(), 24))
	t.Log(FutureDateFromDay(Date(), -24))
}

func TestCompare(t *testing.T) {
	t.Log("lt", Compare("2019-07-01", "2019-07-02"))
	t.Log("eq",Compare("2019-07-01", "2019-07-01"))
	t.Log("-2",Compare("2019-07-100", "2019-07-01"))
}