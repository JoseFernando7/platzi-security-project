package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Another comment
func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received POST request!")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)

		return
	}

	fmt.Println("Body:", string(body))
}

// Comment
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", postHandler).Methods("POST")

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err.Error())
	}
}
