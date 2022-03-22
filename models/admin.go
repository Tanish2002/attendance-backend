package models

type Admin_Details struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	UserName string
	Password string
}
