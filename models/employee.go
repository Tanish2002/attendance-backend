package models

type Employee_Details struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Gender     string
	Company_ID uint
	Company    Company_Details `gorm:"foreignKey:Company_ID"`
}

func AddEmployeeDetails(name string, gender string, company_id uint) (*Employee_Details, error) {
	employee_details := &Employee_Details{
		Name:       name,
		Gender:     gender,
		Company_ID: company_id,
	}
	if err := DB.Create(employee_details).Error; err != nil {
		return nil, err
	}
	return employee_details, nil
}

func GetEmployeeDetails(company_id uint) []Employee_Details {
	var employee_details []Employee_Details
	DB.Where(&Employee_Details{
		Company_ID: company_id,
	}).Find(&employee_details)
	return employee_details
}
func GetEmployeeDetailByID(id string) Employee_Details {
	var employee Employee_Details
	DB.Take(&employee, id)
	return employee
}
