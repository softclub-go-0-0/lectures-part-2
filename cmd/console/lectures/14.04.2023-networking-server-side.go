package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	teachers := []string{
		"Siyovush",
		"Nurullo",
		"Fariz",
	}

	students := []string{
		"Alijon",
		"Mehrdod",
		"Fozil",
		"Muhammad",
	}

	teachersHandler := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, fmt.Sprint(teachers))
	}

	studentsHandler := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, fmt.Sprint(students))
	}

	var name string
	fmt.Scan(&name)
	// our endpoints:
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, `Some cool html`)
	})
	http.HandleFunc("/teachers", teachersHandler)
	http.HandleFunc("/students", studentsHandler)
	http.HandleFunc("/groups", groupsHandler)

	log.Fatal(http.ListenAndServe(":4000", nil))
}

func groupsHandler(w http.ResponseWriter, _ *http.Request) {
	groups := []string{
		"C#",
		"Go",
		"Golang",
		"Golangjon",
	}

	resp, _ := json.Marshal(groups)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, string(resp))
}
