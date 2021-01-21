package elfsquad

import (
	"encoding/json"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
	//TokenType   string `json:"token_type"`
	//Scope       string `json:"scope"`
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

	skipAccessToken := true

	requestConfig := oauth2.RequestConfig{
		URL:             AccessTokenURL,
		BodyModel:       body,
		ResponseModel:   &accessToken,
		SkipAccessToken: &skipAccessToken,
	}

	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	expiresIn, _ := json.Marshal(accessToken.ExpiresIn)
	expiresInJSON := json.RawMessage(expiresIn)

	token := oauth2.Token{
		AccessToken: &accessToken.AccessToken,
		ExpiresIn:   &expiresInJSON,
	}

	return &token, nil
}
