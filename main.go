package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/josepnapitupulu/API_Martumpol/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterMartumpolsRoutes(r)
	http.Handle("/", r)

	fmt.Print("Starting Server localhost:7076")
	log.Fatal(http.ListenAndServe("localhost:7076", r))
}