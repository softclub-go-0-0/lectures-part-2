package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost" // 127.0.0.1
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "exam"
)

func main() {
	// data source name (DSN)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//
	var (
		code       string
		name       string
		area       string
		commission string
		phoneNo    string
		country    string
	)

	rows, err := db.Query("SELECT commission, country FROM agents")
	if err != nil {
		log.Fatal(err)
	}

	cnt := 0
	for rows.Next() {
		err := rows.Scan(&commission, &country)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(code, name, area, commission, phoneNo, country)
		cnt++
	}
	fmt.Println(cnt, "rows")
}
