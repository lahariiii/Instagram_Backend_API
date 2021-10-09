package main

import (
	"dao"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// CREATE POST
func CreatePostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	post.Id = bson.NewObjectId()
	if err := dao.insert(post); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, post)
}

// GET ALL POSTS
func AllPostsEndPoint(w http.ResponseWriter, r *http.Request) {
	posts,
		err := dao.findall()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, posts)
}

// GET POST BY ID
func FindPostEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	post,
		err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Post ID")
		return
	}
	respondWithJson(w, http.StatusOK, post)
}
