package elfsquad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	oauth2 "github.com/Leapforce-nl/go_oauth2"
	types "github.com/Leapforce-nl/go_types"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
	//ExpiresIn   json.RawMessage `json:"expiresIn"`
}

func (es *Elfsquad) GetAccessToken() (*oauth2.Token, error) {
	client := new(http.Client)

	urlString := accessTokenURL

	data := make(map[string]string)
	data["clientId"] = es.clientID
	data["secret"] = es.secret

	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(accessTokenMethod, urlString, bytes.NewReader(dataByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Send out the HTTP request
	response, err := client.Do(req)

	// Check HTTP StatusCode
	if response.StatusCode < 200 || response.StatusCode > 299 {
		fmt.Println(fmt.Sprintf("ERROR in %s", accessTokenMethod))
		fmt.Println("url", urlString)
		fmt.Println("StatusCode", response.StatusCode)

		message := fmt.Sprintf("Server returned statuscode %v", response.StatusCode)
		return nil, &types.ErrorString{message}
	}
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	accessToken := AccessToken{}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &accessToken)
	if err != nil {
		return nil, err
	}

	expiresIn, _ := json.Marshal(accessToken.ExpiresIn / 1000)
	expiresInJson := json.RawMessage(expiresIn)

	token := oauth2.Token{
		AccessToken: &accessToken.AccessToken,
		ExpiresIn:   &expiresInJson,
	}

	return &token, nil
}
