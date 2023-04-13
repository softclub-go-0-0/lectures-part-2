package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"lecture14/pkg/models"
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

	teacher := models.Teacher{
		//ID:        0,
		Name:    "Siyovush",
		Surname: "Radnoy",
		Phone:   "+992987654321",
		//Email:     sql.NullString{},
		//CreatedAt: sql.NullTime{},
		//UpdatedAt: sql.NullTime{},
		//DeletedAt: sql.NullTime{},
	}

	statement := "insert into teachers (name, surname, phone) VALUES ($1, $2, $3)"
	result, err := db.Exec(statement, teacher.Name, teacher.Surname, teacher.Phone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully added teacher", teacher.Name)
	rowsAffected, _ := result.RowsAffected()
	id, _ := result.LastInsertId()

	fmt.Println("Rows affected:", rowsAffected)
	fmt.Println("ID of inserted teacher:", id)
}
