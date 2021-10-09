package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	Server   string
	Database string
}
type PostsDAO struct {
	Server   string
	Database string
}

var admin *mgo.Database

const (
	COLLECTION = "users"
)
const (
	collection = "posts"
)

//MONGODB  CONNECTION
func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	admin = session.DB(m.Database)
}

//MONGODB Insert user DAO
func (m *UsersDAO) Insert(user User) error {
	err := admin.C(COLLECTION).Insert(&user)
	return err
}

//MONGODB Insert post DAO
func (m *UsersDAO) insert(post Post) error {
	err := admin.C(collection).Insert(&post)
	return err
}

// MONGODB FindALL users DAO
func (m *UsersDAO) FindAll() ([]User, error) {
	var users []User
	err := admin.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// MONGODB FindALL posts DAO
func (m *UsersDAO) findAll() ([]Post, error) {
	var posts []Post
	err := admin.C(collection).Find(bson.M{}).All(&posts)
	return posts, err
}

// MONGODB FindById user DAO
func (m *UsersDAO) FindById(id string) (User, error) {
	var user User
	err := admin.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// MONGODB FindById post DAO
func (m *UsersDAO) findbyid(id string) (Post, error) {
	var post Post
	err := admin.C(collection).FindId(bson.ObjectIdHex(id)).One(&post)
	return post, err
}
