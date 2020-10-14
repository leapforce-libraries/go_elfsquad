package elfsquad

import (
	"fmt"

	types "github.com/Leapforce-nl/go_types"
)

type QuotationLinesResponse struct {
	Context  string          `json:"@odata.context"`
	Value    []QuotationLine `json:"value"`
	NextLink string          `json:"@odata.nextLink"`
}

type QuotationLine struct {
	QuotationID                     types.GUID `json:"quotationId"`
	LineNumber                      int32      `json:"lineNumber"`
	ArticleCode                     string     `json:"articleCode"`
	Deliverydate                    string     `json:"deliverydate"`
	Description                     string     `json:"description"`
	FeatureID                       types.GUID `json:"featureId"`
	FeatureModelNodeID              types.GUID `json:"featureModelNodeId"`
	Quantity                        float64    `json:"quantity"`
	ImageValue                      string     `json:"imageValue"`
	TextValue                       string     `json:"textValue"`
	ParentLineID                    types.GUID `json:"parentLineId"`
	GroupedRootLine                 bool       `json:"groupedRootLine"`
	GroupID                         types.GUID `json:"groupId"`
	GroupOrder                      int32      `json:"groupOrder"`
	GroupTitle                      string     `json:"groupTitle"`
	AddedFromConfiguration          bool       `json:"addedFromConfiguration"`
	ConfigurationID                 types.GUID `json:"configurationId"`
	VatID                           types.GUID `json:"vatId"`
	DiscountPct                     float64    `json:"discountPct"`
	MarginPct                       float64    `json:"marginPct"`
	PurchasePriceDiscountPct        float64    `json:"purchasePriceDiscountPct"`
	GroupDiscountPct                float64    `json:"groupDiscountPct"`
	DefaultPurchasePriceDiscountPct float64    `json:"defaultPurchasePriceDiscountPct"`
	UnitPrice                       float64    `json:"unitPrice"`
	OriginalUnitPrice               float64    `json:"originalUnitPrice"`
	ID                              types.GUID `json:"id"`
	CreatedDate                     string     `json:"createdDate"`
	UpdatedDate                     string     `json:"updatedDate"`
	OrganizationID                  types.GUID `json:"organizationId"`
	CreatorID                       types.GUID `json:"creatorId"`
}

func (es *Elfsquad) GetQuotationLines() (*[]QuotationLine, error) {
	top := 100
	skip := 0

	quotationLines := []QuotationLine{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		url := fmt.Sprintf("%s/quotationlines?$top=%v&$skip=%v", apiURLData, top, skip)

		quotationLinesReponse := QuotationLinesResponse{}

		_, err := es.oAuth2.Get(url, &quotationLinesReponse)
		if err != nil {
			return nil, err
		}

		rowCount = len(quotationLinesReponse.Value)

		if rowCount > 0 {
			quotationLines = append(quotationLines, quotationLinesReponse.Value...)
		}

		skip += top
	}

	return &quotationLines, nil
}
