package controllers

import (
	"attendance-backend/models"
	"fmt"
)

func RegisterAdmin(username string, pass string) (*models.Admin_Details, error) {
	admin, err := models.AddAdminDetails(&models.Admin_Details{
		UserName: username,
		Password: pass,
	})
	return admin, err
}
func LoginAdmin(username string, pass string) error {
	res := models.GetAdminDetailsByNamePass(username, pass)
	if res == 0 {
		return fmt.Errorf("company not registered")
	}
	return nil
}
