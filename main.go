package main

import (
	"Emply_DB/controllers"
	"Emply_DB/database"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	database.InitializeMongodb()

	router.GET("/employees", controllers.GetEmployees)
	router.GET("/employees/:id", controllers.GetEmployeeID)
	router.POST("/employees", controllers.CreateEmployee)
	router.PUT("/employees/:id", controllers.UpdateEmployee)
	router.DELETE("employees/:id", controllers.DeleteEmployee)

	err := router.Run("8080")
	if err != nil {
		return
	}

}
