package handlers

import (
	"attendance-backend/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func RegisterFaceHandler(c *gin.Context) {
	// req, _ := httputil.DumpRequest(c.Request, true)
	// fmt.Println(string(req))
	name := c.PostForm("name")
	gender := c.PostForm("gender")
	company_id_query := c.PostForm("company_id")
	company_id := cast.ToUint(company_id_query)
	fmt.Println(name, gender, company_id)
	image, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.SaveUploadedFile(image, "/tmp/image.jpg")
	if err := controllers.RegisterFace(name, gender, company_id, "/tmp/image.jpg"); err != nil {
		c.String(http.StatusServiceUnavailable, err.Error())
		return
	}
	c.String(http.StatusOK, "Face Registered!")
}
