package handlers

import (
	"attendance-backend/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

func AdminRegister_Handler(c *fiber.Ctx) error {
	// name string, lat float32, long float32, entry_time time.Time, exit_time time.Time
	username := c.FormValue("username")
	pass := c.FormValue("pass")
	company_id := c.FormValue("company_id")
	admin, err := controllers.RegisterAdmin(username, pass, cast.ToUint(company_id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(admin)
}

func AdminLogin_Handler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	pass := c.FormValue("pass")
	company_id := c.FormValue("company_id")
	admin, err := controllers.LoginAdmin(username, pass, cast.ToUint(company_id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(admin)
}
