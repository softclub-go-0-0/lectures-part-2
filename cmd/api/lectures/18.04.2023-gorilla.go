package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You see user list\n")
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you try to create new user\n")
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to see user %s\n", vars["id"])
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to update %s\n", vars["login"])
}

func Delete(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	for k, v := range h {
		fmt.Println(k, "\t", v)
	}
	time.Sleep(time.Second * 5)
	fmt.Fprintf(w, "you try to delete user %s", mux.Vars(r)["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetList).Methods("GET")
	r.HandleFunc("/users", Create).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", GetOne).Methods("GET")
	r.HandleFunc("/users/{login}", Update).Methods("PUT").Headers("X-Auth", "test", "anotherOne", "someText")
	r.HandleFunc("/users/{id:[0-9]+}", Delete).Methods("DELETE")

	fmt.Println("starting server at :4007")
	log.Fatal(http.ListenAndServe(":4007", r))
}
