package handlers

import (
	"attendance-backend/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

func RegisterFaceHandler(c *fiber.Ctx) error {
	// req, _ := httputil.DumpRequest(c.Request, true)
	// fmt.Println(string(req))
	name := c.FormValue("name")
	gender := c.FormValue("gender")
	company_id_query := c.FormValue("company_id")
	company_id := cast.ToUint(company_id_query)

	image, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// image := c.PostForm("image")
	// coI := strings.Index(string(image), ",")
	// rawImage := string(image)[coI+1:]
	// unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	// jpgI, err := jpeg.Decode(bytes.NewReader(unbased))
	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// 	return
	// }
	//err = services.Rec.SaveImage("/tmp/image.jpg", jpgI)

	c.SaveFile(image, "/tmp/image.jpg")

	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAaa")
	if err := controllers.RegisterFace(name, gender, company_id, "/tmp/image.jpg"); err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	return c.SendString("Face Registered!")
}
