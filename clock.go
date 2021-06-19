package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Clock struct {
	seconds int
}

func (dt Clock) MarshalText() (text []byte, err error) {
	return []byte(dt.String()), nil
}
func (dt *Clock) UnmarshalText(text []byte) (err error) {
	myT, err := ParseDayTime(string(text))
	dt.seconds = myT.seconds
	return
}
func (dt *Clock) Add(s int) *Clock {
	dt.seconds += s
	return dt
}
func (dt Clock) Second() int {
	return dt.seconds
}
func (dt Clock) HMS() (h, m, s int) {
	s = dt.seconds
	h = s / 3600
	s -= h * 3600
	m = s / 60
	s -= m * 60
	return
}
func GetDayTime(t time.Time) (dt Clock) {
	return CreateDayTime(t.Clock())
}
func CreateDayTime(h, m, s int) (dt Clock) {
	dt.seconds += s
	dt.seconds += (m * 60)
	dt.seconds += (h * 60 * 60)
	return
}

// ParseDayTime
func ParseDayTime(dts string) (dt Clock, err error) {
	hms := strings.Split(dts, ":")
	if len(hms) == 3 {
		h, _ := strconv.Atoi(hms[0])
		m, _ := strconv.Atoi(hms[1])
		s, _ := strconv.Atoi(hms[2])
		dt = CreateDayTime(h, m, s)
	} else {
		err = errors.New("invalid daytime format str;corrent format 12:10:09")
	}
	return
}
func (dt Clock) String() string {
	h, m, s := dt.HMS()
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
