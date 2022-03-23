package main

import (
	"attendance-backend/configuration"
	"attendance-backend/handlers"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	configuration.Init()
	log.Printf("Gin cold start")
	Router = gin.Default()
	Router.Use(cors.Default())

	// Attendance Endpoints
	Router.POST("/entry", handlers.EntryHandler)
	Router.POST("/exit", handlers.ExitHandler)
	Router.POST("/registerface", handlers.RegisterFaceHandler)

	// Employee List
	Router.POST("/employees", handlers.EmployeesListHandler)

	// Company Endpoints
	Router.POST("/companyregister", handlers.CompanyRegister_Handler)
	Router.POST("/companylogin", handlers.CompanyLogin_Handler)

	// Admin Endpoints
	Router.POST("/adminregister", handlers.AdminRegister_Handler)
	Router.POST("/adminlogin", handlers.AdminLogin_Handler)

}

func main() {
	// services.AddFile(&services.Rec, "fotos/1.jpg", "lol")
	// services.Rec.SaveDataset("dataset.json")

	if configuration.Runmode == "dev" {
		Router.Run(":8080")
	}
	Router.Run()
	//lambda.Start(handlers.LambdaHandler)
	//algnhsa.ListenAndServe(router, nil)
}
