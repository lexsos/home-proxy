package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonAccount struct {
	Login       string   `json:"login"`
	Password    *string  `json:"password"`
	Ips         []string `json:"ips"`
	ProfileSlug string   `json:"profile"`
}

type JsonAccounts struct {
	Accounts []JsonAccount `json:"accounts"`
}

func NewJsonHttpAuthenticator(fileName string) (*JsonHttpAuthenticator, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fileName, err)
	}
	var accounts JsonAccounts
	if err := json.Unmarshal(data, &accounts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &JsonHttpAuthenticator{
		accountsByLogin: loginMap(accounts.Accounts),
		accountsByIp:    ipMap(accounts.Accounts),
	}, nil
}

func loginMap(accounts []JsonAccount) map[string]JsonAccount {
	accountsbyLoginMap := make(map[string]JsonAccount)
	for _, account := range accounts {
		accountsbyLoginMap[account.Login] = account
	}
	return accountsbyLoginMap
}

func ipMap(accounts []JsonAccount) map[string]JsonAccount {
	accountsbyIpMap := make(map[string]JsonAccount)
	for _, account := range accounts {
		for _, ip := range account.Ips {
			accountsbyIpMap[ip] = account
		}
	}
	return accountsbyIpMap
}
