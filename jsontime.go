package types

import "time"

type JsonTime time.Time

//实现它的json序列化方法
func (this JsonTime) MarshalBinary() ([]byte, error) {
	return this.MarshalText()
} //实现它的json序列化方法
func (this *JsonTime) UnmarshalBinary(bs []byte) (err error) {
	err = this.UnmarshalText(bs)
	return
}

//实现它的json序列化方法
func (this JsonTime) MarshalText() ([]byte, error) {
	var stamp = time.Time(this).Format(CommonDatetime)
	return []byte(stamp), nil
}

//实现它的json序列化方法
func (this *JsonTime) UnmarshalText(bs []byte) (err error) {
	var t time.Time
	t, err = time.Parse(CommonDatetime, string(bs))
	jt := JsonTime(t)
	this = &jt
	return
}
