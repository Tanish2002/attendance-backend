package handlers

import (
	"attendance-backend/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Admin_Body struct {
	UserName  string `json:"username"`
	Password  string `json:"pass"`
	CompanyID uint   `json:"company_id"`
}

func AdminRegister_Handler(c *fiber.Ctx) error {
	body := new(Admin_Body)
	err := c.BodyParser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	username := body.UserName
	pass := body.Password
	company_id := body.CompanyID
	admin, err := controllers.RegisterAdmin(username, pass, company_id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(admin)
}

func AdminLogin_Handler(c *fiber.Ctx) error {
	body := new(Admin_Body)
	c.BodyParser(body)
	username := body.UserName
	pass := body.Password
	company_id := body.CompanyID
	fmt.Println(username, pass, company_id)
	admin, err := controllers.LoginAdmin(username, pass, company_id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(admin)
}
