package service

import "time"

type CalculatePersonDataRequest struct {
	TaxCode string `json:"taxCode"`
}

type Gender string

const (
	GenderUnknown Gender = "UNKNOWN"
	GenderMale           = "M"
	GenderFemale         = "F"
)

type CalculatePersonDataResponse struct {
	Gender      Gender    `json:"gender"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	BirthPlace  string    `json:"birthPlace"`
	Province    string    `json:"province"`
	TaxCode     string    `json:"taxCode"`
}
