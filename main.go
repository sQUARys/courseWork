package main

import (
	"courseWork/app/services"
	"courseWork/app/sorts"
)

func main() {
	sort := sorts.New()
	service := services.New(sort)

	service.FillByRand(10)
	service.StartSorting()

	//fmt.Println(service.Numbers)
	//service.FillFromFile("test.txt")
	//fmt.Println(service.Numbers)

}
