package elfsquad

import (
	"fmt"
	"net/http"
	url "net/url"
	"strings"
	"time"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type QuotationsResponse struct {
	Context  string      `json:"@odata.context"`
	Value    []Quotation `json:"value"`
	NextLink string      `json:"@odata.nextLink"`
}

type Quotation struct {
	SellerID            *types.Guid             `json:"sellerId,omitempty"`
	SellerContactID     *types.Guid             `json:"sellerContactId,omitempty"`
	DebtorID            *types.Guid             `json:"debtorId,omitempty"`
	DebtorContactID     *types.Guid             `json:"debtorContactId,omitempty"`
	ShipToID            *types.Guid             `json:"shipToId,omitempty"`
	ShipToContactID     *types.Guid             `json:"shipToContactId,omitempty"`
	LanguageISO         string                  `json:"languageIso,omitempty"`
	CurrencyISO         string                  `json:"currencyIso,omitempty"`
	Synced              bool                    `json:"synced,omitempty"`
	QuotationNumber     int64                   `json:"quotationNumber,omitempty"`
	VersionNumber       int32                   `json:"versionNumber,omitempty"`
	Status              *string                 `json:"status,omitempty"`
	StatusID            *types.Guid             `json:"statusId,omitempty"`
	Subject             *string                 `json:"subject,omitempty"`
	TotalPrice          float64                 `json:"totalPrice,omitempty"`
	IsVerified          *bool                   `json:"isVerified,omitempty"`
	CustomerReference   *string                 `json:"customerReference,omitempty"`
	QuotationReference  *string                 `json:"quotationReference,omitempty"`
	Deliverydate        *e_types.DateTimeString `json:"deliverydate,omitempty"`
	Remarks             *string                 `json:"remarks,omitempty"`
	ExpiresDate         *e_types.DateTimeString `json:"expiresDate,omitempty"`
	QuotationTemplateID *types.Guid             `json:"quotationTemplateId,omitempty"`
	ID                  types.Guid              `json:"id,omitempty"`
	CreatedDate         e_types.DateTimeString  `json:"createdDate,omitempty"`
	UpdatedDate         e_types.DateTimeString  `json:"updatedDate,omitempty"`
	OrganizationID      *types.Guid             `json:"organizationId,omitempty"`
	Reference           *string                 `json:"reference,omitempty"`
	CreatorID           types.Guid              `json:"creatorId,omitempty"`
	CustomField1        *string                 `json:"customField1,omitempty"`
	CustomField2        *string                 `json:"customField2,omitempty"`
	CustomField3        *string                 `json:"customField3,omitempty"`
	CustomField4        *string                 `json:"customField4,omitempty"`
	CustomField5        *string                 `json:"customField5,omitempty"`
}

type GetQuotationsParams struct {
	QuotationNumber *int64
	CreatedAfter    *time.Time
	UpdatedAfter    *time.Time
	Status          *string
	Select          *[]string
	OrderBy         *[]string
}

func (service *Service) GetQuotations(params *GetQuotationsParams) (*[]Quotation, *errortools.Error) {
	top := 100
	skip := 0

	filter := []string{}

	if params != nil {
		if params.QuotationNumber != nil {
			filter = append(filter, fmt.Sprintf("QuotationNumber eq %v", *params.QuotationNumber))
		}
		if params.CreatedAfter != nil {
			filter = append(filter, fmt.Sprintf("CreatedDate gt %s", params.CreatedAfter.Format(time.RFC3339)))
		}
		if params.UpdatedAfter != nil {
			filter = append(filter, fmt.Sprintf("UpdatedDate gt %s", params.UpdatedAfter.Format(time.RFC3339)))
		}
		if params.Status != nil {
			filter = append(filter, fmt.Sprintf("Status eq '%s'", *params.Status))
		}
	}

	quotations := []Quotation{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("quotations?$top=%v&$skip=%v", top, skip)

		if len(filter) > 0 {
			urlPath = fmt.Sprintf("%s&$filter=%s", urlPath, url.QueryEscape(strings.Join(filter, " AND ")))
		}
		if params != nil {
			if params.Select != nil {
				urlPath = fmt.Sprintf("%s&$select=%s", urlPath, strings.Join(*params.Select, ","))
			}
			if params.OrderBy != nil {
				urlPath = fmt.Sprintf("%s&$orderby=%s", urlPath, strings.Join(*params.OrderBy, ","))
			}
		}

		quotationsResponse := QuotationsResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.urlData(urlPath),
			ResponseModel: &quotationsResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(quotationsResponse.Value)

		if rowCount > 0 {
			quotations = append(quotations, quotationsResponse.Value...)
		}

		skip += top
	}

	return &quotations, nil
}

func (service *Service) UpdateQuotation(quotationID types.Guid, quotationUpdate *Quotation) *errortools.Error {
	urlPath := fmt.Sprintf("quotations(%s)", quotationID.String())

	requestConfig := go_http.RequestConfig{
		Method:    http.MethodPatch,
		Url:       service.urlData(urlPath),
		BodyModel: quotationUpdate,
	}
	_, _, e := service.httpRequest(&requestConfig)

	return e
}
