package controllers

import (
	"attendance-backend/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

func RegisterComapny(name string, pass string, lat float64, long float64, entry_time datatypes.Time, exit_time datatypes.Time) (*models.Company_Details, error) {
	company_check := &models.Company_Details{
		Name: name,
	}
	_, res := models.GetCompanyDetailsByName(company_check)
	if res != 0 {
		return nil, fmt.Errorf("company name not available")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error when hashing password: %s", err.Error())
	}
	company, err := models.AddCompanyDetails(&models.Company_Details{
		Name:       name,
		Password:   hash,
		Lat:        lat,
		Long:       long,
		Entry_Time: entry_time,
		Exit_Time:  exit_time,
	})
	return company, err
}
func LoginCompany(name string, pass string) (*models.Company_Details, error) {

	company_details := &models.Company_Details{
		Name: name,
	}
	company, res := models.GetCompanyDetailsByName(company_details)
	if res == 0 {
		return nil, fmt.Errorf("company not registered")
	}
	if err := bcrypt.CompareHashAndPassword(company.Password, []byte(pass)); err != nil {
		return nil, fmt.Errorf("wrong password")
	}
	return company, nil
}
