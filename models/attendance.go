package models

import (
	"math"
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

func Measure(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 { // generally used geo measurement function
	var R = 6378.137 // Radius of earth in KM
	var dLat = lat2*math.Pi/180 - lat1*math.Pi/180
	var dLon = lon2*math.Pi/180 - lon1*math.Pi/180
	var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = R * c
	return d * 1000 // meters
}
