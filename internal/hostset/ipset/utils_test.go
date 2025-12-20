package ipset

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (s *UtilsTestSuite) TestToIP4() {
	ip := net.ParseIP("192.168.1.1")
	assert.NotNil(s.T(), toIP4(ip))
}

func (s *UtilsTestSuite) TestToIP6() {
	ip := net.ParseIP("2001:db8::1")
	assert.NotNil(s.T(), toIP6(ip))
}

func (s *UtilsTestSuite) TestFailToIP4() {
	ip := net.ParseIP("2a02:6bf:8080:d21::1:2b")
	assert.Nil(s.T(), toIP4(ip))
}

func (s *UtilsTestSuite) TestFailToIP6() {
	ip := net.ParseIP("192.168.1.1")
	assert.NotNil(s.T(), toIP6(ip))
}

func TestRunUtilsSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}
