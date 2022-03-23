package handlers

import (
	"attendance-backend/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

func EntryHandler(c *fiber.Ctx) error {
	lat := c.FormValue("lat")
	long := c.FormValue("lon")
	image, err := c.FormFile("image")
	company_id := c.FormValue("company_id")
	if err != nil {
		fmt.Println(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	c.SaveFile(image, "/tmp/image.jpg")
	attendance, err := controllers.EntryDetect("/tmp/image.jpg", cast.ToFloat64(lat), cast.ToFloat64(long), cast.ToUint(company_id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"status":  "entry marked",
		"message": attendance,
	})
}
func ExitHandler(c *fiber.Ctx) error {
	lat := c.FormValue("lat")
	long := c.FormValue("lon")
	image, err := c.FormFile("image")
	company_id := c.FormValue("company_id")
	if err != nil {
		fmt.Println(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	c.SaveFile(image, "/tmp/image.jpg")
	attendance, err := controllers.ExitDetect("/tmp/image.jpg", cast.ToFloat64(lat), cast.ToFloat64(long), cast.ToUint(company_id))
	fmt.Println("THE ATTENDANCE OBJECT IS", attendance)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"status":  "exit marked",
		"message": attendance,
	})
}
