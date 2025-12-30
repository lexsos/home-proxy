package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lexsos/home-proxy/internal/utils/logging"
	"github.com/lexsos/home-proxy/internal/utils/request"
)

type AccountData struct {
	Login       string
	Password    *string
	Ips         []string
	ProfileSlug string
}

type InMemoryAuthRepository struct {
	accountsByLogin map[string]AccountData
	accountsByIp    map[string]AccountData
}

func NewInMemoryAuthRepository() *InMemoryAuthRepository {
	return &InMemoryAuthRepository{
		accountsByLogin: make(map[string]AccountData),
		accountsByIp:    make(map[string]AccountData),
	}
}

func (repo *InMemoryAuthRepository) AddWithIps(login string, profileSlug string, ips []string) error {
	data := AccountData{
		Login:       login,
		ProfileSlug: profileSlug,
		Ips:         ips,
	}
	for _, ip := range ips {
		if _, ok := repo.accountsByIp[ip]; ok {
			return fmt.Errorf("ip '%s' already exists", ip)
		}
		if _, ok := repo.accountsByLogin[login]; ok {
			return fmt.Errorf("login '%s' already exists", login)
		}
	}
	for _, ip := range ips {
		repo.accountsByIp[ip] = data
	}
	return nil
}

func (repo *InMemoryAuthRepository) AddWithPassword(login string, profileSlug string, password string) error {
	data := AccountData{
		Login:       login,
		ProfileSlug: profileSlug,
		Password:    &password,
	}
	if _, ok := repo.accountsByLogin[login]; ok {
		return fmt.Errorf("login '%s' already exists", login)
	}
	repo.accountsByLogin[login] = data
	return nil
}

func (repo *InMemoryAuthRepository) AuthByPassword(ctx context.Context, login string, password string) (*Account, error) {
	logger := logging.LogFromContext(ctx)
	logger.Debug("Try auth by login/password")
	data, ok := repo.accountsByLogin[login]
	if !ok {
		logger.Debugf("Login '%s' not found", login)
		return nil, nil
	}
	if data.Password == nil || *data.Password == "" {
		logger.Debugf("Password not set for login '%s'", login)
		return nil, nil
	}
	if *data.Password != password {
		logger.Debugf("Password not match for login '%s'", login)
		return nil, nil
	}
	return &Account{
		Login:       data.Login,
		ProfileSlug: data.ProfileSlug,
	}, nil
}

func (repo *InMemoryAuthRepository) AuthByIp(ctx context.Context, ip string) (*Account, error) {
	logger := logging.LogFromContext(ctx)
	logger.Debug("Try auth by ip")
	data, ok := repo.accountsByIp[ip]
	if !ok {
		logger.Debugf("Account for ip '%s' not found", ip)
		return nil, nil
	}
	return &Account{
		Login:       data.Login,
		ProfileSlug: data.ProfileSlug,
	}, nil
}

func (repo *InMemoryAuthRepository) GetByLogin(login string) (*Account, error) {
	data, ok := repo.accountsByLogin[login]
	if !ok {
		return nil, nil
	}
	return &Account{
		Login:       data.Login,
		ProfileSlug: data.ProfileSlug,
	}, nil
}

func (repo *InMemoryAuthRepository) GetUser(ctx context.Context, r *http.Request) (*Account, error) {
	logger := logging.LogFromContext(ctx)
	lp := request.GetLoginPass(r)
	if lp != nil {
		account, err := repo.AuthByPassword(ctx, lp.Login, lp.Password)
		if err != nil {
			return nil, fmt.Errorf("error auth by password: %w", err)
		}
		if account != nil {
			logger.Debugf("success auth by password for account '%s'", account.Login)
			return account, nil
		}
		logger.Debugf("fail auth by password for account '%s'", lp.Login)
	}
	ip := request.GetClientIpAddress(r)
	account, err := repo.AuthByIp(ctx, ip)
	if err != nil {
		return nil, fmt.Errorf("error auth by ip: %w", err)
	}
	if account == nil {
		logger.Debugf("fail auth by ip for '%s'", ip)
		return nil, nil
	}
	logger.Debugf("success auth by ip for account '%s'", account.Login)
	return account, nil
}
