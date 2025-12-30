package loader

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lexsos/home-proxy/internal/auth"
	"github.com/lexsos/home-proxy/internal/utils/logging"
)

type AccountData struct {
	Login       string   `json:"login"`
	Password    *string  `json:"password"`
	Ips         []string `json:"ips"`
	ProfileSlug string   `json:"profile"`
}

type Accounts struct {
	Accounts []AccountData `json:"accounts"`
}

func LoadAuthRepository(fileName string) (*auth.InMemoryAuthRepository, error) {
	logger := logging.LogFromContext(context.Background())
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var accounts Accounts
	if err := json.Unmarshal(data, &accounts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	repo := auth.NewInMemoryAuthRepository()
	for _, account := range accounts.Accounts {
		if account.Login == "" {
			return nil, fmt.Errorf("account login is empty")
		}
		if account.Password != nil && *account.Password != "" {
			err := repo.AddWithPassword(account.Login, account.ProfileSlug, *account.Password)
			if err != nil {
				return nil, fmt.Errorf("failed to add account: %w", err)
			}
			logger.Infof("Add account '%s' with password", account.Login)
		}
		if len(account.Ips) > 0 {
			err := repo.AddWithIps(account.Login, account.ProfileSlug, account.Ips)
			if err != nil {
				return nil, fmt.Errorf("failed to add account: %w", err)
			}
			logger.Infof("Add account '%s' with ips", account.Login)
		}
	}
	return repo, nil
}
