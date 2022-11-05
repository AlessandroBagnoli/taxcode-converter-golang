package taxcode

import "taxcode-converter/service"

func calculate(req service.CalculateTaxCodeRequest, cityExtractor func(place service.Place) *service.CityCSV) *service.CalculateTaxCodeResponse {
	//TODO implement BL
	return &service.CalculateTaxCodeResponse{TaxCode: "BGNLSN93P19H294L"}
}
