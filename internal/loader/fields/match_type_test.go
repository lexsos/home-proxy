package fields

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MatchTypeTestSuit struct {
	suite.Suite
	mt MatchType
}

func (s *MatchTypeTestSuit) TestEmptyData() {
	raw := []byte("\"\"")
	err := json.Unmarshal(raw, &s.mt)
	s.ErrorContains(err, "unknown match type: ''")
}

func (s *MatchTypeTestSuit) TestUnknownData() {
	raw := []byte("\"unknown\"")
	err := json.Unmarshal(raw, &s.mt)
	s.ErrorContains(err, "unknown match type: 'unknown'")
}

func (s *MatchTypeTestSuit) TestValidData() {
	subTests := []struct {
		raw []byte
		mt  MatchType
	}{
		{[]byte("\"exact\""), ExactHost},
		{[]byte("\"subdomains\""), SubDomainsHosts},
		{[]byte("\"ip\""), Ip},
	}
	for _, tt := range subTests {
		s.T().Run(string(tt.mt), func(t *testing.T) {
			err := json.Unmarshal(tt.raw, &s.mt)
			s.NoError(err)
			s.Equal(tt.mt, s.mt)
		})
	}
}

func TestRunMatchTypeTestSuit(t *testing.T) {
	suite.Run(t, new(MatchTypeTestSuit))
}
