package main

import (
	"attendance-backend/configuration"
	"attendance-backend/handlers"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	configuration.Init()
	log.Printf("Gin cold start")
	handlers.Router = gin.Default()
	handlers.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowCredentials: true,
	}))

	// Attendance Endpoints
	handlers.Router.POST("/entry", handlers.EntryHandler)
	handlers.Router.POST("/exit", handlers.ExitHandler)
	handlers.Router.POST("/registerface", handlers.RegisterFaceHandler)

	// Employee List
	handlers.Router.POST("/employees", handlers.EmployeesListHandler)

	// Company Endpoints
	handlers.Router.POST("/companyregister", handlers.CompanyRegister_Handler)
	handlers.Router.POST("/companylogin", handlers.CompanyLogin_Handler)

}

func main() {
	// services.AddFile(&services.Rec, "fotos/1.jpg", "lol")
	// services.Rec.SaveDataset("dataset.json")

	if configuration.Runmode == "dev" {
		handlers.Router.Run(":8080")
	}
	handlers.Router.Run()
	//lambda.Start(handlers.LambdaHandler)
	//algnhsa.ListenAndServe(router, nil)
}
