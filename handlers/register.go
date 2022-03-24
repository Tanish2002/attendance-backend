package handlers

import (
	"attendance-backend/controllers"
	"attendance-backend/services"
	"bytes"
	"encoding/base64"
	"image/jpeg"

	"github.com/gofiber/fiber/v2"
)

type Register_Body struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Company_ID uint   `json:"company_id"`
	Image      string `json:"image"`
}

func RegisterFaceHandler(c *fiber.Ctx) error {
	body := new(Register_Body)
	err := c.BodyParser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	name := body.Name
	gender := body.Gender
	company_id := body.Company_ID
	rawImage := body.Image
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	jpgI, err := jpeg.Decode(bytes.NewReader(unbased))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = services.Rec.SaveImage("/tmp/image.jpg", jpgI)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := controllers.RegisterFace(name, gender, company_id, "/tmp/image.jpg"); err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	return c.SendString("Face Registered!")
}
