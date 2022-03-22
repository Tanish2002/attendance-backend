package controllers

import (
	"attendance-backend/models"
	"attendance-backend/services"
	"fmt"
	"strconv"
)

func RegisterFace(name string, gender string, company_id uint, imagePath string) error {
	faceId, _ := services.DetectFace(imagePath)
	if faceId != "" {
		return fmt.Errorf("face already registered")
	}
	id, err := models.AddEmployeeDetails(name, gender, company_id)
	if err != nil {
		return err
	}
	if err := services.RegisterFace(imagePath, strconv.FormatUint(uint64(id.ID), 10)); err != nil {
		return err
	}
	return nil
}
