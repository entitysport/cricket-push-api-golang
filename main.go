package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
//example strucher
type Data struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Decode JSON data
	var receivedData Data
	err = json.Unmarshal(body, &receivedData)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Process the data
	// For this example, we'll simply log the received data
	fmt.Printf("Received JSON data: %+v\n", receivedData)

	// Prepare the response
	response := map[string]string{
		"message": "JSON data received successfully",
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Send the response
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error sending response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/data", handleRequest)

	port := "3000" // Replace with your desired port number
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
