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
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
	//
	var (
		wArea      sql.NullString
		agentCount int
	)

	rows, err := db.Query("select working_area, count(*), min(commission) from agents group by working_area")
	if err != nil {
		log.Fatal(err)
	}

	cnt := 0
	fmt.Println("\nAreas with number of agents in there")
	var comm sql.NullFloat64
	for rows.Next() {
		err := rows.Scan(&wArea, &agentCount, &comm)
		if err != nil {
			log.Fatal(err)
		}
		//if !wArea.Valid {
		//	// not string, i.e. it is NULL
		//}
		//if !comm.Valid {
		//	// not float, i.e. it is NULL
		//}
		fmt.Printf("%s: %d\t%f\n", wArea.String, agentCount, comm.Float64)
		cnt++
	}
	fmt.Println(cnt, "rows")

	fmt.Println("Creating new agent. Static string:")
	query := "insert into agents (agent_code, agent_name, working_area, commission, phone_no, country) VALUES ('B074', 'Will Smith', 'Dubai', 0.25, '+992987654321', 'Tajikistan')"
	result, err := db.Exec(query)
	if err != nil {
		log.Fatal(err, result)
	}
	fmt.Println("Success!")
	id, err := result.LastInsertId()
	//if err != nil {
	//	log.Fatal(err)
	//}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID:", id, "\tAffected rows number:", affectedRows)

	var (
		code       string  = "C074"
		name       string  = "Brad Pitt"
		area       string  = "Antalia"
		commission float64 = 0.3
		phone      string  = "+992123456789"
		country    string  = "Tajikistan"
	)

	fmt.Println("\nCreating new agent. Dynamic string:")
	query = "insert into agents (agent_code, agent_name, working_area, commission, phone_no, country) VALUES ($1, $2, $3, $4, $5, $6)"
	result, err = db.Exec(query, code, name, area, commission, phone, country)
	if err != nil {
		log.Fatal(err, result)
	}
	fmt.Println("Success!")
	id, err = result.LastInsertId()
	//if err != nil {
	//	log.Fatal(err)
	//}
	affectedRows, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID:", id, "\tAffected rows number:", affectedRows)

	fmt.Println("\nUpdating existing agent(s). Dynamic string:")
	query = `update agents set agent_name=$1, working_area=$2 where country=$3`

	result, err = db.Exec(query, "Fariz", "Manhattan", "Tajikistan")
	if err != nil {
		log.Fatal(err, result)
	}
	fmt.Println("Success!")
	id, err = result.LastInsertId()
	//if err != nil {
	//	log.Fatal(err)
	//}
	affectedRows, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID:", id, "\tAffected rows number:", affectedRows)

	fmt.Println("\nDeleting existing agent(s). Dynamic string:")
	query = `delete from agents where country is null`

	result, err = db.Exec(query)
	if err != nil {
		log.Fatal(err, result)
	}
	fmt.Println("Success!")
	id, err = result.LastInsertId()
	//if err != nil {
	//	log.Fatal(err)
	//}
	affectedRows, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID:", id, "\tAffected rows number:", affectedRows)
}
