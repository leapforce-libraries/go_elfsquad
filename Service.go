package elfsquad

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	apiName              string = "Elfsquad"
	apiURLData           string = "https://api.elfsquad.io/data/1"
	accessTokenURL       string = "https://api.elfsquad.io/api/2/auth/elfskotconnectlogin"
	accessTokenMethod    string = http.MethodPost
	accessTokenGrantType string = "client_credentials"
	accessTokenScope     string = "Elfskot.Api"
)

// Service stores Service configuration
//
type Service struct {
	clientID     string
	clientSecret string
	oAuth2       *oauth2.OAuth2
}

type ServiceConfig struct {
	ClientID     string
	ClientSecret string
}

// methods
//
func NewService(serviceConfig ServiceConfig, bigQueryService *bigquery.Service) (*Service, *errortools.Error) {
	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	service := Service{
		clientID:     serviceConfig.ClientID,
		clientSecret: serviceConfig.ClientSecret,
	}

	getTokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return google.GetToken(apiName, serviceConfig.ClientID, bigQueryService)
	}

	saveTokenFunction := func(token *oauth2.Token) *errortools.Error {
		return google.SaveToken(apiName, serviceConfig.ClientID, token, bigQueryService)
	}

	newTokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return service.GetAccessToken()
	}

	oAuth2Config := oauth2.OAuth2Config{
		GetTokenFunction:  &getTokenFunction,
		SaveTokenFunction: &saveTokenFunction,
		NewTokenFunction:  &newTokenFunction,
	}
	service.oAuth2 = oauth2.NewOAuth(oAuth2Config)
	return &service, nil
}

func (service *Service) ValidateToken() (*oauth2.Token, *errortools.Error) {
	return service.oAuth2.ValidateToken()
}

func ParseDateString(date string) *time.Time {
	if len(date) >= 19 {
		d, err := time.Parse("2006-01-02T15:04:05", date[:19])
		if err == nil {
			return &d
		}
	}

	return nil
}

// generic Get method
//
func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2.Get(requestConfig)
}

// generic Post method
//
func (service *Service) post(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2.Post(requestConfig)
}

// generic Put method
//
func (service *Service) put(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2.Put(requestConfig)
}

// generic Patch method
//
func (service *Service) patch(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2.Patch(requestConfig)
}

// generic Delete method
//
func (service *Service) delete(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2.Delete(requestConfig)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURLData, path)
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig, skipAccessToken bool) (*http.Request, *http.Response, *errortools.Error) {
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.oAuth2.HTTPRequest(httpMethod, requestConfig, skipAccessToken)

	if e != nil {
		if errorResponse.Error.Message != "" {
			e.SetMessage(errorResponse.Error.Message)
		}

		b, _ := json.Marshal(errorResponse)
		e.SetExtra("error", string(b))
	}

	return request, response, e
}
