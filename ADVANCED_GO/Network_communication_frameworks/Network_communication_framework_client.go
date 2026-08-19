package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func ProcessMessage(client *http.Client, request *http.Request) {
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error)
	} else {
		var response Response
		err := json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(response.Message)
		}
	}
}

func SuccessRequest(frameworkAddress string) *http.Request {
	request, _ := http.NewRequest(
		"POST",
		frameworkAddress+"/users/new-user",
		bytes.NewBuffer([]byte(`{"name": "User", "id": 1}`)),
	)
	return request
}

func Error1Request(frameworkAddress string) *http.Request {
	request, _ := http.NewRequest(
		"GET",
		frameworkAddress+"/error/try-error-1",
		nil,
	)
	return request
}

func Error2Request(frameworkAddress string) *http.Request {
	request, _ := http.NewRequest(
		"GET",
		frameworkAddress+"/error/try-error-2",
		nil,
	)
	return request
}

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	ginAddress := "http://localhost:8080"
	echoAddress := "http://localhost:8081"
	fiberAddress := "http://localhost:8082"

	ginSucReq := SuccessRequest(ginAddress)
	fmt.Println("[GIN] sending suc messages ...")
	ProcessMessage(client, ginSucReq)
	echoSucReq := SuccessRequest(echoAddress)
	fmt.Println("[ECHO] sending suc messages ...")
	ProcessMessage(client, echoSucReq)
	fiberSucReq := SuccessRequest(fiberAddress)
	fmt.Println("[FIBER] sending suc messages ...")
	ProcessMessage(client, fiberSucReq)

	ginFailReq1 := Error1Request(ginAddress)
	fmt.Println("[GIN] sending fail1 messages ...")
	ProcessMessage(client, ginFailReq1)
	echoFailReq1 := Error1Request(echoAddress)
	fmt.Println("[ECHO] sending fail1 messages ...")
	ProcessMessage(client, echoFailReq1)
	fiberFailReq1 := Error1Request(fiberAddress)
	fmt.Println("[FIBER] sending fail1 messages ...")
	ProcessMessage(client, fiberFailReq1)

	ginFailReq2 := Error2Request(ginAddress)
	fmt.Println("[GIN] sending fail2 messages ...")
	ProcessMessage(client, ginFailReq2)
	echoFailReq2 := Error2Request(echoAddress)
	fmt.Println("[ECHO] sending fail2 messages ...")
	ProcessMessage(client, echoFailReq2)
	fiberFailReq2 := Error2Request(fiberAddress)
	fmt.Println("[FIBER] sending fail2 messages ...")
	ProcessMessage(client, fiberFailReq2)

}
