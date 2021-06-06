package elfsquad

import (
	"fmt"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type OrganizationsResponse struct {
	Context  string         `json:"@odata.context"`
	Value    []Organization `json:"value"`
	NextLink string         `json:"@odata.nextLink"`
}

type Organization struct {
	Name                          string                 `json:"name"`
	SettingsID                    types.GUID             `json:"settingsId"`
	DefaultDiscountPct            float64                `json:"defaultDiscountPct"`
	DefaultUpValuePct             float64                `json:"defaultUpValuePct"`
	CurrencyISO                   string                 `json:"currencyIso"`
	LanguageISO                   string                 `json:"languageIso"`
	IsConfigurationModelPreviewer bool                   `json:"isConfigurationModelPreviewer"`
	OrganizationTypeID            types.GUID             `json:"organizationTypeId"`
	ID                            types.GUID             `json:"id"`
	CreatorID                     types.GUID             `json:"creatorId"`
	Reference                     string                 `json:"reference"`
	Synced                        bool                   `json:"synced"`
	Inactive                      bool                   `json:"inactive"`
	CreatedDate                   e_types.DateTimeString `json:"createdDate"`
	UpdatedDate                   e_types.DateTimeString `json:"updatedDate"`
}

func (service *Service) GetOrganizations() (*[]Organization, *errortools.Error) {
	top := 100
	skip := 0

	organizations := []Organization{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("Organizations?$top=%v&$skip=%v", top, skip)

		organizationsResponse := OrganizationsResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           service.urlData(urlPath),
			ResponseModel: &organizationsResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(organizationsResponse.Value)

		if rowCount > 0 {
			organizations = append(organizations, organizationsResponse.Value...)
		}

		skip += top
	}

	return &organizations, nil
}
