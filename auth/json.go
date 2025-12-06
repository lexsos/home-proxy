package auth

import (
	"net/http"
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

func NewJsonHttpAuthenticator(FileName string) (*JsonAccount, error) {
	
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
	return &Account{
		Id:    account.Id,
		Login: account.Login,
		Role:  account.Role,
	}, nil
}
