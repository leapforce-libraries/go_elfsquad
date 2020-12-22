package elfsquad

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
)

type ConfigurationsResponse struct {
	Context  string          `json:"@odata.context"`
	Value    []Configuration `json:"value"`
	NextLink string          `json:"@odata.nextLink"`
}

type Configuration struct {
	Code                string     `json:"code"`
	CurrencyISO         string     `json:"currencyIso"`
	LanguageISO         string     `json:"languageIso"`
	Preview             bool       `json:"preview"`
	FeatureModelID      types.GUID `json:"featureModelId"`
	FeatureModelVersion int32      `json:"featureModelVersion"`
	ID                  types.GUID `json:"id"`
	CreatedDate         string     `json:"createdDate"`
	UpdatedDate         string     `json:"updatedDate"`
	OrganizationID      types.GUID `json:"organizationId"`
	CreatorID           types.GUID `json:"creatorId"`
}

func (es *Elfsquad) GetConfigurations() (*[]Configuration, *errortools.Error) {
	top := 100
	skip := 0

	configurations := []Configuration{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("configurations?$top=%v&$skip=%v", top, skip)

		configurationsReponse := ConfigurationsResponse{}
		_, _, e := es.get(urlPath, &configurationsReponse)
		if e != nil {
			return nil, e
		}

		rowCount = len(configurationsReponse.Value)

		if rowCount > 0 {
			configurations = append(configurations, configurationsReponse.Value...)
		}

		skip += top
	}

	return &configurations, nil
}
