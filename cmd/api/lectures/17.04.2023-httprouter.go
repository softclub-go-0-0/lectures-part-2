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

type handler struct {
	DB *gorm.DB
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected")

	h := handler{DB: db}

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

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var users []models.User
	query := r.URL.Query()
	search := query.Get("name")

	if search != "" {
		h.DB.Where("name ilike ?", "%"+search+"%").Find(&users)
	} else {
		h.DB.Find(&users)
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(usersJSON)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to see a task %s of user %s\n", ps.ByName("id"), ps.ByName("id"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)

	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	res := h.DB.Create(&user)
	if res.Error != nil {
		log.Fatal(res.Error)
	}

	w.Write([]byte(fmt.Sprintf("Successfully created a user. ID: %d", user.ID)))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to update %s\n", ps.ByName("id"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to delete %s\n", ps.ByName("id"))
}
