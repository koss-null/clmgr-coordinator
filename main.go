package main

import (
	"./pkg/connection"
	"fmt"
)

func main() {
	fmt.Println("Started coordinator")
	go connection.RunServer()
	fmt.Println("Finished corinator")
}
