package main

import (
	pb "../../protobuf/compiled/protobuf/simplemsg"
	"fmt"
	"golang.org/x/net/context"
)

var n int32 = 0

func main() {
	fmt.Println("Stub")
}

type simpleServer struct {
}

func (s *simpleServer) GetFeature(ctx context.Context, pingMsg *pb.PingMessage) (*pb.PingMessage, error) {
	f, _ := s.GetFeature(ctx, pingMsg)
	fmt.Printf("Got ping: %s %d", f.Query, f.Number)

	getNextN := func() int32 { n += 1; return n }
	return &pb.PingMessage{"Ping", getNextN()}, nil
}
