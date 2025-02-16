package main

import (
	"fmt"
	"net/http"

	"nilushancosta/request-viewer/internal/handlers"
)

func main() {
	http.HandleFunc("/rest", handlers.RESTHandler)

	fmt.Println("Starting HTTP Server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error occurred when starting server")
		return
	}
}
