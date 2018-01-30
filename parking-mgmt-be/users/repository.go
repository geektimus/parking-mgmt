package users

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repository allow us to interact with the mongo database
type Repository struct{}

//TODO change for the docker container name

// SERVER points to the mongo database server with port number
const SERVER = "parking-mgmt-db:27017"

// DBNAME points to the database name
const DBNAME = "parking-mgmt"

// DOCNAME points to the collection name
const DOCNAME = "users"

// GetUsers returns the list of Users
func (r Repository) GetUsers() Users {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Users{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// GetUser returns an User
func (r Repository) GetUser(id string) User {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()

	// Grab id
	oid := bson.ObjectIdHex(id)
	// Find user
	c := session.DB(DBNAME).C(DOCNAME)
	user := User{}
	if err := c.Find(bson.M{"_id": oid}).One(&user); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return user
}

// AddUser inserts an User in the DB
func (r Repository) AddUser(user User) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	user.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(user)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateUser updates an user in the DB (not used for now)
func (r Repository) UpdateUser(user User) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(user.ID, user)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteUser deletes an User (not used for now)
func (r Repository) DeleteUser(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	// Write status
	return "OK"
}
