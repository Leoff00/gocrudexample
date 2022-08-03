package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Leoff00/go-crud-api/db"
	"github.com/Leoff00/go-crud-api/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	var user models.User
	fmt.Println(user)
	db.GetDatabase().First(&user).Scan(&user)
	bt, _ := json.Marshal(user)
	fmt.Println(string(bt))
}
