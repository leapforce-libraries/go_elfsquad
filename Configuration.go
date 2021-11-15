package elfsquad

import (
	"fmt"
	"net/http"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type ConfigurationsResponse struct {
	Context  string          `json:"@odata.context"`
	Value    []Configuration `json:"value"`
	NextLink string          `json:"@odata.nextLink"`
}

type Configuration struct {
	Code                string                 `json:"code"`
	CurrencyISO         string                 `json:"currencyIso"`
	LanguageISO         string                 `json:"languageIso"`
	Preview             bool                   `json:"preview"`
	FeatureModelID      types.GUID             `json:"featureModelId"`
	FeatureModelVersion int32                  `json:"featureModelVersion"`
	ID                  types.GUID             `json:"id"`
	CreatedDate         e_types.DateTimeString `json:"createdDate"`
	UpdatedDate         e_types.DateTimeString `json:"updatedDate"`
	OrganizationID      *types.GUID            `json:"organizationId"`
	Reference           *string                `json:"reference"`
	CreatorID           types.GUID             `json:"creatorId"`
	CustomField1        *string                `json:"customField1"`
	CustomField2        *string                `json:"customField2"`
	CustomField3        *string                `json:"customField3"`
	CustomField4        *string                `json:"customField4"`
	CustomField5        *string                `json:"customField5"`
}

func (service *Service) GetConfigurations() (*[]Configuration, *errortools.Error) {
	top := 100
	skip := 0

	configurations := []Configuration{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("configurations?$top=%v&$skip=%v", top, skip)

		configurationsResponse := ConfigurationsResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.urlData(urlPath),
			ResponseModel: &configurationsResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(configurationsResponse.Value)

		if rowCount > 0 {
			configurations = append(configurations, configurationsResponse.Value...)
		}

		skip += top
	}

	return &configurations, nil
}
