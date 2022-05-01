package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Dancheg97/gRPC_api_example/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type server struct {
	pb.UnimplementedTestServerServer
}

func (s *server) Unary(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return in, nil
}

func (s *server) StreamOut(in *pb.Message, stream pb.TestServer_StreamOutServer) error {
	fmt.Println(in)
	for _, i := range []int32{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		err := stream.Send(&pb.Message{
			Id:      i,
			Message: `hello there`,
			Content: []byte{1, 2, 3, 4, 5},
		})
		if err != nil {
			fmt.Println(err)
			return nil
		}
		time.Sleep(time.Second)
	}
	return nil
}
