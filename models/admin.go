package models

type Admin_Details struct {
	ID         uint `gorm:"primaryKey"`
	UserName   string
	Password   []byte
	Company_ID uint
	Company    Company_Details `gorm:"foreignKey:Company_ID"`
}

func AddAdminDetails(admin_details *Admin_Details) (*Admin_Details, error) {
	if err := DB.Create(admin_details).Error; err != nil {
		return nil, err
	}
	return admin_details, nil
}

func GetAdminDetailsByName(admin_details *Admin_Details) (*Admin_Details, int64) {
	res := DB.Where(admin_details).Take(admin_details)
	return admin_details, res.RowsAffected
}
