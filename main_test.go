package main_test

import (
	"encoding/json"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/khushalpatel499/gofr_api/model"
	"github.com/khushalpatel499/gofr_api/router"
)

func TestRouter(t *testing.T) {
	// Coverage target
	coverage := 0.0

	// GET /api/cars
	req, err := http.NewRequest("GET", "/api/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := router.Router()
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code 200 for GET /api/cars, got %d", recorder.Code)
	}
	coverage++

	// POST /api/car
	req, err = http.NewRequest("POST", "/api/car", strings.NewReader(`{"ownername": "John Doe", "modalname": "Toyota Camry", "carnumber": "ABC123"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code 201 for POST /api/car, got %d", recorder.Code)
	}
	coverage++

	// Decode the response body and check the content
	var result model.Garage
	err = json.Unmarshal(recorder.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}

	// Convert the inserted ID to a string
	insertedIDStr := result.ID.Hex()

	// Check if the ID is not empty
	if insertedIDStr == "" {
		t.Error("Expected non-empty ObjectID, got empty ID")
	}

	// PUT /api/cars/{id} with valid data
	req, err = http.NewRequest("PUT", "/api/cars/"+insertedIDStr, strings.NewReader(`{"repair": true}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code 200 for PUT /api/cars/{id}, got %d", recorder.Code)
	}
	coverage++

	// PUT /api/cars/{id} with invalid data
	req, err = http.NewRequest("PUT", "/api/cars/56789", strings.NewReader(`{"invalid_field": "dummy"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400 for PUT /api/cars/{id} with invalid data, got %d", recorder.Code)
	}
	coverage++

	// DELETE /api/cars/{id}
	req, err = http.NewRequest("DELETE", "/api/cars/"+insertedIDStr, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected status code 204 for DELETE /api/cars/{id}, got %d", recorder.Code)
	}
	coverage++

	// Invalid route
	req, err = http.NewRequest("POST", "/invalid/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404 for invalid route, got %d", recorder.Code)
	}
	coverage++

	// Calculate and report coverage
	actualCoverage := (coverage / 6) * 100
	if actualCoverage < 60 {
		t.Errorf("Test coverage fell below target: %f%% (expected 60%%)", actualCoverage)
	}
}
