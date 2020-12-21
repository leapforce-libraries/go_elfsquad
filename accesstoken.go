package elfsquad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
	//ExpiresIn   json.RawMessage `json:"expiresIn"`
}

func (es *Elfsquad) GetAccessToken() (*oauth2.Token, *errortools.Error) {
	client := new(http.Client)

	urlString := accessTokenURL

	data := make(map[string]string)
	data["clientId"] = es.clientID
	data["secret"] = es.secret

	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	request, err := http.NewRequest(accessTokenMethod, urlString, bytes.NewReader(dataByte))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	// Send out the HTTP request
	response, err := client.Do(request)

	// Check HTTP StatusCode
	if response.StatusCode < 200 || response.StatusCode > 299 {
		fmt.Println(fmt.Sprintf("ERROR in %s", accessTokenMethod))
		fmt.Println("url", urlString)
		fmt.Println("StatusCode", response.StatusCode)

		e := new(errortools.Error)
		e = new(errortools.Error)
		e.SetRequest(request)
		e.SetResponse(response)

		e.SetMessage(fmt.Sprintf("Server returned statuscode %v", response.StatusCode))
	}
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	defer response.Body.Close()

	accessToken := AccessToken{}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	err = json.Unmarshal(b, &accessToken)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	expiresIn, _ := json.Marshal(accessToken.ExpiresIn / 1000)
	expiresInJson := json.RawMessage(expiresIn)

	token := oauth2.Token{
		AccessToken: &accessToken.AccessToken,
		ExpiresIn:   &expiresInJson,
	}

	return &token, nil
}
