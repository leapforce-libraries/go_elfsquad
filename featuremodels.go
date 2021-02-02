package elfsquad

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type FeatureModelsResponse struct {
	Context  string         `json:"@odata.context"`
	Value    []FeatureModel `json:"value"`
	NextLink string         `json:"@odata.nextLink"`
}

type FeatureModel struct {
	RootFeatureID    types.GUID `json:"rootFeatureId"`
	Order            int32      `json:"order"`
	DisplayPrices    bool       `json:"displayPrices"`
	HideInShowroom   bool       `json:"hideInShowroom"`
	HideInOrderEntry bool       `json:"hideInOrderEntry"`
	AutodeskUrn      string     `json:"autodeskUrn"`
	ID               types.GUID `json:"id"`
	CreatedDate      string     `json:"createdDate"`
	UpdatedDate      string     `json:"updatedDate"`
	OrganizationID   types.GUID `json:"organizationId"`
	CreatorID        types.GUID `json:"creatorId"`
}

func (service *Service) GetFeatureModels() (*[]FeatureModel, *errortools.Error) {
	top := 100
	skip := 0

	featureModels := []FeatureModel{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("featuremodels?$top=%v&$skip=%v", top, skip)

		featureModelsResponse := FeatureModelsResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           service.url(urlPath),
			ResponseModel: &featureModelsResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(featureModelsResponse.Value)

		if rowCount > 0 {
			featureModels = append(featureModels, featureModelsResponse.Value...)
		}

		skip += top
	}

	return &featureModels, nil
}
