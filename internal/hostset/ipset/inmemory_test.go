package ipset

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InMemoryIpSetTestSuite struct {
	suite.Suite
	IpSet *InMemoryIpSet
}

func (s *InMemoryIpSetTestSuite) SetupTest() {
	s.IpSet = NewInMemoryIpSet()
}

func (s *InMemoryIpSetTestSuite) TestNoData() {
	ip := net.ParseIP("192.168.1.1")
	contains, _ := s.IpSet.Contains(ip)
	assert.False(s.T(), contains)
}

func (s *InMemoryIpSetTestSuite) TestFind4ByAddress() {
	s.IpSet.Add("192.168.1.1")
	ip := net.ParseIP("192.168.1.1")
	contains, _ := s.IpSet.Contains(ip)
	assert.True(s.T(), contains)
}

func (s *InMemoryIpSetTestSuite) TestFind4ByCIDR() {
	s.IpSet.Add("192.168.1.0/24")
	ip := net.ParseIP("192.168.1.1")
	contains, _ := s.IpSet.Contains(ip)
	assert.True(s.T(), contains)
}

func (s *InMemoryIpSetTestSuite) TestFind6ByAddress() {
	s.IpSet.Add("2001:db8::1")
	ip := net.ParseIP("2001:db8::1")
	contains, _ := s.IpSet.Contains(ip)
	assert.True(s.T(), contains)
}

func (s *InMemoryIpSetTestSuite) TestFind6ByCIDR() {
	s.IpSet.Add("2001:db8::/64")
	ip := net.ParseIP("2001:db8::1")
	contains, _ := s.IpSet.Contains(ip)
	assert.True(s.T(), contains)
}

func (s *InMemoryIpSetTestSuite) TestMiss() {
	s.IpSet.Add("192.168.1.0/24")
	s.IpSet.Add("192.168.2.1")
	ip := net.ParseIP("192.168.3.1")
	contains, _ := s.IpSet.Contains(ip)
	assert.False(s.T(), contains)
}

func (s *InMemoryIpSetTestSuite) TestAddInvalid() {
	err := s.IpSet.Add("192.168.1.0/33")
	assert.EqualError(s.T(), err, "invalid ip or network: 192.168.1.0/33")
}

func TestRunnMemoryIpSetSuite(t *testing.T) {
	suite.Run(t, new(InMemoryIpSetTestSuite))
}
