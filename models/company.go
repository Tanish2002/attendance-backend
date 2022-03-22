package models

import (
	"gorm.io/datatypes"
)

type Company_Details struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Password   string
	Lat        float64
	Long       float64
	Entry_Time datatypes.Time
	Exit_Time  datatypes.Time
}

func AddCompanyDetails(company_details *Company_Details) (*Company_Details, error) {
	if err := DB.Create(company_details).Error; err != nil {
		return nil, err
	}
	return company_details, nil
}

func GetCompanyDetailByID(id string) Company_Details {
	var company Company_Details
	DB.Take(&company, id)
	return company
}

func GetCompanyDetailsByNamePass(name string, pass string) int64 {
	res := DB.Where(&Company_Details{
		Name:     name,
		Password: pass,
	}).Take(&Company_Details{})
	return res.RowsAffected
}
