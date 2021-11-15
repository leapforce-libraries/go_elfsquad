package elfsquad

import (
	"net/http"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

/*type QuotationTemplatesResponse struct {
	Context  string              `json:"@odata.context"`
	Value    []QuotationTemplate `json:"value"`
	NextLink string              `json:"@odata.nextLink"`
}*/

type QuotationTemplate struct {
	Name                  string                 `json:"name"`
	URL                   string                 `json:"url"`
	IsDefault             bool                   `json:"isDefault"`
	LanguageISO           string                 `json:"languageIso"`
	TenantDefaultLanguage string                 `json:"tenantDefaultLanguage"`
	ID                    types.GUID             `json:"id"`
	CreatorID             types.GUID             `json:"creatorId"`
	Synced                bool                   `json:"synced"`
	Inactive              bool                   `json:"inactive"`
	CreatedDate           e_types.DateTimeString `json:"createdDate"`
	UpdatedDate           e_types.DateTimeString `json:"updatedDate"`
}

func (service *Service) GetQuotationTemplates() (*[]QuotationTemplate, *errortools.Error) {
	quotationTemplates := []QuotationTemplate{}

	urlPath := "QuotationTemplates"

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url(urlPath),
		ResponseModel: &quotationTemplates,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &quotationTemplates, nil
}
