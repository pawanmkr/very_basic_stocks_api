package router

import (
	"github.com/gorilla/mux"
	"stox/middlewares"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middlewares.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middlewares.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", middlewares.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middlewares.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middlewares.DeleteStock).Methods("DELETE", "OPTIONS")
	return router
}
