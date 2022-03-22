package handlers

import (
	"attendance-backend/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func EntryHandler(c *gin.Context) {
	lat := c.PostForm("lat")
	long := c.PostForm("lon")
	image, err := c.FormFile("image")
	company_id := c.PostForm("company_id")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.SaveUploadedFile(image, "/tmp/image.jpg")
	attendance, err := controllers.EntryDetect("/tmp/image.jpg", cast.ToFloat64(lat), cast.ToFloat64(long), cast.ToUint(company_id))
	if err != nil {
		c.String(http.StatusServiceUnavailable, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "entry marked",
		"message": attendance,
	})
	// if err := controllers.EntryDetect("/tmp/image.jpg"); err != nil {
	// 	c.String(http.StatusServiceUnavailable, err.Error())
	// 	return
	// }
}
func ExitHandler(c *gin.Context) {
	lat := c.PostForm("lat")
	long := c.PostForm("lon")
	image, err := c.FormFile("image")
	company_id := c.PostForm("company_id")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.SaveUploadedFile(image, "/tmp/image.jpg")
	attendance, err := controllers.ExitDetect("/tmp/image.jpg", cast.ToFloat64(lat), cast.ToFloat64(long), cast.ToUint(company_id))
	fmt.Println("THE ATTENDANCE OBJECT IS", attendance)
	if err != nil {
		c.String(http.StatusServiceUnavailable, err.Error())
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "exit marked",
		"message": attendance,
	})
	// if err := controllers.EntryDetect("/tmp/image.jpg"); err != nil {
	// 	c.String(http.StatusServiceUnavailable, err.Error())
	// 	return
	// }
}
