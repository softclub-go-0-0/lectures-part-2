package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Chose what to do:")
	fmt.Println("\t0. Exit")
	fmt.Println("\t1. Get the list of tasks from todo-service")
	fmt.Println("\t2. Create a task at todo-service")
	fmt.Println("\t3. Update a task at todo-service")
	fmt.Println("\t4. Delete a task at todo-service")
	fmt.Print("Your command: ")
	var command int
	fmt.Scan(&command)
	switch command {
	case 0:
		fmt.Println("Exiting...")
		return
	case 1:
		fmt.Println("Making GET request to the todo-service...")
		resp, err := http.Get("http://127.0.0.1:4000/")
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
		fmt.Printf("%v\t%T", bodyAsObject, bodyAsObject)

		bodyPrettified, err := json.MarshalIndent(bodyAsObject, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nMarshalled response body with indents:")
		//fmt.Println(bodyPrettified)
		fmt.Println(string(bodyPrettified))
	case 2:
		fmt.Print("\nTo create a task enter it's text: ")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		text := stdin.Text()

		fmt.Println("Making POST request to the todo-service...")
		//resp, err := http.PostForm("http://13.235.247.96:4000/tasks", url.Values{"text": {text}})

		requestBody, err := json.Marshal(map[string]string{"text": text})
		if err != nil {
			log.Fatal(err)
		}

		requestBuffer := bytes.NewReader(requestBody)

		resp, err := http.Post("http://13.235.247.96:4000/tasks", "application/json", requestBuffer)
		if err != nil {
			log.Fatal(err)
		}
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

}
