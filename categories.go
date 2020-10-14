package elfsquad

import (
	"fmt"

	types "github.com/Leapforce-nl/go_types"
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

func (es *Elfsquad) GetCategories() (*[]Category, error) {
	top := 100
	skip := 0

	categories := []Category{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		url := fmt.Sprintf("%s/categories?$top=%v&$skip=%v", apiURLData, top, skip)

		categoriesReponse := CategoriesResponse{}

		_, err := es.oAuth2.Get(url, &categoriesReponse)
		if err != nil {
			return nil, err
		}

		rowCount = len(categoriesReponse.Value)

		if rowCount > 0 {
			categories = append(categories, categoriesReponse.Value...)
		}

		skip += top
	}

	return &categories, nil
}
