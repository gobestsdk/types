package types

const (
	DateFormat     = "20060102"
	CommonDatetime = "2006-01-02 15:04:05"
	CommonDate     = "2006-01-02"
	SimpleDate     = "2006-1-2"
	RFC3339Milli   = "2006-01-02T15:04:05.999999Z07:00"
)

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day() + 1)
	return GetZeroTime(d)
}
//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}
 
//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}