package ipset

import (
	"fmt"
	"net"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SignatureTestSuite struct {
	suite.Suite
}

func (s *SignatureTestSuite) TestIs4() {
	sig, _ := NewIpSignature(net.ParseIP("192.168.1.1"))
	assert.True(s.T(), sig.Is4())
	assert.False(s.T(), sig.Is6())
}

func (s *SignatureTestSuite) TestIs6() {
	sig, _ := NewIpSignature(net.ParseIP("2001:db8::1"))
	assert.False(s.T(), sig.Is4())
	assert.True(s.T(), sig.Is6())
}

func (s *SignatureTestSuite) TestMaskFor4() {
	subTests := []struct {
		maskLen  int
		expected string
	}{
		{32, "192.168.1.1"},
		{24, "192.168.1.0"},
		{16, "192.168.0.0"},
		{8, "192.0.0.0"},
		{0, "0.0.0.0"},
	}

	for _, data := range subTests {
		s.T().Run(data.expected, func(t *testing.T) {
			sig, _ := NewIpSignature(net.ParseIP("192.168.1.1"))
			expected := toIP4(net.ParseIP(data.expected))
			masked, err := sig.GetForMask4(data.maskLen)
			assert.Nil(s.T(), err)
			assert.Equal(s.T(), *expected, masked)
		})
	}
}

func (s *SignatureTestSuite) TestInvalidMaskLen() {
	subTests := []int{-1, 33}
	for _, maskLen := range subTests {
		name := fmt.Sprintf("invalid mask len: %d", maskLen)
		s.T().Run(name, func(t *testing.T) {
			sig, _ := NewIpSignature(net.ParseIP("192.168.1.1"))
			_, err := sig.GetForMask4(maskLen)
			assert.EqualError(s.T(), err, "can't get ip for mask4: invalid mask len: "+strconv.Itoa(maskLen))
		})
	}
}

func (s *SignatureTestSuite) TestWrongVersion() {
	sig, _ := NewIpSignature(net.ParseIP("2001:db8::1"))
	_, err := sig.GetForMask4(32)
	assert.EqualError(s.T(), err, "can't get ip for mask4: invalid ip version: 6")
}

func (s *SignatureTestSuite) TestMaskFor6() {
	subTests := []struct {
		maskLen  int
		expected string
	}{
		{128, "2001:db8::1"},
		{64, "2001:db8::0"},
	}

	for _, data := range subTests {
		s.T().Run(data.expected, func(t *testing.T) {
			ip := net.ParseIP("2001:db8::1")
			sig, _ := NewIpSignature(ip)
			masked, err := sig.GetForMask6(data.maskLen)
			assert.Nil(s.T(), err)
			assert.Equal(s.T(), *toIP6(net.ParseIP(data.expected)), masked)
		})
	}
}

func TestRunSignatureSuite(t *testing.T) {
	suite.Run(t, new(SignatureTestSuite))
}
