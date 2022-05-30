package elfsquad

import (
	"fmt"
	"net/http"

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
	ParentID       types.Guid             `json:"parentId"`
	Order          int32                  `json:"order"`
	ID             types.Guid             `json:"id"`
	CreatedDate    e_types.DateTimeString `json:"createdDate"`
	UpdatedDate    e_types.DateTimeString `json:"updatedDate"`
	OrganizationID types.Guid             `json:"organizationId"`
	CreatorID      types.Guid             `json:"creatorId"`
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
			Method:        http.MethodGet,
			Url:           service.urlData(urlPath),
			ResponseModel: &categoriesResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
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
