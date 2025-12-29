package hostset

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/lexsos/home-proxy/internal/hostset/domainset"
	"github.com/lexsos/home-proxy/internal/hostset/ipset"
)

type InMemoryHostRepositoryTestSuit struct {
	suite.Suite
	repo    *InMemoryHostRepository
	ips     *ipset.InMemoryIpSet
	domains *domainset.InMemoryDomainsSet
}

func (s *InMemoryHostRepositoryTestSuit) SetupTest() {
	s.repo = NewInMemoryHostRepository()
	s.ips = ipset.NewInMemoryIpSet()
	s.domains = domainset.NewInMemoryDomainsSet()
	s.repo.AddHostSet("test", s.ips, s.domains)
}

func (s *InMemoryHostRepositoryTestSuit) TestNoData() {
	subTests := []string{"127.0.0.1", "example.com"}
	for _, host := range subTests {
		s.T().Run(host, func(t *testing.T) {
			s.False(s.repo.Contains(host, []string{"test"}))
		})
	}
}

func (s *InMemoryHostRepositoryTestSuit) TestNoSet() {
	subTests := []string{"127.0.0.1", "example.com"}
	for _, host := range subTests {
		s.T().Run(host, func(t *testing.T) {
			s.False(s.repo.Contains(host, []string{"wrong"}))
		})
	}
}

func (s *InMemoryHostRepositoryTestSuit) TestContainsIp() {
	s.ips.Add("127.0.0.1")
	s.True(s.repo.Contains("127.0.0.1", []string{"test"}))
}

func (s *InMemoryHostRepositoryTestSuit) TestContainsDomain() {
	s.domains.Add("example.com", domainset.ExactDomain)
	s.True(s.repo.Contains("example.com", []string{"test"}))
}

func (s *InMemoryHostRepositoryTestSuit) TestSeveralSets() {
	ips := ipset.NewInMemoryIpSet()
	domains := domainset.NewInMemoryDomainsSet()
	ips.Add("127.0.0.1")
	domains.Add("example.com", domainset.ExactDomain)
	s.repo.AddHostSet("example", ips, domains)
	subTests := []string{"127.0.0.1", "example.com"}
	for _, host := range subTests {
		s.T().Run(host, func(t *testing.T) {
			s.True(s.repo.Contains(host, []string{"test", "example"}))
		})
	}
}

func (s *InMemoryHostRepositoryTestSuit) TestFilterBySetsList() {
	ips := ipset.NewInMemoryIpSet()
	domains := domainset.NewInMemoryDomainsSet()
	ips.Add("127.0.0.1")
	domains.Add("example.com", domainset.ExactDomain)
	s.repo.AddHostSet("example", ips, domains)
	subTests := []string{"127.0.0.1", "example.com"}
	for _, host := range subTests {
		s.T().Run(host, func(t *testing.T) {
			s.False(s.repo.Contains(host, []string{"test"}))
		})
	}
}

func TestRunInMemoryHostRepositoryTestSuit(t *testing.T) {
	suite.Run(t, new(InMemoryHostRepositoryTestSuit))
}
