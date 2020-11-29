package helper

import (
	"fmt"
	"time"
)

const EARTH_RADIUS float64 = 6378.137
const LocalTimeFormat = "2006-01-02 15:04:05 -0700"

func GetUnixTimeStamp() int {
	now := time.Now().Unix()
	return int(now)
}

func GetNanoTimeStamp() int64 {
	now := time.Now().UnixNano()
	return now
}

func ToDatetimeStr(t *time.Time) string {
	if t.IsZero() {
		return "1970-01-01 08:00:00"
	}

	return fmt.Sprintf("%4d-%0.2d-%0.2d %0.2d:%0.2d:%0.2d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
}

func GetDateStr(t *time.Time) string {
	if t.IsZero() {
		return "1970-01-01 08:00:00"
	}

	return fmt.Sprintf("%4d-%0.2d-%0.2d 00:00:00",
		t.Year(),
		t.Month(),
		t.Day())
}

func GetDateInt(t *time.Time) int {
	if t.IsZero() {
		return 19700101
	}
	year, month, day := t.Date()
	return year*10000 + int(month)*100 + day
}

func ToLocalTimeSeconds(date string) (int64, error) {
	date += " +0800"
	t, e := time.Parse(LocalTimeFormat, date)
	if e != nil {
		return 0, e
	}
	return t.Unix(), nil
}

func ToLocalTimeMilliSeconds(date string) (int64, error) {
	date += " +0800"
	t, e := time.Parse(LocalTimeFormat, date)
	if e != nil {
		return 0, e
	}
	return t.UnixNano() / 1000000, nil
}

func ToLocalTime(date string) (time.Time, error) {
	date += " +0800"
	return time.Parse(LocalTimeFormat, date)
}

/*
 * timeOp: 0:一年以内，1：今天，2：最近7天，3：最近30天
 */
func CalBegin(timeOp int) int64 {

	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	var tm time.Time
	switch timeOp {
	case 1:
		tm = t
	case 2:
		tm = t.AddDate(0, 0, -7)
	case 3:
		tm = t.AddDate(0, -30, 0)
	default:
		tm = t.AddDate(-1, 0, 0)
	}

	return tm.Unix()
}

func MilliSec2time(milliSec int64) time.Time {
	return time.Unix(milliSec/1000, (milliSec%1000)*1000000)
}

func GetNowMilliSec() int64 {
	return time.Now().UnixNano() / 1e6
}
