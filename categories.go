package elfsquad

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
)

type CategoriesResponse struct {
	Context  string     `json:"@odata.context"`
	Value    []Category `json:"value"`
	NextLink string     `json:"@odata.nextLink"`
}

type Category struct {
	Name           string     `json:"name"`
	ParentID       types.GUID `json:"parentId"`
	Order          int32      `json:"order"`
	ID             types.GUID `json:"id"`
	CreatedDate    string     `json:"createdDate"`
	UpdatedDate    string     `json:"updatedDate"`
	OrganizationID types.GUID `json:"organizationId"`
	CreatorID      types.GUID `json:"creatorId"`
}

func (es *Elfsquad) GetCategories() (*[]Category, *errortools.Error) {
	top := 100
	skip := 0

	categories := []Category{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("categories?$top=%v&$skip=%v", top, skip)

		categoriesReponse := CategoriesResponse{}
		_, _, e := es.get(urlPath, &categoriesReponse)
		if e != nil {
			return nil, e
		}

		rowCount = len(categoriesReponse.Value)

		if rowCount > 0 {
			categories = append(categories, categoriesReponse.Value...)
		}

		skip += top
	}

	return &categories, nil
}
