package handlers

import (
	"attendance-backend/controllers"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"gorm.io/datatypes"
)

type Company_Body struct {
	Name       string  `json:"name"`
	Password   string  `json:"pass"`
	Lat        float64 `json:"lat"`
	Long       float64 `json:"long"`
	Entry_Time time.Time
	Exit_Time  time.Time
}

func CompanyRegister_Handler(c *fiber.Ctx) error {
	body := new(Company_Body)
	err := c.BodyParser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	name := body.Name
	pass := body.Password
	lat := body.Lat
	long := body.Long
	entry_time_T := body.Entry_Time
	exit_time_T := body.Exit_Time
	entry_time := datatypes.NewTime(entry_time_T.Hour(), entry_time_T.Minute(), entry_time_T.Second(), entry_time_T.Nanosecond())
	exit_time := datatypes.NewTime(exit_time_T.Hour(), exit_time_T.Minute(), exit_time_T.Second(), exit_time_T.Nanosecond())
	company, err := controllers.RegisterComapny(name, pass, cast.ToFloat64(lat), cast.ToFloat64(long), entry_time, exit_time)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(company)

}

func CompanyLogin_Handler(c *fiber.Ctx) error {
	body := new(Company_Body)
	err := c.BodyParser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	name := body.Name
	pass := body.Password
	fmt.Println("name ", name)
	fmt.Println("pass", pass)
	company, err := controllers.LoginCompany(name, pass)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(company)

}
