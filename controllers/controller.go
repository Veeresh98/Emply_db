package controllers

import (
	"Emply_DB/models"
	"Emply_DB/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetEmployees(c *gin.Context) {

	employees, err := service.GetAllEmployees()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while getting the employees"})
		return
	}

	c.JSON(http.StatusOK, &employees)

}

func CreateEmployee(c *gin.Context) {

	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while binding the json"})
		return
	}

	createdEmployee, err := service.CreateEmployee(employee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while creating the employee"})
		return
	}

	c.JSON(http.StatusOK, createdEmployee)

}

func GetEmployeeID(c *gin.Context) {

	id := c.Param("id")
	employee, err := service.GetEmployeeByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while getting the id"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {

	id := c.Param("id")

	var updatedEmployee models.Employee

	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad"})
		log.Println("error: error while binding the json")
		return
	}
	employee, err := service.UpdateEmployee(id, updatedEmployee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "status not found"})
		return
	}

	c.JSON(http.StatusOK, employee)

}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	if err := service.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error deleting the id"})
		return
	}
	c.Status(http.StatusNoContent)
}
