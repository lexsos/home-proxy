package times

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type DayTimeTestSuite struct {
	suite.Suite
}

func (s *DayTimeTestSuite) TestSecondsSinceMidnight() {
	subTests := []struct {
		time string
		sec  int
	}{
		{"00:00:00", 0},
		{"00:00:10", 10},
		{"00:01:10", 70},
		{"12:00:00", 12 * 60 * 60},
		{"23:59:59", 23*60*60 + 59*60 + 59},
	}
	for _, tt := range subTests {
		s.T().Run(tt.time, func(t *testing.T) {
			parsed, err := time.Parse(timeLayout, tt.time)
			s.NoError(err)
			seconds := int(SecondsSinceMidnight(parsed))
			s.Equal(tt.sec, seconds)
		})
	}
}

func (s *DayTimeTestSuite) TestParseDayTime() {
	seconds, err := ParseTime("00:00:10")
	s.NoError(err)
	s.Equal(10, int(seconds))
}

func (s *DayTimeTestSuite) TestParserError() {
	_, err := ParseTime("00:00")
	s.ErrorContains(err, "fail parse '00:00' as time")
}

func TestRunDayTimeTestSuite(t *testing.T) {
	suite.Run(t, new(DayTimeTestSuite))
}
