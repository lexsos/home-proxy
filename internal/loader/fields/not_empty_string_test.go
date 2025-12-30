package fields

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NotEmptyStringTestSuit struct {
	suite.Suite
	str NotEmptyString
}

func (s *NotEmptyStringTestSuit) TestEmptyString() {
	err := json.Unmarshal([]byte("\"\""), &s.str)
	s.ErrorContains(err, "string is empty")
}

func (s *NotEmptyStringTestSuit) TestValidString() {
	err := json.Unmarshal([]byte("\"test\""), &s.str)
	s.NoError(err)
	s.Equal("test", string(s.str))
}

func (s *NotEmptyStringTestSuit) TestToString() {
	str := NotEmptyString("test")
	s.Equal("test", str.String())
}

func TestRunNotEmptyStringTestSuit(t *testing.T) {
	suite.Run(t, new(NotEmptyStringTestSuit))
}
