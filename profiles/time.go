package profiles

import (
	"time"
)

const (
	timeLayout = "15:04:05"
)

type DayTime uint32

func SecondsSinceMidnight(t time.Time) DayTime {
	return DayTime(t.Hour()*3600 + t.Minute()*60 + t.Second())
}

func ParseTime(strTime string) (DayTime, error) {
	start, err := time.Parse(timeLayout, strTime)
	if err != nil {
		return 0, err
	}
	return SecondsSinceMidnight(start), nil
}

func NowTimeInLocation(location *time.Location) DayTime {
	currentTime := time.Now().In(location)
	return SecondsSinceMidnight(currentTime)
}
