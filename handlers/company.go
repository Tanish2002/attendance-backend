package handlers

import (
	"attendance-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/datatypes"
)

func CompanyRegister_Handler(c *gin.Context) {
	// name string, lat float32, long float32, entry_time time.Time, exit_time time.Time
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	lat := c.PostForm("lat")
	long := c.PostForm("long")
	entry_time_query := c.PostForm("entry_time")
	entry_time_T := cast.ToTime(entry_time_query)
	exit_time_query := c.PostForm("exit_time")
	exit_time_T := cast.ToTime(exit_time_query)
	entry_time := datatypes.NewTime(entry_time_T.Hour(), entry_time_T.Minute(), entry_time_T.Second(), entry_time_T.Nanosecond())
	exit_time := datatypes.NewTime(exit_time_T.Hour(), exit_time_T.Minute(), exit_time_T.Second(), exit_time_T.Nanosecond())
	company, err := controllers.RegisterComapny(name, pass, cast.ToFloat64(lat), cast.ToFloat64(long), entry_time, exit_time)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, company)

}

func CompanyLogin_Handler(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	if err := controllers.LoginCompany(name, pass); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "Logged IN")

}
