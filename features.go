package elfsquad

import (
	"fmt"
)

type FeaturesResponse struct {
	Value []Feature `json:"value"`
}

type Feature struct {
	Context  string `json:"@odata.context"`
	Name     string `json:"name"`
	NextLink string `json:"@odata.nextLink"`
}

func (es *Elfsquad) GetFeatures() (*[]Feature, error) {
	top := 100
	skip := 0

	features := []Feature{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		url := fmt.Sprintf("%s/features?$top=%v&$skip=%v", apiURLData, top, skip)

		featuresReponse := FeaturesResponse{}

		_, err := es.oAuth2.Get(url, &featuresReponse)
		if err != nil {
			return nil, err
		}

		rowCount = len(featuresReponse.Value)

		if rowCount > 0 {
			features = append(features, featuresReponse.Value...)
		}

		skip += top
	}

	return &features, nil
}
