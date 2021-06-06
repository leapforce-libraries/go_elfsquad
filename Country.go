package elfsquad

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type CountriesResponse struct {
	Context  string    `json:"@odata.context"`
	Value    []Country `json:"value"`
	NextLink string    `json:"@odata.nextLink"`
}

type Country struct {
	ISO         string `json:"iso"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	EnglishName string `json:"englishName"`
	PhonePrefix string `json:"phonePrefix"`
	Capital     string `json:"capital"`
}

func (service *Service) GetCountries() (*[]Country, *errortools.Error) {
	top := 100
	skip := 0

	countries := []Country{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("countries?$top=%v&$skip=%v", top, skip)

		countriesResponse := CountriesResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           service.urlData(urlPath),
			ResponseModel: &countriesResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(countriesResponse.Value)

		if rowCount > 0 {
			countries = append(countries, countriesResponse.Value...)
		}

		skip += top
	}

	return &countries, nil
}
