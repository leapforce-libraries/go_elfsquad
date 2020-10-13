package elfsquad

import (
	"net/http"

	bigquerytools "github.com/Leapforce-nl/go_bigquerytools"

	oauth2 "github.com/Leapforce-nl/go_oauth2"
)

const (
	apiName         string = "Elfsquad"
	apiURLData      string = "https://api.elfsquad.io/data/1"
	apiURL          string = "https://api.elfsquad.io/api/2"
	tokenURL        string = "https://api.elfsquad.io/api/2/auth/elfskotconnectlogin"
	tokenHTTPMethod string = http.MethodPost
)

// Elfsquad stores Elfsquad configuration
//
type Elfsquad struct {
	oAuth2 *oauth2.OAuth2
}

// methods
//
func NewElfsquad(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery, isLive bool) (*Elfsquad, error) {
	es := Elfsquad{}
	es.oAuth2 = oauth2.NewOAuth(apiName, clientID, clientSecret, scope, "", "", tokenURL, tokenHTTPMethod, bigQuery, isLive)
	return &es, nil
}

func (es *Elfsquad) ValidateToken() (*oauth2.Token, error) {
	return es.oAuth2.ValidateToken()
}
