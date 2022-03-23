package handlers

import (
	"attendance-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func AdminRegister_Handler(c *gin.Context) {
	// name string, lat float32, long float32, entry_time time.Time, exit_time time.Time
	username := c.PostForm("username")
	pass := c.PostForm("pass")
	company_id := c.PostForm("company_id")
	admin, err := controllers.RegisterAdmin(username, pass, cast.ToUint(company_id))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}

func AdminLogin_Handler(c *gin.Context) {
	username := c.PostForm("username")
	pass := c.PostForm("pass")
	company_id := c.PostForm("company_id")
	admin, err := controllers.LoginAdmin(username, pass, cast.ToUint(company_id))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}
