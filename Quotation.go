package elfsquad

import (
	"fmt"
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
	SellerID            *types.GUID             `json:"sellerId,omitempty"`
	SellerContactID     *types.GUID             `json:"sellerContactId,omitempty"`
	DebtorID            *types.GUID             `json:"debtorId,omitempty"`
	DebtorContactID     *types.GUID             `json:"debtorContactId,omitempty"`
	ShipToID            *types.GUID             `json:"shipToId,omitempty"`
	ShipToContactID     *types.GUID             `json:"shipToContactId,omitempty"`
	LanguageISO         string                  `json:"languageIso,omitempty"`
	CurrencyISO         string                  `json:"currencyIso,omitempty"`
	Synced              bool                    `json:"synced,omitempty"`
	QuotationNumber     int64                   `json:"quotationNumber,omitempty"`
	VersionNumber       int32                   `json:"versionNumber,omitempty"`
	Status              *string                 `json:"status,omitempty"`
	StatusID            *types.GUID             `json:"statusId,omitempty"`
	Subject             *string                 `json:"subject,omitempty"`
	TotalPrice          float64                 `json:"totalPrice,omitempty"`
	IsVerified          *bool                   `json:"isVerified,omitempty"`
	CustomerReference   *string                 `json:"customerReference,omitempty"`
	QuotationReference  *string                 `json:"quotationReference,omitempty"`
	Deliverydate        *e_types.DateTimeString `json:"deliverydate,omitempty"`
	Remarks             *string                 `json:"remarks,omitempty"`
	ExpiresDate         *e_types.DateTimeString `json:"expiresDate,omitempty"`
	QuotationTemplateID *types.GUID             `json:"quotationTemplateId,omitempty"`
	ID                  types.GUID              `json:"id,omitempty"`
	CreatedDate         e_types.DateTimeString  `json:"createdDate,omitempty"`
	UpdatedDate         e_types.DateTimeString  `json:"updatedDate,omitempty"`
	OrganizationID      *types.GUID             `json:"organizationId,omitempty"`
	Reference           *string                 `json:"reference,omitempty"`
	CreatorID           types.GUID              `json:"creatorId,omitempty"`
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
			URL:           service.url(urlPath),
			ResponseModel: &quotationsResponse,
		}
		_, _, e := service.get(&requestConfig)
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

func (service *Service) UpdateQuotation(quotationID types.GUID, quotationUpdate *Quotation) *errortools.Error {
	urlPath := fmt.Sprintf("quotations(%s)", quotationID.String())

	requestConfig := go_http.RequestConfig{
		URL:       service.url(urlPath),
		BodyModel: quotationUpdate,
	}
	_, _, e := service.patch(&requestConfig)

	return e
}
