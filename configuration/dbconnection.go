package configuration

import (
	"attendance-backend/models"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	HostName     string
	UserName     string
	Password     string
	DatabaseName string
}

func registerDatabase() {
	MYSQLCONFIG := &MysqlConfig{
		HostName:     viper.GetString(Runmode + ".mysql.host"),
		UserName:     viper.GetString(Runmode + ".mysql.user"),
		Password:     viper.GetString(Runmode + ".mysql.password"),
		DatabaseName: viper.GetString(Runmode + ".mysql.database"),
	}

	mysqlConf := MYSQLCONFIG.UserName + ":" +
		MYSQLCONFIG.Password + "@tcp(" +
		MYSQLCONFIG.HostName + ")/" +
		MYSQLCONFIG.DatabaseName + "?parseTime=true"
	log.Println("conf", mysqlConf)
	var err error
	models.DB, err = gorm.Open(mysql.Open(mysqlConf), &gorm.Config{
		// Silent ORM Logs
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	// Create Tables if they aren't already created
	models.DB.AutoMigrate(&models.Attendance_Taken{}, &models.Admin_Details{}, &models.Employee_Details{}, &models.Company_Details{}, &models.Dataset{})
}
