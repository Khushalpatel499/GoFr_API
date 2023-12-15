package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/khushalpatel499/gofr_api/router"
)

func main() {
	fmt.Println("REST gofr API")
	r := router.Router()

	fmt.Println("Server is getting started....")
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening to port 5000...")
}
