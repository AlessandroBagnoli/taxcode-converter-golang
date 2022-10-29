package service

type CalculateTaxCodeRequest struct {
	Gender      Gender    `json:"gender" enums:"UNKNOWN,MALE,FEMALE" validate:"required,notblank,oneof=MALE FEMALE"`
	Name        string    `json:"name" validate:"required,notblank"`
	Surname     string    `json:"surname" validate:"required,notblank"`
	DateOfBirth CivilTime `json:"dateOfBirth" format:"date" validate:"required"`
	BirthPlace  string    `json:"birthPlace" validate:"required,notblank"`
	Province    string    `json:"province" validate:"required,notblank"`
}

type CalculateTaxCodeResponse struct {
	TaxCode string `json:"taxCode"`
}

type CalculatePersonDataRequest struct {
	TaxCode string `json:"taxCode" validate:"required,notblank"`
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
