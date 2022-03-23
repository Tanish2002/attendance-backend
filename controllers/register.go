package controllers

import (
	"attendance-backend/services"
	"fmt"
)

func RegisterFace(name string, gender string, company_id uint, imagePath string) error {
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAaa")
	faceId, err := services.DetectFace(imagePath)
	fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	fmt.Println(faceId, err)
	if faceId != "" {
		return fmt.Errorf("face already registered")
	}
	if err != nil {
		return fmt.Errorf("error while detecting face: %s", err)
	}
	// id, err := models.AddEmployeeDetails(name, gender, company_id)
	// if err != nil {
	// 	return fmt.Errorf("error while adding details to database: %s", err)
	// }
	// if err := services.RegisterFace(imagePath, strconv.FormatUint(uint64(id.ID), 10)); err != nil {
	// 	return fmt.Errorf("error while registering face %s", err)
	// }
	return nil
}
