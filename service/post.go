package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Leoff00/go-crud-api/db"
	"github.com/Leoff00/go-crud-api/models"
)

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	payload, _ := io.ReadAll(r.Body)
	var res models.User
	err := json.Unmarshal(payload, &res)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(&res)

	result := db.GetDatabase().Create(&res)
	// fmt.Fprintln(w, string(payload))
	fmt.Println(result)
}
