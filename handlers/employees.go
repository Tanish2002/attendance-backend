package handlers

import (
	"attendance-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

type Employee_Body struct {
	Company_ID uint `json:"company_id"`
}

func EmployeesListHandler(c *fiber.Ctx) error {
	body := new(Register_Body)
	err := c.BodyParser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	company_id := body.Company_ID
	employees := controllers.EmployeeList(company_id)
	return c.JSON(employees)
}
