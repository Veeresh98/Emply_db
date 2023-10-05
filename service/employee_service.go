package service

import (
	"Emply_DB/models"
	"Emply_DB/repositories"
)

func GetAllEmployees() ([]models.Employee, error) {
	return repositories.GetAllEmployees()
}

func CreateEmployee(employee models.Employee) (models.Employee, error) {
	return repositories.CreateEmployee(employee)
}

func GetEmployeeByID(id string) (models.Employee, error) {
	return repositories.GetEmployeeByID(id)
}

func UpdateEmployee(id string, updateEmployee models.Employee) (models.Employee, error) {
	return repositories.UpdateEmployee(id, updateEmployee)
}

func DeleteEmployee(id string) error {
	return repositories.DeleteEmployee(id)
}
