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
	sort := sorts.New()           // create sorts
	service := services.New(sort) // create service

	controller := controller.New(service) // create controller
	router := routers.New(controller)     // create router

	router.SetRoutes()     // set current routes
	server := http.Server{ // create server
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
		Addr:         ":8080",
		Handler:      router.Router,
	}

	err := server.ListenAndServe() // start server
	if err != nil {
		log.Println("Error in main : ", err)
		return
	}
}
