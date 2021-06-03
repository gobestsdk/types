package types

import (
	"errors"
	"time"
)

const (
	DateFormat = "20060102"

	CommonDate = "2006-01-02"
	SimpleDate = "2006-1-2"
	LDate      = "2006/1/2"

	CommonDatetime = "2006-01-02 15:04:05"
	RFC3339Milli   = "2006-01-02T15:04:05.999999Z07:00"
)

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

//GetLastZeroTime 获取某一天的23:59:59时间
func GetLastZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}
func MustParseDate(v string) (date time.Time, err error) {
	dates := []string{DateFormat, CommonDatetime, CommonDate, SimpleDate, LDate}
	for _, datepattern := range dates {
		date, err = time.ParseInLocation(datepattern, v, time.Now().Location())
		if err == nil {
			return
		}
	}
	err = errors.New("err date format")
	return
}
