package auth

import (
	"context"
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

func TestRunInMemoryAuthRepositoryTestSuit(t *testing.T) {
	suite.Run(t, new(InMemoryAuthRepositoryTestSuit))
}
