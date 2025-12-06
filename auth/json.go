package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type JsonAccount struct {
	Id       int64  `json:"id"`
	Login    string `json:"login"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type JsonHttpAuthenticator struct {
	accounts map[string]JsonAccount
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
	accountsMap := make(map[string]JsonAccount, len(accounts))
	for _, account := range accounts {
		accountsMap[account.Login] = account
	}
	return &JsonHttpAuthenticator{
		accounts: accountsMap,
	}, nil
}

func (jsonAuth *JsonHttpAuthenticator) GetUser(r *http.Request) (*Account, error) {
	lp := getLoginPass(r)
	if lp == nil {
		return nil, nil
	}
	account, ok := jsonAuth.accounts[lp.Login]
	if !ok {
		return nil, nil
	}
	if account.Password != lp.Password {
		return nil, nil
	}
	return &Account{
		Id:    account.Id,
		Login: account.Login,
		Role:  account.Role,
	}, nil
}
