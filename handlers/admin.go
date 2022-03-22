package handlers

import (
	"attendance-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRegister_Handler(c *gin.Context) {
	// name string, lat float32, long float32, entry_time time.Time, exit_time time.Time
	username := c.PostForm("username")
	pass := c.PostForm("pass")
	admin, err := controllers.RegisterAdmin(username, pass)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}

func AdminLogin_Handler(c *gin.Context) {
	username := c.PostForm("username")
	pass := c.PostForm("pass")
	if err := controllers.LoginAdmin(username, pass); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "Logged IN")

}
