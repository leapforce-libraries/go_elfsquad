package elfsquad

import (
	"net/http"
	"time"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	errortools "github.com/leapforce-libraries/go_errortools"

	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	APIName              string = "Elfsquad"
	APIURLData           string = "https://api.elfsquad.io/data/1"
	AccessTokenURL       string = "https://login.elfsquad.io/connect/token"
	AccessTokenMethod    string = http.MethodPost
	AccessTokenGrantType string = "client_credentials"
	AccessTokenScope     string = "Elfskot.Api"
)

// Elfsquad stores Elfsquad configuration
//
type Elfsquad struct {
	clientID     string
	clientSecret string
	oAuth2       *oauth2.OAuth2
}

// methods
//
func NewElfsquad(clientID string, clientSecret string, bigQuery *bigquerytools.BigQuery) (*Elfsquad, *errortools.Error) {
	es := Elfsquad{clientID: clientID, clientSecret: clientSecret}

	tokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return es.GetAccessToken()
	}

	config := oauth2.OAuth2Config{
		APIName:       APIName,
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		TokenFunction: &tokenFunction,
	}
	es.oAuth2 = oauth2.NewOAuth(config, bigQuery)
	return &es, nil
}

func (es *Elfsquad) ValidateToken() (*oauth2.Token, *errortools.Error) {
	return es.oAuth2.ValidateToken()
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
