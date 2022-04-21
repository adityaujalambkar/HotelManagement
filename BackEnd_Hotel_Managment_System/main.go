package main

import (
	"backend_hotel_management_system/controllers"
	"log"
	"net/http"
)

func main() {
	initDB()

	router := mux.NewRouter()
	router.HandleFunc("*", Cors)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", setHeaders(router)))
}
func initaliseHandlers(router *mux.Router) {

	router.HandleFunc("/user/login", controllers.Login).Methods("POST")

}

func initDB() {

}
