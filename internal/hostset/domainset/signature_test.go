package domainset

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DomainSignatureTestSuite struct {
	suite.Suite
	DomainSig *DomainSig
}

func (s *DomainSignatureTestSuite) SetupTest() {
	s.DomainSig = NewDomainSig("www.example.com")
}

func (s *DomainSignatureTestSuite) TestOriginal() {
	assert.Equal(s.T(), "www.example.com", s.DomainSig.Original())
}

func (s *DomainSignatureTestSuite) TestParents() {
	assert.Equal(s.T(), []string{"example.com", "com"}, s.DomainSig.Parents())
}

func (s *DomainSignatureTestSuite) TestTopDomain() {
	sig := NewDomainSig("com")

	assert.Equal(s.T(), 0, len(sig.Parents()))
}

func TestRunDomainSignatureTestSuite(t *testing.T) {
	suite.Run(t, new(DomainSignatureTestSuite))
}
