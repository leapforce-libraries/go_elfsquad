package elfsquad

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	types "github.com/leapforce-libraries/go_types"
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

func (service *Service) GetQuotationLines() (*[]QuotationLine, *errortools.Error) {
	top := 100
	skip := 0

	quotationLines := []QuotationLine{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("quotationlines?$top=%v&$skip=%v", top, skip)

		quotationLinesResponse := QuotationLinesResponse{}
		requestConfig := oauth2.RequestConfig{
			URL:           service.url(urlPath),
			ResponseModel: &quotationLinesResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(quotationLinesResponse.Value)

		if rowCount > 0 {
			quotationLines = append(quotationLines, quotationLinesResponse.Value...)
		}

		skip += top
	}

	return &quotationLines, nil
}
