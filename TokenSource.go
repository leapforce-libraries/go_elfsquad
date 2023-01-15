package elfsquad

import (
	"encoding/json"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_token "github.com/leapforce-libraries/go_oauth2/token"
)

type TokenSource struct {
	token   *go_token.Token
	service *Service
}

func NewTokenSource(service *Service) (*TokenSource, *errortools.Error) {
	if service == nil {
		return nil, errortools.ErrorMessage("Service is a nil pointer")
	}

	return &TokenSource{
		service: service,
	}, nil
}

func (t *TokenSource) Token() *go_token.Token {
	return t.token
}

func (t *TokenSource) NewToken() (*go_token.Token, *errortools.Error) {
	return t.service.GetAccessToken()
}

func (t *TokenSource) SetToken(token *go_token.Token, save bool) *errortools.Error {
	t.token = token

	if !save {
		return nil
	}

	return t.SaveToken()
}

func (t *TokenSource) RetrieveToken() *errortools.Error {
	return nil
}

func (t *TokenSource) SaveToken() *errortools.Error {
	return nil
}

func (m *TokenSource) UnmarshalToken(b []byte) (*go_token.Token, *errortools.Error) {
	var token go_token.Token

	err := json.Unmarshal(b, &token)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	return &token, nil
}
