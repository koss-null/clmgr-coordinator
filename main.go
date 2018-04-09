package main

import (
	"fmt"
	"myproj.com/clmgr-coordinator/pkg/connection"
)

func main() {
	fmt.Println("Started coordinator")
	go connection.RunServer()
	fmt.Println("Finished corinator")
}
