package elfsquad

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type QuotationPropertiesResponse struct {
	Context  string              `json:"@odata.context"`
	Value    []QuotationProperty `json:"value"`
	NextLink string              `json:"@odata.nextLink"`
}

type QuotationProperty struct {
	Description    string     `json:"description"`
	IsRequired     bool       `json:"isRequired"`
	IsReadonly     bool       `json:"isReadonly"`
	Type           *string    `json:"type"`
	Order          int32      `json:"order"`
	ID             types.GUID `json:"id"`
	CreatedDate    string     `json:"createdDate"`
	UpdatedDate    string     `json:"updatedDate"`
	OrganizationID types.GUID `json:"organizationId"`
	CreatorID      types.GUID `json:"creatorId"`
	CustomField1   string     `json:"customField1"`
	CustomField2   string     `json:"customField2"`
	CustomField3   string     `json:"customField3"`
	CustomField4   string     `json:"customField4"`
	CustomField5   string     `json:"customField5"`
}

func (service *Service) GetQuotationProperties() (*[]QuotationProperty, *errortools.Error) {
	top := 100
	skip := 0

	quotationProperties := []QuotationProperty{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("quotationproperties?$top=%v&$skip=%v", top, skip)

		quotationPropertiesResponse := QuotationPropertiesResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.urlData(urlPath),
			ResponseModel: &quotationPropertiesResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(quotationPropertiesResponse.Value)

		if rowCount > 0 {
			quotationProperties = append(quotationProperties, quotationPropertiesResponse.Value...)
		}

		skip += top
	}

	return &quotationProperties, nil
}
