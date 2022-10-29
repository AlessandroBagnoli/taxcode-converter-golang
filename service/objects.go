package service

type CalculateTaxCodeRequest struct {
	Gender      Gender    `json:"gender" enums:"UNKNOWN,MALE,FEMALE"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth CivilTime `json:"dateOfBirth" format:"date"`
	BirthPlace  string    `json:"birthPlace"`
	Province    string    `json:"province"`
}

type CalculateTaxCodeResponse struct {
	TaxCode string `json:"taxCode"`
}

type CalculatePersonDataRequest struct {
	TaxCode string `json:"taxCode"`
}

type CalculatePersonDataResponse struct {
	Gender      Gender    `json:"gender" enums:"UNKNOWN,MALE,FEMALE"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth CivilTime `json:"dateOfBirth" format:"date"`
	BirthPlace  string    `json:"birthPlace"`
	Province    string    `json:"province"`
	TaxCode     string    `json:"taxCode"`
}

type Gender string

const (
	GenderUnknown Gender = "UNKNOWN"
	GenderMale           = "MALE"
	GenderFemale         = "FEMALE"
)
