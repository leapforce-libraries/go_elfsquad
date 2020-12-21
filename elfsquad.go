package elfsquad

import (
	"net/http"
	"time"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	errortools "github.com/leapforce-libraries/go_errortools"

	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	apiName    string = "Elfsquad"
	apiURLData string = "https://api.elfsquad.io/data/1"
	//apiURL            string = "https://api.elfsquad.io/api/2"
	accessTokenURL string = "https://api.elfsquad.io/api/2/auth/elfskotconnectlogin"
	//accessTokenURL    string = "https://login.elfsquad.io/connect/token"
	accessTokenMethod string = http.MethodPost
)

// Elfsquad stores Elfsquad configuration
//
type Elfsquad struct {
	clientID string
	secret   string
	oAuth2   *oauth2.OAuth2
}

// methods
//
func NewElfsquad(clientID string, secret string, bigQuery *bigquerytools.BigQuery) (*Elfsquad, *errortools.Error) {
	es := Elfsquad{clientID: clientID, secret: secret}

	tokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return es.GetAccessToken()
	}

	config := oauth2.OAuth2Config{
		APIName:       apiName,
		ClientID:      clientID,
		ClientSecret:  secret,
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
