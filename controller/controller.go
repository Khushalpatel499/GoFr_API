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

/*func insertOneCar(car model.Garage) {
	inserted, err := collection.InsertOne(context.Background(), car)

	checkNilError(err)

	fmt.Println("Inserted 1 car detail in Database with id:", inserted.InsertedID)
	//return inserted.InsertedID.(string)
}*/

func insertOneCar(car model.Garage) (primitive.ObjectID, error) {
	inserted, err := collection.InsertOne(context.Background(), car)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Print the inserted ID to the console
	fmt.Println("Inserted 1 car detail in Database with id:", inserted.InsertedID)

	// Return the inserted ID
	return inserted.InsertedID.(primitive.ObjectID), nil
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
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// Check if the request body is empty
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please add the car detail")
		return
	}

	// Decode JSON body into the car model
	var car model.Garage
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error decoding JSON: " + err.Error())
		return
	}

	// Check for empty fields in the JSON
	if car.OwnerName == "" || car.CarNumber == "" || car.ModalName == "" {
		json.NewEncoder(w).Encode("Some fields are empty in JSON")
		return
	}

	// Insert the car into the database
	insertedID, err := insertOneCar(car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error inserting car into the database: " + err.Error())
		return
	}

	// Respond with the inserted car and its ID
	car.ID = insertedID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

/*func InsertOneCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/cars-detail")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Add the car detail")
	}

	var car model.Garage

	_ = json.NewDecoder(r.Body).Decode(&car)

	//_, _ = json.Marshal(&car)
	//checkNilError(err)
	//check for empty field in json
	if car.OwnerName == "" || car.CarNumber == "" || car.ModalName == "" {
		json.NewEncoder(w).Encode("Some field are empty in json")
		return
	}
	insertOneCar(car)

	json.NewEncoder(w).Encode(car)

}*/

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
