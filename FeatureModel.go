package elfsquad

import (
	"fmt"
	"net/http"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
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
	RootFeatureID          types.GUID             `json:"rootFeatureId"`
	Order                  int32                  `json:"order"`
	DisplayPrices          bool                   `json:"displayPrices"`
	HideInShowroom         bool                   `json:"hideInShowroom"`
	HideInOrderEntry       bool                   `json:"hideInOrderEntry"`
	StartingPricesInclVat  map[string]string      `json:"startingPricesInclVat"`
	StartingPricesExclVat  map[string]string      `json:"startingPricesExclVat"`
	AutodeskURN            string                 `json:"autodeskUrn"`
	ForeignAutodeskURNs    map[string]string      `json:"foreignAutodeskUrns"`
	ForeignAttachmentNodes map[string]string      `json:"foreignAttachmentNodes"`
	FeatureModelHash       string                 `json:"featureModelHash"`
	PurchaseDiscount       float64                `json:"purchaseDiscount"`
	ID                     types.GUID             `json:"id"`
	CreatorID              types.GUID             `json:"creatorId"`
	Synced                 bool                   `json:"synced"`
	Inactive               bool                   `json:"inactive"`
	CreatedDate            e_types.DateTimeString `json:"createdDate"`
	UpdatedDate            e_types.DateTimeString `json:"updatedDate"`
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
			Method:        http.MethodGet,
			URL:           service.urlData(urlPath),
			ResponseModel: &featureModelsResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
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
