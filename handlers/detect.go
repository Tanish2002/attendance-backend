package handlers

import (
	"attendance-backend/controllers"
	"attendance-backend/services"
	"bytes"
	"encoding/base64"
	"image/jpeg"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

type Attendance_Body struct {
	Lat        float64 `json:"lat"`
	Long       float64 `json:"long"`
	Company_ID uint    `json:"company_id"`
	Image      string  `json:"image"`
}

func EntryHandler(c *fiber.Ctx) error {
	body := new(Attendance_Body)
	err := c.BodyParser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	rawImage := body.Image
	lat := body.Lat
	long := body.Long
	company_id := body.Company_ID
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	jpgI, err := jpeg.Decode(bytes.NewReader(unbased))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = services.Rec.SaveImage("/tmp/image.jpg", jpgI)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
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
	body := new(Attendance_Body)
	err := c.BodyParser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	rawImage := body.Image
	lat := body.Lat
	long := body.Long
	company_id := body.Company_ID
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	jpgI, err := jpeg.Decode(bytes.NewReader(unbased))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = services.Rec.SaveImage("/tmp/image.jpg", jpgI)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	attendance, err := controllers.ExitDetect("/tmp/image.jpg", cast.ToFloat64(lat), cast.ToFloat64(long), cast.ToUint(company_id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"status":  "exit marked",
		"message": attendance,
	})
}
