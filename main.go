package main

import (
	"log"
	"net/http"

	"github.com/Leoff00/go-crud-api/db"
	"github.com/Leoff00/go-crud-api/service"
)

func main() {
	go db.StartDB()

	http.HandleFunc("/get", service.Get)
	http.HandleFunc("/post", service.Post)
	http.HandleFunc("/update/:id", service.Update)
	http.HandleFunc("/delete/:id", service.Delete)
	log.Println("ONLINE! 🚀 \n Listening on PORT http://localhost:3001")
	http.ListenAndServe(":3001", nil)
}
