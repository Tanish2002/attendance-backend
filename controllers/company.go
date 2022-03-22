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
func LoginCompany(name string, pass string) error {
	res := models.GetCompanyDetailsByNamePass(name, pass)
	if res == 0 {
		return fmt.Errorf("company not registered")
	}
	return nil
}
