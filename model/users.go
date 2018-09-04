package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/jinzhu/gorm"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

// delete user
