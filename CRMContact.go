package elfsquad

import (
	"fmt"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type CRMContactsResponse struct {
	Context  string       `json:"@odata.context"`
	Value    []CRMContact `json:"value"`
	NextLink string       `json:"@odata.nextLink"`
}

type CRMContact struct {
	FullName         string                 `json:"fullName"`
	FullLastName     string                 `json:"fullLastName"`
	FirstName        string                 `json:"firstName"`
	LastName         string                 `json:"lastName"`
	Email            string                 `json:"email"`
	UseParentAddress bool                   `json:"useParentAddress"`
	CRMAccountID     types.GUID             `json:"crmAccountId"`
	ID               types.GUID             `json:"id"`
	CreatorID        types.GUID             `json:"creatorId"`
	Reference        string                 `json:"reference"`
	Synced           bool                   `json:"synced"`
	Inactive         bool                   `json:"inactive"`
	CreatedDate      e_types.DateTimeString `json:"createdDate"`
	UpdatedDate      e_types.DateTimeString `json:"updatedDate"`
}

func (service *Service) GetCRMContacts() (*[]CRMContact, *errortools.Error) {
	top := 100
	skip := 0

	crmContacts := []CRMContact{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("CrmContacts?$top=%v&$skip=%v", top, skip)

		crmContactsResponse := CRMContactsResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           service.urlData(urlPath),
			ResponseModel: &crmContactsResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(crmContactsResponse.Value)

		if rowCount > 0 {
			crmContacts = append(crmContacts, crmContactsResponse.Value...)
		}

		skip += top
	}

	return &crmContacts, nil
}
