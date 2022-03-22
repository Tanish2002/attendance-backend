package handlers

import (
	"attendance-backend/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EntryHandler(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.SaveUploadedFile(image, "/tmp/image.jpg")
	attendance, err := controllers.EntryDetect("/tmp/image.jpg")
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
	image, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.SaveUploadedFile(image, "/tmp/image.jpg")
	attendance, err := controllers.ExitDetect("/tmp/image.jpg")
	fmt.Println("THE ATTENDANCE OBJECT IS", attendance)
	if err != nil {
		c.String(http.StatusServiceUnavailable, err.Error())
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
