package handlers

import (
	"attendance-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func EmployeesListHandler(c *gin.Context) {
	company_id_query := c.PostForm("company_id")
	if company_id_query == "" {
		c.String(http.StatusBadRequest, "company_id parameter is required")
		return
	}
	company_id := cast.ToUint(company_id_query)
	employees := controllers.EmployeeList(company_id)
	c.JSON(http.StatusOK, employees)
}
