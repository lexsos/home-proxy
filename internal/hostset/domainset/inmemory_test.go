package domainset

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DomainSetTestSuite struct {
	suite.Suite
	domainSet *InMemoryDomainsSet
}

func (s *DomainSetTestSuite) SetupTest() {
	s.domainSet = NewInMemoryDomainsSet()
}

func (s *DomainSetTestSuite) TestNoData() {
	contains, _ := s.domainSet.Contains("example.com")
	assert.False(s.T(), contains)
}

func (s *DomainSetTestSuite) TestContainsExactDomain() {
	s.domainSet.Add("example.com", ExactDomain)
	contains, _ := s.domainSet.Contains("example.com")
	assert.True(s.T(), contains)
}

func (s *DomainSetTestSuite) TestContainsSubDomain() {
	s.domainSet.Add("example.com", SubDomains)
	subTests := []string{"example.com", "sub.example.com", "sub.sub.example.com"}
	for _, domain := range subTests {
		s.T().Run(domain, func(t *testing.T) {
			contains, _ := s.domainSet.Contains(domain)
			assert.True(t, contains)
		})
	}
}

func (s *DomainSetTestSuite) TestMiss() {
	s.domainSet.Add("example.com", ExactDomain)
	s.domainSet.Add("other.com", SubDomains)
	subTests := []string{"invalid.com", "www.example.com"}
	for _, domain := range subTests {
		s.T().Run(domain, func(t *testing.T) {
			contains, _ := s.domainSet.Contains(domain)
			assert.False(t, contains)
		})
	}
}

func (s *DomainSetTestSuite) TestCaseInsensitiveExact() {
	s.domainSet.Add("Example.com", ExactDomain)
	subTests := []string{"EXAMPLE.COM", "ExAmPlE.CoM"}
	for _, domain := range subTests {
		s.T().Run(domain, func(t *testing.T) {
			contains, _ := s.domainSet.Contains(domain)
			assert.True(t, contains)
		})
	}
}

func (s *DomainSetTestSuite) TestCaseInsensitiveSub() {
	s.domainSet.Add("example.com", SubDomains)
	subTests := []string{"www.EXAMPLE.COM", "wWw.ExAmPlE.CoM"}
	for _, domain := range subTests {
		s.T().Run(domain, func(t *testing.T) {
			contains, _ := s.domainSet.Contains(domain)
			assert.True(t, contains)
		})
	}
}

func TestRunDomainSetTestSuite(t *testing.T) {
	suite.Run(t, new(DomainSetTestSuite))
}
