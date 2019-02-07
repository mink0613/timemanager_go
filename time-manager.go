package timemanager

import (
	"time"
)

var zone = getZone()
var timeFormat = "2006-01-02 " + zone
var timeDetailFormat = "2006-01-02 15:04:05 " + zone

func getZone() string {

	zone, _ := time.Now().Local().Zone()
	return zone
}

// GetNow will return current local time with format YYYY-MM-DD
func GetNow() string {

	return time.Now().Local().Format(timeFormat)
}

// GetNowDetail will return current local time with format YYYY-MM-DD HH:MM:SS
func GetNowDetail() string {
	return time.Now().Local().Format(timeDetailFormat)
}

// GetTimeFromString will convert string to time format YYYY-MM-DD
func GetTimeFromString(date string) time.Time {

	converted, _ := time.Parse(timeFormat, date)
	return converted
}

// GetTimeDetailFromString will convert string to time format YYYY-MM-DD HH:MM:SS
func GetTimeDetailFromString(date string) time.Time {

	converted, _ := time.Parse(timeDetailFormat, date)
	return converted
}

// GetTimeDiffInDays will convert string to time format and get the time difference of two values as days
// Return value must be date1 - date2
func GetTimeDiffInDays(date1 string, date2 string) int {

	converted1 := GetTimeFromString(date1)
	converted2 := GetTimeFromString(date2)

	return int(converted1.Sub(converted2).Hours() / 24)
}

// GetTimeDiffInSeconds will convert string to time format and get the time difference of two values as seconds
// Return value must be date1 - date2
func GetTimeDiffInSeconds(date1 string, date2 string) int {

	converted1 := GetTimeDetailFromString(date1)
	converted2 := GetTimeDetailFromString(date2)

	return int(converted1.Sub(converted2).Seconds())
}

// GetYesterday will return yestrday with format YYYY-MM-DD
func GetYesterday() string {
	return time.Now().Add(-24 * time.Hour).Local().Format(timeFormat)
}

// IsSameDay will compare two different date strings, nowSimple is a date with format YYYY-MM-DD
// and nowDetail is a date with format YYYY-MM-DD HH:MM:SS
func IsSameDay(nowSimple string, nowDetail string) bool {

	t1, _ := time.Parse(timeFormat, nowSimple)
	t2, _ := time.Parse(timeDetailFormat, nowDetail)

	if t1.Year() != t2.Year() || t1.Month() != t2.Month() || t1.Day() != t2.Day() {
		return false
	}
	return true
}

// GetTimeDiffFromNowHour will calculate time diff and return diff hours
func GetTimeDiffFromNowHour(timeDetail string) int {

	t1, _ := time.Parse(timeDetailFormat, timeDetail)
	t2, _ := time.Parse(timeDetailFormat, GetNowDetail())

	return int(t2.Sub(t1.Local()).Hours())
}
