package elfsquad

import (
	"fmt"

	types "github.com/Leapforce-nl/go_types"
)

type QuotationsResponse struct {
	Context  string      `json:"@odata.context"`
	Value    []Quotation `json:"value"`
	NextLink string      `json:"@odata.nextLink"`
}

type Quotation struct {
	SellerID            types.GUID `json:"sellerId"`
	SellerContactID     types.GUID `json:"sellerContactId"`
	DebtorID            types.GUID `json:"debtorId"`
	DebtorContactID     types.GUID `json:"debtorContactId"`
	ShipToID            types.GUID `json:"shipToId"`
	ShipToContactID     types.GUID `json:"shipToContactId"`
	Synced              bool       `json:"synced"`
	QuotationNumber     int64      `json:"quotationNumber"`
	VersionNumber       int32      `json:"versionNumber"`
	Status              string     `json:"status"`
	Subject             string     `json:"subject"`
	TotalPrice          float64    `json:"totalPrice"`
	IsVerified          bool       `json:"isVerified"`
	CustomerReference   string     `json:"customerReference"`
	QuotationReference  string     `json:"quotationReference"`
	Deliverydate        string     `json:"deliverydate"`
	Remarks             string     `json:"remarks"`
	ExpiresDate         string     `json:"expiresDate"`
	QuotationTemplateID types.GUID `json:"quotationTemplateId"`
	ID                  types.GUID `json:"id"`
	CreatedDate         string     `json:"createdDate"`
	UpdatedDate         string     `json:"updatedDate"`
	OrganizationID      types.GUID `json:"organizationId"`
	CreatorID           types.GUID `json:"creatorId"`
}

func (es *Elfsquad) GetQuotations() (*[]Quotation, error) {
	top := 100
	skip := 0

	quotations := []Quotation{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		url := fmt.Sprintf("%s/quotations?$top=%v&$skip=%v", apiURLData, top, skip)

		quotationsReponse := QuotationsResponse{}

		_, err := es.oAuth2.Get(url, &quotationsReponse)
		if err != nil {
			return nil, err
		}

		rowCount = len(quotationsReponse.Value)

		if rowCount > 0 {
			quotations = append(quotations, quotationsReponse.Value...)
		}

		skip += top
	}

	return &quotations, nil
}
