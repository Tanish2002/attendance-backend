package handlers

import (
	"attendance-backend/controllers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

func EmployeesListHandler(c *fiber.Ctx) error {
	company_id_query := c.FormValue("company_id")
	if company_id_query == "" {
		return fiber.NewError(http.StatusBadRequest, "company_id parameter is required")
	}
	company_id := cast.ToUint(company_id_query)
	employees := controllers.EmployeeList(company_id)
	return c.JSON(employees)
}
