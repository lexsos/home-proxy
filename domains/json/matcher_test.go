package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatcherTestSuite struct {
	suite.Suite
}

func (suite *MatcherTestSuite) TestDomainToLower() {
	data := prepareDomain("Api.Google.com")
	assert.Equal(suite.T(), "api.google.com", data.original, "Domain to lower")
}

func (suite *MatcherTestSuite) TestDomainParents() {
	data := prepareDomain("api.google.com")
	assert.Equal(suite.T(), []string{"google.com", "com"}, data.parents, "Domain parents")
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(MatcherTestSuite))
}
