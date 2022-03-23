package controllers

import (
	"attendance-backend/models"
	"fmt"

	"gorm.io/datatypes"
)

func RegisterComapny(name string, pass string, lat float64, long float64, entry_time datatypes.Time, exit_time datatypes.Time) (*models.Company_Details, error) {
	company, err := models.AddCompanyDetails(&models.Company_Details{
		Name:       name,
		Password:   pass,
		Lat:        lat,
		Long:       long,
		Entry_Time: entry_time,
		Exit_Time:  exit_time,
	})
	return company, err
}
func LoginCompany(name string, pass string) (*models.Company_Details, error) {
	company_details := &models.Company_Details{
		Name:     name,
		Password: pass,
	}
	company, res := models.GetCompanyDetailsByNamePass(company_details)
	fmt.Println(company)
	fmt.Println(res)
	if res == 0 {
		return nil, fmt.Errorf("company not registered")
	}
	return company, nil
}
