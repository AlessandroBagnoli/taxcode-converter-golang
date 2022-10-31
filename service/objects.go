package service

import (
	"cloud.google.com/go/civil"
)

// CalculateTaxCodeRequest is the request for calculating tax code starting from necessary data of a person.
type CalculateTaxCodeRequest struct {
	Gender      Gender     `json:"gender" validate:"required,notblank,oneof=MALE FEMALE" enums:"MALE,FEMALE"`
	Name        string     `json:"name" validate:"required,notblank"`
	Surname     string     `json:"surname" validate:"required,notblank"`
	DateOfBirth civil.Date `json:"dateOfBirth" validate:"required,inthepast" format:"date" swaggertype:"string"`
	BirthPlace  string     `json:"birthPlace" validate:"required,notblank"`
	Province    string     `json:"province" validate:"required,notblank"`
}

// CalculateTaxCodeResponse contains the calculated tax code of a person.
type CalculateTaxCodeResponse struct {
	TaxCode string `json:"taxCode"`
}

// CalculatePersonDataRequest is the request for calculating the data of a person starting from its taxcode.
type CalculatePersonDataRequest struct {
	TaxCode string `json:"taxCode" validate:"required,notblank,taxcode"`
}

// CalculatePersonDataResponse contains the calculated data of a person.
type CalculatePersonDataResponse struct {
	Gender      Gender     `json:"gender" enums:"MALE,FEMALE"`
	Name        string     `json:"name"`
	Surname     string     `json:"surname"`
	DateOfBirth civil.Date `json:"dateOfBirth" format:"date" swaggertype:"string"`
	BirthPlace  string     `json:"birthPlace"`
	Province    string     `json:"province"`
	TaxCode     string     `json:"taxCode"`
}

// Gender is an enum for the gender of a person.
type Gender string

const (
	// GenderUnknown unknown gender
	GenderUnknown Gender = "UNKNOWN"
	// GenderMale male gender
	GenderMale = "MALE"
	// GenderFemale female gender
	GenderFemale = "FEMALE"
)

type GenericRequest interface {
	CalculateTaxCodeRequest | CalculatePersonDataRequest
}
