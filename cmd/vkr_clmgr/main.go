package main

import (
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("")
	// if err != nil {
	// 	...
	// }
	client := pb.NewRouteGuideClient(conn)
	defer conn.Close()
}
