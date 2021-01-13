package elfsquad

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (service *Service) GetAccessToken() (*oauth2.Token, *errortools.Error) {
	client := new(http.Client)

	urlString := AccessTokenURL

	data := url.Values{}
	data.Set("client_id", service.clientID)
	data.Set("client_secret", service.clientSecret)
	data.Set("grant_type", AccessTokenGrantType)
	data.Set("scope", AccessTokenScope)

	e := new(errortools.Error)

	request, err := http.NewRequest(AccessTokenMethod, urlString, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	e.SetRequest(request)

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	// Send out the HTTP request
	response, err := client.Do(request)

	// Check HTTP StatusCode
	if response.StatusCode < 200 || response.StatusCode > 299 {
		fmt.Println(fmt.Sprintf("ERROR in %s", AccessTokenMethod))
		fmt.Println("url", urlString)
		fmt.Println("StatusCode", response.StatusCode)

		e.SetResponse(response)
		e.SetMessage(fmt.Sprintf("Server returned statuscode %v", response.StatusCode))
		return nil, e
	} else if err != nil {
		e.SetResponse(response)
		e.SetMessage(err)
		return nil, e
	}

	defer response.Body.Close()

	accessToken := AccessToken{}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		e.SetMessage(err)
		return nil, e
	}

	err = json.Unmarshal(b, &accessToken)
	if err != nil {
		e.SetMessage(err)
		return nil, e
	}

	expiresIn, _ := json.Marshal(accessToken.ExpiresIn)
	expiresInJson := json.RawMessage(expiresIn)

	token := oauth2.Token{
		AccessToken: &accessToken.AccessToken,
		ExpiresIn:   &expiresInJson,
		TokenType:   &accessToken.TokenType,
	}

	return &token, nil
}
