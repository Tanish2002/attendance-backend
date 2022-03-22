package models

type Admin_Details struct {
	ID       uint `gorm:"primaryKey"`
	UserName string
	Password string
}

func AddAdminDetails(admin_details *Admin_Details) (*Admin_Details, error) {
	if err := DB.Create(admin_details).Error; err != nil {
		return nil, err
	}
	return admin_details, nil
}
func GetAdminDetailsByNamePass(username string, pass string) int64 {
	res := DB.Where(&Admin_Details{
		UserName: username,
		Password: pass,
	}).Take(&Admin_Details{})
	return res.RowsAffected
}
