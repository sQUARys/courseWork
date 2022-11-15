package main

import (
	"courseWork/app/services"
	"fmt"
)

func main() {
	service := services.New()

	service.FillByRand(10)

	fmt.Println(service.Numbers)

	service.FillFromFile("test.txt")
	fmt.Println(service.Numbers)

}
