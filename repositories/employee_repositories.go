package repositories

import (
	db "Emply_DB/database"
	"Emply_DB/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	EmployeeCollection *mongo.Collection
)

func init() {
	EmployeeCollection = db.MongoClient.Database("employeeDatabase").Collection("employee")
}

func GetAllEmployees() ([]models.Employee, error) {

	var employee []models.Employee

	cursor, err := EmployeeCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Println("error: Error for getting the EmployeeCollection")
		return employee, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
		}
	}(cursor, context.Background())

	err = cursor.All(context.TODO(), &employee)
	if err != nil {
		log.Println("error: Error")
	}

	return employee, nil
}

func CreateEmployee(employee models.Employee) (models.Employee, error) {

	result, err := EmployeeCollection.InsertOne(context.TODO(), employee)

	if err != nil {
		log.Println("error: error while inserting the employee to the database")
		return models.Employee{}, err
	}
	employee.ID = result.InsertedID.(primitive.ObjectID)

	return employee, nil
}

func GetEmployeeByID(id string) (models.Employee, error) {

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("error: error while getting the employee id")
		return models.Employee{}, err
	}
	var employee models.Employee

	err = EmployeeCollection.FindOne(context.TODO(), bson.M{"_id": employeeID}).Decode(&employee)
	if err != nil {
		log.Println("error: error while decoding")
		return models.Employee{}, err
	}

	return employee, nil

}

func UpdateEmployee(id string, updateEmployee models.Employee) (models.Employee, error) {

	employeeID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("error: error while getting the employee id")
		return models.Employee{}, err
	}

	update := bson.M{"$set": updateEmployee}

	_, err = EmployeeCollection.UpdateOne(context.TODO(), bson.M{"_id": employeeID}, update)

	if err != nil {
		log.Println(err)
		return models.Employee{}, err
	}
	updateEmployee.ID = employeeID

	return updateEmployee, nil
}

func DeleteEmployee(id string) error {
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("error: error while getting the id")
		return err
	}
	_, err = EmployeeCollection.DeleteOne(context.TODO(), bson.M{"_id": employeeID})
	if err != nil {
		log.Println("error: error while deleting the id")
		return err
	}

	return nil
}
