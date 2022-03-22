package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Attendance_Taken struct {
	ID          uint `gorm:"primaryKey"`
	Employee_ID uint
	Name        string
	Entry_Time  time.Time
	Exit_Time   time.Time
	Attendance  int
	Employee    Employee_Details `gorm:"foreignKey:Employee_ID"`
}

func AddEntryDetails(attendance *Attendance_Taken) (*Attendance_Taken, error) {
	if err := DB.Create(attendance).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}
func UpdateExitDetails(attendance *Attendance_Taken, exit_time time.Time) (*Attendance_Taken, error) {
	if err := DB.Model(attendance).UpdateColumns(Attendance_Taken{
		Exit_Time: exit_time,
	}).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}
func CheckAttendanceForDay(id uint) Attendance_Taken {
	var attendance Attendance_Taken
	DB.Where(&Attendance_Taken{
		Employee_ID: id,
	}).Last(&attendance)
	return attendance
}
func CountWorkingDays(eid uint) int64 {
	return DB.Where(&Attendance_Taken{
		Employee_ID: eid,
	}).Find(&Attendance_Taken{}).RowsAffected
}
