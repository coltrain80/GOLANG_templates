// api_template.go
/*
Generic API Template using Go
-----------------------------

This script sets up a basic API using the standard `net/http` package in Go.
It defines a generic API server with endpoints that can be extended as needed.

Quick Reference for Expanding the API:
--------------------------------------

1. **Adding New Routes:**
   - To add a new endpoint, define a new handler function and map it to a route using `http.HandleFunc`.
   - Example:
     func newEndpointHandler(w http.ResponseWriter, r *http.Request) {
         // Define the logic for this endpoint
         fmt.Fprintln(w, "This is a new endpoint")
     }
     http.HandleFunc("/new_endpoint", newEndpointHandler)

2. **Encapsulating Further with New Structs and Methods:**
   - Create additional structs or methods to handle different parts of the API (e.g., separate handlers for different resources).
   - Example:
     type UserAPI struct {
         // Define fields for UserAPI
     }

     func (api *UserAPI) getUsersHandler(w http.ResponseWriter, r *http.Request) {
         // Logic for handling users
         fmt.Fprintln(w, "Users: []")
     }

3. **Adding New Methods:**
   - Define new methods within the structs to encapsulate functionality.
   - Example:
     func (api *UserAPI) processUserData(data string) string {
         // Perform operations on data
         return processedData
     }

4. **Integrating Database:**
   - Use a Go database/sql driver to interact with a database.
   - Example:
     import "database/sql"
     db, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")

5. **Implementing Error Handling:**
   - Add error handling in handler functions and return appropriate HTTP status codes.
   - Example:
     http.Error(w, "Not Found", http.StatusNotFound)

6. **Adding Middleware:**
   - Use middleware functions for cross-cutting concerns like logging, authentication, etc.
   - Example:
     func loggingMiddleware(next http.Handler) http.Handler {
         return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
             log.Println(r.RequestURI)
             next.ServeHTTP(w, r)
         })
     }

7. **Authentication and Authorization:**
   - Integrate JWT, OAuth2, or another method to protect routes.
   - Example:
     import "github.com/dgrijalva/jwt-go"

Usage:
    1. Install Go if not already installed: https://golang.org/dl/
    2. Run the script: go run api_template.go
    3. Access the API at http://localhost:8080/

Example:
    You can use this template to build your own API by extending the GenericAPI struct.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GenericAPI defines the structure for our API server
type GenericAPI struct {
	Port string
}

// NewGenericAPI creates a new instance of GenericAPI
func NewGenericAPI(port string) *GenericAPI {
	return &GenericAPI{Port: port}
}

// SetupRoutes defines the routes for the API
func (api *GenericAPI) SetupRoutes() {
	http.HandleFunc("/", api.indexHandler)
	http.HandleFunc("/data", api.getDataHandler)
	http.HandleFunc("/data", api.createDataHandler).Methods("POST")
	http.HandleFunc("/data/", api.updateDataHandler).Methods("PUT")
	http.HandleFunc("/data/", api.deleteDataHandler).Methods("DELETE")
}

// indexHandler is the default route for the API
func (api *GenericAPI) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Generic API!")
}

// getDataHandler is an example GET endpoint to retrieve data
func (api *GenericAPI) getDataHandler(w http.ResponseWriter, r *http.Request) {
	sampleData := map[string]interface{}{
		"id":          1,
		"name":        "Sample Data",
		"description": "This is some sample data.",
	}
	json.NewEncoder(w).Encode(sampleData)
}

// createDataHandler is an example POST endpoint to create data
func (api *GenericAPI) createDataHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// updateDataHandler is an example PUT endpoint to update data
func (api *GenericAPI) updateDataHandler(w http.ResponseWriter, r *http.Request) {
	dataID := r.URL.Path[len("/data/"):]
	var updatedData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedData["id"] = dataID
	json.NewEncoder(w).Encode(updatedData)
}

// deleteDataHandler is an example DELETE endpoint to delete data
func (api *GenericAPI) deleteDataHandler(w http.ResponseWriter, r *http.Request) {
	dataID := r.URL.Path[len("/data/"):]
	response := map[string]string{"message": fmt.Sprintf("Data with ID %s deleted successfully.", dataID)}
	json.NewEncoder(w).Encode(response)
}

// Run starts the API server
func (api *GenericAPI) Run() {
	fmt.Printf("Starting server on port %s\n", api.Port)
	log.Fatal(http.ListenAndServe(api.Port, nil))
}

// Main function to start the server
func main() {
	api := NewGenericAPI(":8080")
	api.SetupRoutes()
	api.Run()
}
