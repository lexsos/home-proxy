package times

import (
	"time"
)

type WeekTime struct {
	Time DayTime
	Day  time.Weekday
}

func NowTimeInLocation(location *time.Location) *WeekTime {
	currentTime := time.Now().In(location)
	return &WeekTime{
		Time: SecondsSinceMidnight(currentTime),
		Day:  currentTime.Weekday(),
	}
}
