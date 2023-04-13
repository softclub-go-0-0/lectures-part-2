package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Making GET request to test todo-service...")
	resp, err := http.Get("http://13.235.247.96:4000/tasks")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nResponse headers:")
	for k, v := range resp.Header {
		fmt.Println(k, "\t", v)
	}

	fmt.Println("\nResponse body:")
	defer resp.Body.Close()
	bodyOriginal, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(body)
	fmt.Println(string(bodyOriginal))

	var bodyAsObject interface{}
	json.Unmarshal(bodyOriginal, &bodyAsObject)

	fmt.Println("\nUnmarshalled response body:")
	fmt.Println(bodyAsObject)

	bodyPrettified, err := json.MarshalIndent(bodyAsObject, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nMarshalled response body with indents:")
	//fmt.Println(bodyPrettified)
	fmt.Println(string(bodyPrettified))
}
