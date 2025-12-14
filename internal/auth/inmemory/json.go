package inmemory

import (
	"encoding/json"
	"fmt"
	"os"
)

type Accounts struct {
	Accounts []AccountData `json:"accounts"`
}

func NewHttpAuthenticatorFromJson(fileName string) (*HttpAuthenticator, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fileName, err)
	}
	var accounts Accounts
	if err := json.Unmarshal(data, &accounts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &HttpAuthenticator{
		accountsByLogin: loginMap(accounts.Accounts),
		accountsByIp:    ipMap(accounts.Accounts),
	}, nil
}

func loginMap(accounts []AccountData) map[string]AccountData {
	accountsbyLoginMap := make(map[string]AccountData)
	for _, account := range accounts {
		accountsbyLoginMap[account.Login] = account
	}
	return accountsbyLoginMap
}

func ipMap(accounts []AccountData) map[string]AccountData {
	accountsbyIpMap := make(map[string]AccountData)
	for _, account := range accounts {
		for _, ip := range account.Ips {
			accountsbyIpMap[ip] = account
		}
	}
	return accountsbyIpMap
}
