package elfsquad

import (
	"encoding/json"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_token "github.com/leapforce-libraries/go_oauth2/token"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
}

func (service *Service) GetAccessToken() (*go_token.Token, *errortools.Error) {
	body := struct {
		ClientID string `json:"clientId"`
		Secret   string `json:"secret"`
	}{
		service.clientID,
		service.clientSecret,
	}

	accessToken := AccessToken{}

	requestConfig := go_http.RequestConfig{
		Method:        accessTokenMethod,
		Url:           accessTokenURL,
		BodyModel:     body,
		ResponseModel: &accessToken,
	}

	_, _, e := service.httpRequestWithoutAccessToken(&requestConfig)
	if e != nil {
		return nil, e
	}

	expiresIn, _ := json.Marshal(accessToken.ExpiresIn / 1000)
	expiresInJSON := json.RawMessage(expiresIn)

	token := go_token.Token{
		AccessToken: &accessToken.AccessToken,
		ExpiresIn:   &expiresInJSON,
	}

	return &token, nil
}
