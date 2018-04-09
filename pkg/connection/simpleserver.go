package connection

import (
	"fmt"
	"io"
	"net"

	. "myproj.com/clmgr-coordinator/pkg/common"
	pb "myproj.com/clmgr-coordinator/protobuf/compiled/protobuf/ping"

	"google.golang.org/grpc"
)

type simpleServer struct{}

func RunServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 2222))
	if err != nil {
		fmt.Printf("Can't run the server, %s\n", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPingerServer(grpcServer, &simpleServer{})
	grpcServer.Serve(lis)
}

func (s *simpleServer) GetPing(stream pb.Pinger_GetPingServer) error {
	SetGlobalCounter("pingServer", Counter32())
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("Got a ping reqvest %d %s\n", in.Number, in.Query)
		ans := pb.PingMessage{"Ping OK", Count("pingServer").(int32)}
		if err := stream.Send(&ans); err != nil {
			return err
		}
	}
}
