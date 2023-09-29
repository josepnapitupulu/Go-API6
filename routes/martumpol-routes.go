package routes

import (
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Martumpol/controllers"
)

var RegisterMartumpolsRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.Index).Methods("GET")
	// router.HandleFunc("/jemaatBaru", controllers.Create).Methods("POST")
	router.HandleFunc("/martumpol/", controllers.CreateMartumpol).Methods("POST")
	router.HandleFunc("/martumpol/", controllers.GetMartumpol).Methods("GET")
	router.HandleFunc("/martumpol/{martumpolId}", controllers.GetMartumpolById).Methods("GET")
	router.HandleFunc("/martumpol/{martumpolId}", controllers.UpdateMartumpol).Methods("PUT")
	router.HandleFunc("/martumpol/{martumpolId}", controllers.DeleteMartumpol).Methods("DELETE")
}
