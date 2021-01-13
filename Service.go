package elfsquad

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	APIName              string = "Service"
	APIURLData           string = "https://api.elfsquad.io/data/1"
	AccessTokenURL       string = "https://login.elfsquad.io/connect/token"
	AccessTokenMethod    string = http.MethodPost
	AccessTokenGrantType string = "client_credentials"
	AccessTokenScope     string = "Elfskot.Api"
)

// Service stores Service configuration
//
type Service struct {
	clientID     string
	clientSecret string
	oAuth2       *oauth2.OAuth2
}

// methods
//
func NewService(clientID string, clientSecret string, bigQuery *google.BigQuery) (*Service, *errortools.Error) {
	service := Service{clientID: clientID, clientSecret: clientSecret}

	tokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return service.GetAccessToken()
	}

	config := oauth2.OAuth2Config{
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		NewTokenFunction: &tokenFunction,
	}
	service.oAuth2 = oauth2.NewOAuth(config)
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
func (service *Service) get(urlPath string, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, urlPath, nil, responseModel)
}

// generic Post method
//
func (service *Service) post(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, urlPath, bodyModel, responseModel)
}

// generic Put method
//
func (service *Service) put(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, urlPath, bodyModel, responseModel)
}

// generic Patch method
//
func (service *Service) patch(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPatch, urlPath, bodyModel, responseModel)
}

// generic Delete method
//
func (service *Service) delete(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, urlPath, bodyModel, responseModel)
}

func (service *Service) httpRequest(httpMethod string, urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	url := fmt.Sprintf("%s/%s", APIURLData, urlPath)
	//fmt.Println(url)

	e := new(errortools.Error)

	errorResponse := ErrorResponse{}

	request, response, e := service.oAuth2.HTTP(httpMethod, url, bodyModel, responseModel, &errorResponse)

	if e != nil {
		if errorResponse.Error.Message != "" {
			e.SetMessage(errorResponse.Error.Message)
		}

		b, _ := json.Marshal(errorResponse)
		e.SetExtra("error", string(b))

		return nil, nil, e
	}

	return request, response, e
}
