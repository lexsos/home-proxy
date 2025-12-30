package auth

import (
	"context"
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InMemoryAuthRepositoryTestSuit struct {
	suite.Suite
	repo *InMemoryAuthRepository
}

func (s *InMemoryAuthRepositoryTestSuit) SetupTest() {
	s.repo = NewInMemoryAuthRepository()
}

func (s *InMemoryAuthRepositoryTestSuit) TestDuplicateLogin() {
	err := s.repo.AddWithPassword("test", "test", "test")
	s.NoError(err)
	err = s.repo.AddWithPassword("test", "test", "test")
	s.ErrorContains(err, "login 'test' already exists")
}

func (s *InMemoryAuthRepositoryTestSuit) TestDuplicateIp() {
	err := s.repo.AddWithIps("test", "test", []string{"127.0.0.1", "127.0.0.2"})
	s.NoError(err)
	err = s.repo.AddWithIps("test", "test", []string{"127.0.0.2"})
	s.ErrorContains(err, "ip '127.0.0.2' already exists")
}

func (s *InMemoryAuthRepositoryTestSuit) TestAuthByPassword() {
	s.repo.AddWithPassword("test", "test", "test")
	acc, err := s.repo.AuthByPassword(context.Background(), "test", "test")
	s.NoError(err)
	s.Equal("test", acc.Login)
}

func (s *InMemoryAuthRepositoryTestSuit) TestFailAuthByPassword() {
	s.repo.AddWithPassword("test", "test", "test")
	account, err := s.repo.AuthByPassword(context.Background(), "test", "wrong")
	s.NoError(err)
	s.Nil(account)
}

func (s *InMemoryAuthRepositoryTestSuit) TestAuthByIp() {
	s.repo.AddWithIps("test", "test", []string{"127.0.0.1", "127.0.0.2"})
	account, err := s.repo.AuthByIp(context.Background(), "127.0.0.1")
	s.NoError(err)
	s.Equal("test", account.Login)
}

func (s *InMemoryAuthRepositoryTestSuit) TestFailAuthByIp() {
	s.repo.AddWithIps("test", "test", []string{"127.0.0.1", "127.0.0.2"})
	account, err := s.repo.AuthByIp(context.Background(), "127.0.0.3")
	s.NoError(err)
	s.Nil(account)
}

func (s *InMemoryAuthRepositoryTestSuit) TestGetUserByPassword() {
	s.repo.AddWithPassword("testuser", "testprofile", "testpass")
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	auth := base64.StdEncoding.EncodeToString([]byte("testuser:testpass"))
	req.Header.Set("Proxy-Authorization", "Basic "+auth)
	account, err := s.repo.GetUser(context.Background(), req)
	s.NoError(err)
	s.Equal("testuser", account.Login)
}

func (s *InMemoryAuthRepositoryTestSuit) TestGetUserByIp() {
	s.repo.AddWithIps("ipuser", "ipprofile", []string{"192.168.1.1"})
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.RemoteAddr = "192.168.1.1:12345"
	account, err := s.repo.GetUser(context.Background(), req)
	s.NoError(err)
	s.Equal("ipuser", account.Login)
}

func (s *InMemoryAuthRepositoryTestSuit) TestGetUserNoAuth() {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.RemoteAddr = "10.0.0.1:12345"
	account, err := s.repo.GetUser(context.Background(), req)
	s.NoError(err)
	s.Nil(account)
}

func (s *InMemoryAuthRepositoryTestSuit) TestGetByLogin() {
	s.repo.AddWithPassword("testuser", "testprofile", "testpass")
	account, err := s.repo.GetByLogin("testuser")
	s.NoError(err)
	s.NotNil(account)
	s.Equal("testprofile", account.ProfileSlug)
}

func TestRunInMemoryAuthRepositoryTestSuit(t *testing.T) {
	suite.Run(t, new(InMemoryAuthRepositoryTestSuit))
}
