package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux" //in terminal install, go get -u github.com/gorilla/mux
)

// Define a struct for the data you want to expose
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Slice to store items (in-memory "database")
var items = []Item{
	{ID: 1, Name: "Item 1", Price: 19.99},
	{ID: 2, Name: "Item 2", Price: 29.99},
	{ID: 3, Name: "Item 3", Price: 39.99},
}

// Handler function to get all items
func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {
	myRouter := mux.newRouter().StrictSlash(true)

	// Register handler function for the route "/items"
	//http.HandleFunc("/items", getAllItems)

	myRouter.HandleFunc("/items", getAllItems).Methods("GET")
	myRouter.HandleFunc("/items", getAllItems).Methods("POST")

	// Start the server on port 8081
	port := 8081
	fmt.Printf("Server is listening on :%d...\n", port)
	//err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), myRouter)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
