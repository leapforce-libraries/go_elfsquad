package elfsquad

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	o_token "github.com/leapforce-libraries/go_oauth2/token"
)

const (
	apiName              string = "Elfsquad"
	apiURLData           string = "https://api.elfsquad.io/data/1"
	apiURL               string = "https://api.elfsquad.io/api/2"
	accessTokenURL       string = "https://api.elfsquad.io/api/2/auth/elfskotconnectlogin"
	accessTokenMethod    string = http.MethodPost
	accessTokenGrantType string = "client_credentials"
	accessTokenScope     string = "Elfskot.Api"
)

// Service stores Service configuration
//
type Service struct {
	clientID      string
	clientSecret  string
	oAuth2Service *oauth2.Service
}

type ServiceConfig struct {
	ClientID     string
	ClientSecret string
}

// methods
//
func NewService(serviceConfig *ServiceConfig, bigQueryService *bigquery.Service) (*Service, *errortools.Error) {
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

	tokenSource, e := NewTokenSource(&service)
	if e != nil {
		return nil, e
	}

	oAuth2ServiceConfig := oauth2.ServiceConfig{
		TokenSource: tokenSource,
	}
	oauth2Service, e := oauth2.NewService(&oAuth2ServiceConfig)
	if e != nil {
		return nil, e
	}
	service.oAuth2Service = oauth2Service

	return &service, nil
}

func (service *Service) ValidateToken() (*o_token.Token, *errortools.Error) {
	return service.oAuth2Service.ValidateToken()
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

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
}

func (service *Service) urlData(path string) string {
	return fmt.Sprintf("%s/%s", apiURLData, path)
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.oAuth2Service.HttpRequest(requestConfig)

	if e != nil {
		if errorResponse.Error.Message != "" {
			e.SetMessage(errorResponse.Error.Message)
		}

		b, _ := json.Marshal(errorResponse)
		e.SetExtra("error", string(b))
	}

	return request, response, e
}

func (service *Service) httpRequestWithoutAccessToken(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.oAuth2Service.HttpRequestWithoutAccessToken(requestConfig)

	if e != nil {
		if errorResponse.Error.Message != "" {
			e.SetMessage(errorResponse.Error.Message)
		}

		b, _ := json.Marshal(errorResponse)
		e.SetExtra("error", string(b))
	}

	return request, response, e
}

func (service Service) ApiName() string {
	return apiName
}

func (service Service) ApiKey() string {
	return service.clientID
}

func (service Service) ApiCallCount() int64 {
	return service.oAuth2Service.ApiCallCount()
}

func (service Service) ApiReset() {
	service.oAuth2Service.ApiReset()
}
