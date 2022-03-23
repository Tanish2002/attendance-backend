package handlers

import (
	"attendance-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func RegisterFaceHandler(c *gin.Context) {
	// image := c.PostForm("image")

	name := c.PostForm("name")
	gender := c.PostForm("gender")
	company_id_query := c.PostForm("company_id")
	company_id := cast.ToUint(company_id_query)
	image, err := c.FormFile("image")
	// r, err := awslambda.NewReaderMultipart(ProxyRequest)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// 	return
	// }
	// part, err := r.NextPart()
	// for part.FormName() != "image" {
	// 	part, err = r.NextPart()
	// }
	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// 	return
	// }
	// content, err := io.ReadAll(part)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// 	return
	// }
	// err = os.WriteFile("/tmp/image.jpg", content, 0666)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// }
	c.SaveUploadedFile(image, "/tmp/image.jpg")
	if err := controllers.RegisterFace(name, gender, company_id, "/tmp/image.jpg"); err != nil {
		c.String(http.StatusServiceUnavailable, err.Error())
		return
	}
	c.String(http.StatusOK, "Face Registered!")
}
