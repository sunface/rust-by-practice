package talent

import (
	"errors"
	"time"
)

func Time2String(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.999")
}

//将"2016-02-15 12:00:00"或者"2016-04-18 09:33:56.694"等格式转化为time.Time
func StringToTime(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	return t, err
}

//将"2016-04-22T21:47:49.694123232+08:00"或者"2016-04-22T21:47:49+08:00"等格式转化为time.Time
func StringToTime1(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", s, loc)
	return t, err
}

func NSToTime(ns int64) (time.Time, error) {
	if ns <= 0 {
		return time.Time{}, errors.New("ns is err")
	}
	sec := ns / 1e9
	nsec := ns - sec*1e9
	return time.Unix(sec, nsec), nil
}
