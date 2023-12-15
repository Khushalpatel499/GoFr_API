package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khushalpatel499/gofr_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://khushal:Khushal4989@cluster0.xirtede.mongodb.net/?retryWrites=true&w=majority"
const dbName = "garage"
const colName = "carList"

var collection *mongo.Collection

type cars struct {
	model.Garage
}

func (c *cars) IsEmpty() bool {
	return c.OwnerName == "" || c.CarNumber == "" || c.ModalName == ""
}

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	checkNilError(err)

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance

	fmt.Println("Collection instance is ready")
}

// MongoDB helper

//insert 1 record of car

func insertOneCar(car cars) {
	inserted, err := collection.InsertOne(context.Background(), car)

	checkNilError(err)

	fmt.Println("Inserted 1 car detail in Database with id:", inserted.InsertedID)
}

//update 1 car record

func updateOneCar(carId string) {
	id, err := primitive.ObjectIDFromHex(carId)

	checkNilError(err)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"repair": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	checkNilError(err)

	fmt.Println("Update detail count:", result.ModifiedCount)
}

// Delete 1 car record

func deleteOneCar(carId string) {
	id, err := primitive.ObjectIDFromHex(carId)

	checkNilError(err)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	checkNilError(err)
	fmt.Println("Car detail is deleted with count:", deleteCount)
}

// Delete all car detail

func deleteAllCar() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})

	checkNilError(err)
	fmt.Println("Number of car detail deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all car detail

func getAllCars() []primitive.D {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	checkNilError(err)

	var cars []primitive.D

	for cur.Next(context.Background()) {
		var car bson.D
		err := cur.Decode(&car)
		checkNilError(err)
		cars = append(cars, car)
	}
	defer cur.Close(context.Background())
	return cars

}

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/cars-detail")
	allCars := getAllCars()
	json.NewEncoder(w).Encode(allCars)
}

func InsertOneCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/cars-detail")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Add the car detail")
	}

	var car cars

	_ = json.NewDecoder(r.Body).Decode(&car)

	//check for empty field in json
	if car.IsEmpty() {
		json.NewEncoder(w).Encode("Some field are empty in json")
		return
	}
	insertOneCar(car)
	json.NewEncoder(w).Encode(car)

}

func UpdateOneCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/cars-detail")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneCar(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteACars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/cars-detail")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneCar(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/cars-detail")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllCar()
	json.NewEncoder(w).Encode(count)
}

//check err

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
