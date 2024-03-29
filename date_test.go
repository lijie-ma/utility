package utility

import (
	"testing"
)

func TestDate(t *testing.T) {
	d := Date(YYYY_MM_DD)
	t.Log(d)
	t.Log(Date(YYYY_MM_DD_H_I_S))
}

func TestDate2Unix(t *testing.T) {
	d := Date2Unix(Date(YYYY_MM_DD_H_I_S))
	t.Log(d)
	t.Log(Unix2Time(d))
}

func TestFutureDateFromDay(t *testing.T) {
	t.Log(FutureDateFromDay(Date(YYYY_MM_DD), 24))
	t.Log(FutureDateFromDay(Date(YYYY_MM_DD), -24))
}

func TestCompare(t *testing.T) {
	t.Log("lt", DateCompare("2019-07-01", "2019-07-02"))
	t.Log("eq", DateCompare("2019-07-01", "2019-07-01"))
	t.Log("-2", DateCompare("2019-07-100", "2019-07-01"))
}
