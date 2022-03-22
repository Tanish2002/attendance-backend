package controllers

import "attendance-backend/models"

type Employee_List struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Gender       string
	Working_Days int64
}

func EmployeeList(company_id uint) []Employee_List {
	employees := models.GetEmployeeDetails(company_id)
	employees_list := []Employee_List{}
	for _, v := range employees {
		employee := Employee_List{
			ID:           v.ID,
			Name:         v.Name,
			Gender:       v.Gender,
			Working_Days: models.CountWorkingDays(v.ID),
		}
		employees_list = append(employees_list, employee)
	}
	return employees_list
}
