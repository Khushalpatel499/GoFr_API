# GoFr_API
A Car Garage Management System with CRUD opertion with go language using MongoDB altas to store database and  
use Thunder for testing json query for api to add detail of car in garage.
# Installation
    1.Clone the repository  
        git clone https://github.com/Khushalpatel499/GoFr_API.git
    2.Navigate to the directory
    3.cd GoFr_API
    4.Build and run 
# Run 
 1.Use Gorilla mux for routing system:     
        go get -u github.com/gorilla/mux  
 2.work with mongo driver to add database:        
       go get go.mongodb.org/mongo-driver/mongo
       
## API

1.localhost:5000/api/cars :   
      GET : Get all car detail   
2./api/car:   
      POST : Enter a new car detail   
3./api/cars/id:   
     PUT : Update a car detail  
4./api/cars/id:   
     DELETE: delete a car detail   
5./api/cars/deleteallcars:
     DELETE : delete all cars   
