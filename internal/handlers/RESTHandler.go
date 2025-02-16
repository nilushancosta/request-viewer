package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func RESTHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("REST API REQUEST DETAILS")
	fmt.Println("Method: " + req.Method)

	fmt.Println("\nHEADERS")
	for k, _ := range req.Header {
		fmt.Println(k, req.Header.Get(k))
	}

	fmt.Println("\nQUERY PARAMETERS")
	for k, _ := range req.URL.Query() {
		fmt.Println(k + "=" + req.URL.Query().Get(k))
	}
	
	fmt.Println("\nBODY")
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read request body", err)
	}
	req.Body.Close()
	fmt.Println("Size: " + strconv.Itoa(int(req.ContentLength)) + " bytes")
	fmt.Println("Content:\n" + string(body) + "\n")
}