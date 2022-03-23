package handlers

import (
	"attendance-backend/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"gorm.io/datatypes"
)

func CompanyRegister_Handler(c *fiber.Ctx) error {
	// name string, lat float32, long float32, entry_time time.Time, exit_time time.Time
	name := c.FormValue("name")
	pass := c.FormValue("pass")
	lat := c.FormValue("lat")
	long := c.FormValue("long")
	entry_time_query := c.FormValue("entry_time")
	entry_time_T := cast.ToTime(entry_time_query)
	exit_time_query := c.FormValue("exit_time")
	exit_time_T := cast.ToTime(exit_time_query)
	entry_time := datatypes.NewTime(entry_time_T.Hour(), entry_time_T.Minute(), entry_time_T.Second(), entry_time_T.Nanosecond())
	exit_time := datatypes.NewTime(exit_time_T.Hour(), exit_time_T.Minute(), exit_time_T.Second(), exit_time_T.Nanosecond())
	company, err := controllers.RegisterComapny(name, pass, cast.ToFloat64(lat), cast.ToFloat64(long), entry_time, exit_time)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(company)

}

func CompanyLogin_Handler(c *fiber.Ctx) error {
	name := c.FormValue("name")
	pass := c.FormValue("pass")
	company, err := controllers.LoginCompany(name, pass)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(company)

}
