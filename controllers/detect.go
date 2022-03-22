package controllers

import (
	"attendance-backend/models"
	"attendance-backend/services"
	"fmt"
	"time"
)

func EntryDetect(imagePath string) (*models.Attendance_Taken, error) {
	id, err := services.DetectFace(imagePath)
	if err != nil {
		return nil, err
	}
	current_time := time.Now().Local()
	employee := models.GetEmployeeDetailByID(id)
	attendance_check := models.CheckAttendanceForDay(employee.ID)
	if attendance_check.Entry_Time.Truncate(24 * time.Hour).Equal(current_time.Truncate(24 * time.Hour)) {
		return nil, fmt.Errorf("entry already marked for today")
	}
	attendance, err := models.AddEntryDetails(&models.Attendance_Taken{
		Employee_ID: employee.ID,
		Name:        employee.Name,
		Entry_Time:  current_time,
	})
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func ExitDetect(imagePath string) (*models.Attendance_Taken, error) {
	id, err := services.DetectFace(imagePath)
	if err != nil {
		return nil, err
	}
	current_time := time.Now().Local()
	employee := models.GetEmployeeDetailByID(id)
	attendance_check := models.CheckAttendanceForDay(employee.ID)
	// Check if entry is marked for today
	if attendance_check.Entry_Time.IsZero() {
		return nil, fmt.Errorf("entry not marked for today")
	}
	// Check if exit is already marked
	if attendance_check.Exit_Time.Truncate(24 * time.Hour).Equal(current_time.Truncate(24 * time.Hour)) {
		return nil, fmt.Errorf("exit already marked for today")
	}
	attendance, err := models.UpdateExitDetails(&attendance_check, current_time)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}
