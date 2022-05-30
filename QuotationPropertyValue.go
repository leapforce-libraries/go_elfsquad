package elfsquad

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type QuotationPropertyValuesResponse struct {
	Context  string                   `json:"@odata.context"`
	Value    []QuotationPropertyValue `json:"value"`
	NextLink string                   `json:"@odata.nextLink"`
}

type QuotationPropertyValue struct {
	EntityID         types.Guid `json:"entityId"`
	EntityPropertyID types.Guid `json:"entityPropertyId"`
	Value            string     `json:"value"`
	ID               types.Guid `json:"id"`
	CreatedDate      string     `json:"createdDate"`
	UpdatedDate      string     `json:"updatedDate"`
	OrganizationID   types.Guid `json:"organizationId"`
	CreatorID        types.Guid `json:"creatorId"`
	CustomField1     string     `json:"customField1"`
	CustomField2     string     `json:"customField2"`
	CustomField3     string     `json:"customField3"`
	CustomField4     string     `json:"customField4"`
	CustomField5     string     `json:"customField5"`
}

type GetQuotationPropertyValuesParams struct {
	EntityID         *types.Guid
	EntityPropertyID *types.Guid
	Select           *[]string
}

func (service *Service) GetQuotationPropertyValues(params *GetQuotationPropertyValuesParams) (*[]QuotationPropertyValue, *errortools.Error) {
	top := 100
	skip := 0

	filter := []string{}

	if params != nil {
		if params.EntityID != nil {
			filter = append(filter, fmt.Sprintf("EntityID eq %s", params.EntityID.String()))
		}
		if params.EntityPropertyID != nil {
			filter = append(filter, fmt.Sprintf("EntityPropertyID eq %s", params.EntityPropertyID.String()))
		}
	}

	quotationPropertyValues := []QuotationPropertyValue{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("quotationpropertyvalues?$top=%v&$skip=%v", top, skip)

		if len(filter) > 0 {
			urlPath = fmt.Sprintf("%s&$filter=%s", urlPath, url.QueryEscape(strings.Join(filter, " AND ")))
		}
		if params != nil {
			if params.Select != nil {
				urlPath = fmt.Sprintf("%s&$select=%s", urlPath, strings.Join(*params.Select, ","))
			}
		}

		quotationPropertyValuesResponse := QuotationPropertyValuesResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.urlData(urlPath),
			ResponseModel: &quotationPropertyValuesResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(quotationPropertyValuesResponse.Value)

		if rowCount > 0 {
			quotationPropertyValues = append(quotationPropertyValues, quotationPropertyValuesResponse.Value...)
		}

		skip += top
	}

	return &quotationPropertyValues, nil
}
