package utime

import (
	"fmt"
	"time"
)

const (
	YYYY_MM_DD_HH_MM_SS string = "2006-01-02 15:04:05"
	HH_MM_SS            string = "15:04:05"
)

func TimeS(date time.Time) int64 {
	return date.Unix()
}

func TimeMS(date time.Time) int64 {
	return ToTimeNS(date) / 1000000
}

func ToTimeNS(date time.Time) int64 {
	return date.UnixNano()
}

func Full(date time.Time) string {
	return fmt.Sprintf("%s %s %d %d %s", GetShortWeekDay(date.Weekday().String()), GetShortMonth(date.Month().String()), date.Day(), date.Year(), Time(date))
}

func Date(date time.Time) string {
	return fmt.Sprintf("%s %d %d", GetShortMonth(date.Month().String()), date.Day(), date.Year())
}

func Time(date time.Time) string {
	return date.Format(HH_MM_SS)
}

func StringToDate(dateStr string) (time.Time, error) {
	local, err := time.LoadLocation("Local")
	if err != nil {
		return time.Now(), err
	}
	dt, err := time.ParseInLocation(YYYY_MM_DD_HH_MM_SS, dateStr, local)
	if err != nil {
		return time.Now(), err
	}
	return dt, nil
}

func GetShortMonth(month string) string {
	switch month {
	case time.February.String():
		return "Feb"
	case time.March.String():
		return "Mar"
	case time.April.String():
		return "Apr"
	case time.May.String():
		return "May"
	case time.June.String():
		return "Jun"
	case time.July.String():
		return "Jul"
	case time.August.String():
		return "Aug"
	case time.September.String():
		return "Sept"
	case time.October.String():
		return "Oct"
	case time.November.String():
		return "Nov"
	case time.December.String():
		return "Dec"
	default:
		return "Jan"
	}
}

func GetShortWeekDay(weekDay string) string {
	switch weekDay {
	case time.Monday.String():
		return "Mon"
	case time.Tuesday.String():
		return "Tue"
	case time.Wednesday.String():
		return "Wed"
	case time.Thursday.String():
		return "Thur"
	case time.Friday.String():
		return "Fri"
	case time.Saturday.String():
		return "Sat"
	default:
		return "Sun"
	}
}
