package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//User STRUCT DEFINING
type User struct {
	Id       bson.ObjectId `json:"Id"`
	Name     string        `json:"Name"`
	Email    string        `json:"Email"`
	Password string        `json:"Password"`
}

//Post STRUCT DEFINING
type Post struct {
	Id               bson.ObjectId `json:"Id"`
	Caption          string        `json:"Caption"`
	Image_URL        string        `json:"Image URL"`
	Posted_Timestamp string        `json:"Posted Timestamp"`
}

// CALLING CRUD Operations
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users/{id}", FindUserEndpoint).Methods("GET")

	r.HandleFunc("/posts", AllPostsEndPoint).Methods("GET")
	r.HandleFunc("/posts", CreatePostEndPoint).Methods("POST")
	r.HandleFunc("posts/{id}", FindPostEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
