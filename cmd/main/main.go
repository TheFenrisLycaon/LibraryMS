package main

import (
	"github.com/TheFenrisLycaon/LibraryMS/pkg/routes"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterLibraryRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}
