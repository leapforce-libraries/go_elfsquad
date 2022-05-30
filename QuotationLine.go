package elfsquad

import (
	"fmt"
	"net/http"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type QuotationLinesResponse struct {
	Context  string          `json:"@odata.context"`
	Value    []QuotationLine `json:"value"`
	NextLink string          `json:"@odata.nextLink"`
}

type QuotationLine struct {
	QuotationID                     types.Guid              `json:"quotationId"`
	LineNumber                      int32                   `json:"lineNumber"`
	ArticleCode                     *string                 `json:"articleCode"`
	DeliveryDate                    *e_types.DateTimeString `json:"deliverydate"`
	Description                     string                  `json:"description"`
	FeatureID                       types.Guid              `json:"featureId"`
	FeatureModelNodeID              types.Guid              `json:"featureModelNodeId"`
	Quantity                        float64                 `json:"quantity"`
	ImageValue                      *string                 `json:"imageValue"`
	TextValue                       *string                 `json:"textValue"`
	ParentLineID                    *types.Guid             `json:"parentLineId"`
	GroupedRootLine                 bool                    `json:"groupedRootLine"`
	GroupID                         *types.Guid             `json:"groupId"`
	ParentGroupID                   *types.Guid             `json:"parentGroupId"`
	GroupOrder                      int32                   `json:"groupOrder"`
	GroupTitle                      *string                 `json:"groupTitle"`
	AddedFromConfiguration          bool                    `json:"addedFromConfiguration"`
	ConfigurationID                 *types.Guid             `json:"configurationId"`
	VatID                           *types.Guid             `json:"vatId"`
	DiscountPct                     float64                 `json:"discountPct"`
	MarginPct                       float64                 `json:"marginPct"`
	PurchasePriceDiscountPct        float64                 `json:"purchasePriceDiscountPct"`
	GroupDiscountPct                float64                 `json:"groupDiscountPct"`
	DefaultPurchasePriceDiscountPct float64                 `json:"defaultPurchasePriceDiscountPct"`
	UnitPrice                       float64                 `json:"unitPrice"`
	OriginalUnitPrice               float64                 `json:"originalUnitPrice"`
	ID                              types.Guid              `json:"id"`
	CreatedDate                     e_types.DateTimeString  `json:"createdDate"`
	UpdatedDate                     e_types.DateTimeString  `json:"updatedDate"`
	OrganizationID                  *types.Guid             `json:"organizationId"`
	Reference                       *string                 `json:"reference"`
	CreatorID                       types.Guid              `json:"creatorId"`
	CustomField1                    *string                 `json:"customField1"`
	CustomField2                    *string                 `json:"customField2"`
	CustomField3                    *string                 `json:"customField3"`
	CustomField4                    *string                 `json:"customField4"`
	CustomField5                    *string                 `json:"customField5"`
}

func (service *Service) GetQuotationLines() (*[]QuotationLine, *errortools.Error) {
	top := 100
	skip := 0

	quotationLines := []QuotationLine{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("QuotationLines?$top=%v&$skip=%v", top, skip)

		quotationLinesResponse := QuotationLinesResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.urlData(urlPath),
			ResponseModel: &quotationLinesResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
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
