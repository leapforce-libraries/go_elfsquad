package elfsquad

import (
	"encoding/json"
	"fmt"
	"net/http"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type CRMAccountsResponse struct {
	Context  string       `json:"@odata.context"`
	Value    []CRMAccount `json:"value"`
	NextLink string       `json:"@odata.nextLink"`
}

type CRMAccount struct {
	FullName    string                 `json:"fullName"`
	CompanyName string                 `json:"companyName"`
	StreetName  string                 `json:"streetName"`
	City        string                 `json:"city"`
	CountryISO  string                 `json:"countryIso"`
	Type        json.RawMessage        `json:"type"`
	ID          types.GUID             `json:"id"`
	CreatorID   types.GUID             `json:"creatorId"`
	Reference   string                 `json:"reference"`
	Synced      bool                   `json:"synced"`
	Inactive    bool                   `json:"inactive"`
	CreatedDate e_types.DateTimeString `json:"createdDate"`
	UpdatedDate e_types.DateTimeString `json:"updatedDate"`
}

func (service *Service) GetCRMAccounts() (*[]CRMAccount, *errortools.Error) {
	top := 100
	skip := 0

	crmAccounts := []CRMAccount{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("CrmAccounts?$top=%v&$skip=%v", top, skip)

		crmAccountsResponse := CRMAccountsResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.urlData(urlPath),
			ResponseModel: &crmAccountsResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(crmAccountsResponse.Value)

		if rowCount > 0 {
			crmAccounts = append(crmAccounts, crmAccountsResponse.Value...)
		}

		skip += top
	}

	return &crmAccounts, nil
}
