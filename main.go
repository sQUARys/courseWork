package main

import (
	"courseWork/app/controller"
	"courseWork/app/routers"
	"courseWork/app/services"
	"courseWork/app/sorts"
	"log"
	"net/http"
	"time"
)

func main() {
	sort := sorts.New()
	service := services.New(sort)

	controller := controller.New(service)
	router := routers.New(controller)

	router.SetRoutes()
	server := http.Server{
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
		Addr:         ":8080",
		Handler:      router.Router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error in main : ", err)
		return
	}
}
