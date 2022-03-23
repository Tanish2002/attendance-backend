package controllers

import (
	"attendance-backend/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func RegisterAdmin(username string, pass string, company_id uint) (*models.Admin_Details, error) {
	admin_check := &models.Admin_Details{
		UserName:   username,
		Company_ID: company_id,
		Company: models.Company_Details{
			ID: company_id,
		},
	}
	_, res := models.GetAdminDetailsByName(admin_check)
	if res != 0 {
		return nil, fmt.Errorf("admin name not available")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error when hashing password: %s", err.Error())
	}
	admin, err := models.AddAdminDetails(&models.Admin_Details{
		UserName: username,
		Password: hash,
		Company: models.Company_Details{
			ID: company_id,
		},
	})
	return admin, err
}

// func LoginAdmin(username string, pass string) (*models.Admin_Details, error) {
// 	admin_details := &models.Admin_Details{
// 		UserName: username,
// 		Password: pass,
// 	}
// 	admin, res := models.GetAdminDetailsByNamePass(admin_details)
// 	if res == 0 {
// 		return nil, fmt.Errorf("company not registered")
// 	}
// 	return admin, nil
// }
func LoginAdmin(username string, pass string, company_id uint) (*models.Admin_Details, error) {
	admin_details := &models.Admin_Details{
		UserName:   username,
		Company_ID: company_id,
		Company: models.Company_Details{
			ID: company_id,
		},
	}
	admin, res := models.GetAdminDetailsByName(admin_details)
	if res == 0 {
		return nil, fmt.Errorf("admin not registered")
	}
	if err := bcrypt.CompareHashAndPassword(admin.Password, []byte(pass)); err != nil {
		return nil, fmt.Errorf("wrong password")
	}
	return admin, nil
}
