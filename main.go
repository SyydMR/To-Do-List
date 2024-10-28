package main

import (
	"log"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SyydMR/To-Do-List/routes"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	routes.RegisterToDoListRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", r))
}