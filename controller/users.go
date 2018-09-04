package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/oneArc_backend/model"
)

// func CreatePerson(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)

// }

// CreatePerson is a method to create a user entry in the db
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "onearc")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	vars := mux.Vars(r)

	fmt.Println("this is vars in users.go in controller", vars)
	firstname := vars["firstname"]
	lastname := vars["lastname"]
	password := vars["password"]
	email := vars["email"]

	db.Create(&model.User{Firstname: firstname, Lastname: lastname, Email: email, Password: password})
	fmt.Fprintf(w, "New user is created.")
}
