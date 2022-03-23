package main

import (
	"attendance-backend/configuration"
	"attendance-backend/handlers"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//var Router *gin.Engine

func init() {
	configuration.Init()
}

func main() {
	log.Printf("Fiber start")
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	// Router.Use(cors.Default())

	// Attendance Endpoints
	app.Post("/entry", handlers.EntryHandler)
	app.Post("/exit", handlers.ExitHandler)
	app.Post("/registerface", handlers.RegisterFaceHandler)

	// Employee List
	app.Post("/employees", handlers.EmployeesListHandler)

	// Company Endpoints
	app.Post("/companyregister", handlers.CompanyRegister_Handler)
	app.Post("/companylogin", handlers.CompanyLogin_Handler)

	// Admin Endpoints
	app.Post("/adminregister", handlers.AdminRegister_Handler)
	app.Post("/adminlogin", handlers.AdminLogin_Handler)

	if configuration.Runmode == "dev" {
		app.Listen(":8080")
	}
	app.Listen(":8080")
}
