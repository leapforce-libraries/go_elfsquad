package elfsquad

import (
	"fmt"

	e_types "github.com/leapforce-libraries/go_elfsquad/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
)

type CategoriesResponse struct {
	Context  string     `json:"@odata.context"`
	Value    []Category `json:"value"`
	NextLink string     `json:"@odata.nextLink"`
}

type Category struct {
	Name           string                 `json:"name"`
	ParentID       types.GUID             `json:"parentId"`
	Order          int32                  `json:"order"`
	ID             types.GUID             `json:"id"`
	CreatedDate    e_types.DateTimeString `json:"createdDate"`
	UpdatedDate    e_types.DateTimeString `json:"updatedDate"`
	OrganizationID types.GUID             `json:"organizationId"`
	CreatorID      types.GUID             `json:"creatorId"`
}

func (service *Service) GetCategories() (*[]Category, *errortools.Error) {
	top := 100
	skip := 0

	categories := []Category{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("categories?$top=%v&$skip=%v", top, skip)

		categoriesResponse := CategoriesResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           service.urlData(urlPath),
			ResponseModel: &categoriesResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(categoriesResponse.Value)

		if rowCount > 0 {
			categories = append(categories, categoriesResponse.Value...)
		}

		skip += top
	}

	return &categories, nil
}
