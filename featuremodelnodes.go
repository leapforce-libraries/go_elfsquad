package elfsquad

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
)

type FeatureModelNodesResponse struct {
	Context  string             `json:"@odata.context"`
	Value    []FeatureModelNode `json:"value"`
	NextLink string             `json:"@odata.nextLink"`
}

type FeatureModelNode struct {
	FeatureModelID               types.GUID `json:"featureModelId"`
	FeatureID                    types.GUID `json:"featureId"`
	HideInQuotation              bool       `json:"hideInQuotation"`
	HideInConfigurator           bool       `json:"hideInConfigurator"`
	HideInOrderEntry             bool       `json:"hideInOrderEntry"`
	HideInOverview               bool       `json:"hideInOverview"`
	IsQuotationGroup             bool       `json:"isQuotationGroup"`
	IsMandatory                  bool       `json:"isMandatory"`
	IsPreconfiguration           bool       `json:"isPreconfiguration"`
	IsPhantom                    bool       `json:"isPhantom"`
	UnitPriceIncVAT              float64    `json:"unitPriceIncVAT"`
	UnitPriceExVAT               float64    `json:"unitPriceExVAT"`
	UnitPriceExVATExExchangeRate float64    `json:"unitPriceExVATExExchangeRate"`
	TotalPriceIncVAT             float64    `json:"totalPriceIncVAT"`
	TotalPriceExVAT              float64    `json:"totalPriceExVAT"`
	UnitPriceIncVATLabel         float64    `json:"unitPriceIncVATLabel"`
	UnitPriceExVATLabel          float64    `json:"unitPriceExVATLabel"`
	TotalPriceIncVATLabel        float64    `json:"totalPriceIncVATLabel"`
	TotalPriceExVATLabel         float64    `json:"totalPriceExVATLabel"`
	ID                           types.GUID `json:"id"`
	CreatorID                    types.GUID `json:"creatorId"`
	Synced                       bool       `json:"synced"`
	Inactive                     bool       `json:"inactive"`
	CreatedDate                  string     `json:"createdDate"`
	UpdatedDate                  string     `json:"updatedDate"`
}

func (es *Elfsquad) GetFeatureModelNodes() (*[]FeatureModelNode, *errortools.Error) {
	top := 100
	skip := 0

	featureModelNodes := []FeatureModelNode{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		url := fmt.Sprintf("%s/featuremodelnodes?$top=%v&$skip=%v", APIURLData, top, skip)

		featureModelNodesReponse := FeatureModelNodesResponse{}

		_, _, e := es.oAuth2.Get(url, &featureModelNodesReponse, nil)
		if e != nil {
			return nil, e
		}

		rowCount = len(featureModelNodesReponse.Value)

		if rowCount > 0 {
			featureModelNodes = append(featureModelNodes, featureModelNodesReponse.Value...)
		}

		skip += top
	}

	return &featureModelNodes, nil
}
