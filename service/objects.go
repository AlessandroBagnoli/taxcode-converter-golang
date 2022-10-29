package service

type CalculateTaxCodeRequest struct {
	Gender      Gender    `json:"gender" enums:"UNKNOWN,MALE,FEMALE" validate:"required,oneof=UNKNOWN MALE FEMALE"`
	Name        string    `json:"name" validate:"required"`
	Surname     string    `json:"surname" validate:"required"`
	DateOfBirth CivilTime `json:"dateOfBirth" format:"date" validate:"required"`
	BirthPlace  string    `json:"birthPlace" validate:"required"`
	Province    string    `json:"province" validate:"required"`
}

type CalculateTaxCodeResponse struct {
	TaxCode string `json:"taxCode"`
}

type CalculatePersonDataRequest struct {
	TaxCode string `json:"taxCode" validate:"required"`
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
