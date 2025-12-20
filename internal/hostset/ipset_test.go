package hostset

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IpSetTestSuite struct {
	suite.Suite
	IpSet *InMemoryIpSet
}

func (s *IpSetTestSuite) SetupTest() {
	s.IpSet = NewInMemoryIpSet()
}

func (s *IpSetTestSuite) TestNoData() {
	ip := net.ParseIP("192.168.1.1")
	assert.False(s.T(), s.IpSet.Contains(ip))
}

func (s *IpSetTestSuite) TestFind4ByAddress() {
	s.IpSet.Add("192.168.1.1")
	ip := net.ParseIP("192.168.1.1")
	assert.True(s.T(), s.IpSet.Contains(ip))
}

func (s *IpSetTestSuite) TestFind4ByCIDR() {
	s.IpSet.Add("192.168.1.0/24")
	ip := net.ParseIP("192.168.1.1")
	assert.True(s.T(), s.IpSet.Contains(ip))
}

func (s *IpSetTestSuite) TestFind6ByAddress() {
	s.IpSet.Add("2001:db8::1")
	ip := net.ParseIP("2001:db8::1")
	assert.True(s.T(), s.IpSet.Contains(ip))
}

func (s *IpSetTestSuite) TestFind6ByCIDR() {
	s.IpSet.Add("2001:db8::/64")
	ip := net.ParseIP("2001:db8::1")
	assert.True(s.T(), s.IpSet.Contains(ip))
}

func (s *IpSetTestSuite) TestMiss() {
	s.IpSet.Add("192.168.1.0/24")
	s.IpSet.Add("192.168.2.1")
	ip := net.ParseIP("192.168.3.1")
	assert.False(s.T(), s.IpSet.Contains(ip))
}

func (s *IpSetTestSuite) TestAddInvalid() {
	err := s.IpSet.Add("192.168.1.0/33")
	assert.EqualError(s.T(), err, "invalid ip or network: 192.168.1.0/33")
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(IpSetTestSuite))
}
