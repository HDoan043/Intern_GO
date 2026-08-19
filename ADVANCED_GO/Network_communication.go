package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// User struct
type Elements struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result struct {
	Res int `json:"result"`
}

func main() {
	//-------------------------------------------------------------------------
	// SERVER SIDE
	//	New mux
	mux := http.NewServeMux()
	//	Attatch handler to the mux
	mux.HandleFunc(
		// API
		"/api/sum/",
		// handler
		func(w http.ResponseWriter, r *http.Request) {
			// Read the body of the request and load it into a struct object
			var elements Elements
			json.NewDecoder(r.Body).Decode(&elements)
			// Calculate result
			result := elements.A + elements.B
			// Convert result to string json to write in response's body
			json_result := fmt.Sprintf(`{"result": %d}`, result)
			// Write the response
			w.Write([]byte(json_result))
		},
	)
	//	Open a port to listening to the request to the API "/api/sum/" in an independent goroutine
	go http.ListenAndServe(":8080", mux)
	//	Main goroutine waits for the port opened completely
	time.Sleep(5 * time.Second)

	//-------------------------------------------------------------------------
	// CLIENT SIDE
	//	New client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	//	Context
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	//	Http request
	request, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"http://localhost:8080/api/sum/",
		bytes.NewBuffer([]byte(`{"a": 5, "b": 10}`)),
	)

	// Send request in the main goroutine
	if err != nil {
		fmt.Println("Error when create request", err)
	} else {
		// start := time.Now()
		resp, er := client.Do(request)
		if er != nil {
			fmt.Println("Error when server response: ", er)
		} else {
			defer resp.Body.Close()
			// Get result from server's response
			var result Result
			json.NewDecoder(resp.Body).Decode(&result)
			fmt.Println("Result from server: ", result)
		}
	}

}
