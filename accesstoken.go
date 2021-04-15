package elfsquad

import (
	"encoding/json"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
}

func (service *Service) GetAccessToken() (*oauth2.Token, *errortools.Error) {
	body := struct {
		ClientID string `json:"clientId"`
		Secret   string `json:"secret"`
	}{
		service.clientID,
		service.clientSecret,
	}

	accessToken := AccessToken{}

	requestConfig := go_http.RequestConfig{
		URL:           accessTokenURL,
		BodyModel:     body,
		ResponseModel: &accessToken,
	}

	_, _, e := service.httpRequest(accessTokenMethod, &requestConfig, true)
	if e != nil {
		return nil, e
	}

	expiresIn, _ := json.Marshal(accessToken.ExpiresIn / 1000)
	expiresInJSON := json.RawMessage(expiresIn)

	token := oauth2.Token{
		AccessToken: &accessToken.AccessToken,
		ExpiresIn:   &expiresInJSON,
	}

	return &token, nil
}
