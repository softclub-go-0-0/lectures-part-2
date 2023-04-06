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
	dbname   = "lectures_part_2_db"
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

	fmt.Println("Which table do you want to see?")
	var table string
	fmt.Scan(&table)

	fmt.Println("Okay, here it is:")

	rows, err := db.Query("select * from " + table)
	if err != nil {
		log.Fatal(err)
	}

	cnt := 0
	switch table {
	case "agents":
		var (
			code       string
			name       string
			area       string
			commission string
			phone      string
			country    string
		)
		for rows.Next() {
			err := rows.Scan(&code, &name, &area, &commission, &phone, &country)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(code, name, area, commission, phone, country)
			cnt++
		}
		fmt.Println(cnt, "rows")
	case "customers":
		var (
			code      string
			name      string
			city      string
			area      string
			country   string
			grade     string
			opAmt     string
			reAmt     string
			paAmt     string
			ouAmt     string
			phone     string
			agentCode string
		)
		for rows.Next() {
			err := rows.Scan(&code, &name, &city, &area, &country, &grade, &opAmt, &reAmt, &paAmt, &ouAmt, &phone, &agentCode)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(code, name, city, area, country, grade, opAmt, reAmt, paAmt, ouAmt, phone, agentCode)
			cnt++
		}
		fmt.Println(cnt, "rows")
	case "orders":
		var (
			number       string
			amount       string
			advAmt       string
			date         string
			customerCode string
			agentCode    string
			description  string
		)
		for rows.Next() {
			err := rows.Scan(&number, &amount, &advAmt, &date, &customerCode, &agentCode, &description)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(number, amount, advAmt, date, customerCode, agentCode, description)
			cnt++
		}
		fmt.Println(cnt, "rows")
	}
}
