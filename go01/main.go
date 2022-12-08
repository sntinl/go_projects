package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go01/db"
	"go01/models"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := models.Product{
		Name:        "Chocolate",
		Price:       12.00,
		Description: "Chocolate",
	}
	db.DB.Create(&p)
	fmt.Fprintf(w, "Hello Universe")
}
func main() {
	db.DbConnnection()
	err := db.DB.AutoMigrate(models.Product{})
	if err != nil {
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
