package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/lexsos/home-proxy/request"
)

type JsonAccount struct {
	Id       int64    `json:"id"`
	Login    string   `json:"login"`
	Role     string   `json:"role"`
	Password *string  `json:"password"`
	Ips      []string `json:"ips"`
}

type JsonHttpAuthenticator struct {
	accountsByLogin map[string]JsonAccount
	accountsByIp    map[string]JsonAccount
}

func NewJsonHttpAuthenticator(fileName string) (*JsonHttpAuthenticator, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fileName, err)
	}
	var accounts []JsonAccount
	if err := json.Unmarshal(data, &accounts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &JsonHttpAuthenticator{
		accountsByLogin: loginMap(accounts),
		accountsByIp:    ipMap(accounts),
	}, nil
}

func loginMap(accounts []JsonAccount) map[string]JsonAccount {
	accountsbyLoginMap := make(map[string]JsonAccount, len(accounts))
	for _, account := range accounts {
		accountsbyLoginMap[account.Login] = account
	}
	return accountsbyLoginMap
}

func ipMap(accounts []JsonAccount) map[string]JsonAccount {
	accountsbyIpMap := make(map[string]JsonAccount, len(accounts))
	for _, account := range accounts {
		for _, ip := range account.Ips {
			accountsbyIpMap[ip] = account
		}
	}
	return accountsbyIpMap
}

func (jsonAuth *JsonHttpAuthenticator) GetUser(r *http.Request) (*Account, error) {
	lp := request.GetLoginPass(r)
	if lp == nil {
		return nil, nil
	}
	account, ok := jsonAuth.accountsByLogin[lp.Login]
	if !ok {
		return nil, nil
	}
	if account.Password != nil && *account.Password != lp.Password {
		return nil, nil
	}
	return &Account{
		Id:    account.Id,
		Login: account.Login,
		Role:  account.Role,
	}, nil
}
