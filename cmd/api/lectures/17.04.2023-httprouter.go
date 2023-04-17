package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"lecture14/pkg/models"
	"log"
	"net/http"
)

const (
	DBHost     = "localhost" // 127.0.0.1
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "blog_service"
)

type handlers struct {
	DB *gorm.DB
}

func newHandlers(db *gorm.DB) (h handlers) {
	return handlers{DB: db}
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected")

	h := handlers{DB: db}

	db.AutoMigrate(
		&models.User{},
	)

	router := httprouter.New()

	router.GET("/users", h.GetUsers)
	router.POST("/users", h.CreateUser)
	router.GET("/users/:id", h.GetUser)
	router.PUT("/users/:id", h.UpdateUser)
	router.DELETE("/users/:id", h.DeleteUser)

	fmt.Println("starting server at :4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func (h *handlers) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var users []models.User

	h.DB.Find(&users)
	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(usersJSON)
}

func (h *handlers) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to see a task %s of user %s\n", ps.ByName("id"), ps.ByName("id"))
}

func (h *handlers) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)

	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	fmt.Println("test 2")
	res := h.DB.Create(&user)
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	fmt.Println("test 3")

	w.Write([]byte("Ok"))
}

func (h *handlers) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to update %s\n", ps.ByName("id"))
}

func (h *handlers) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to delete %s\n", ps.ByName("id"))
}
