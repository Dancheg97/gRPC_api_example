package main

import (
	"context"
	"log"
	"net"

	"github.com/Dancheg97/gRPC_api_example/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBasicsServer(s, &server{})
	pb.RegisterConstructionsServer(s, &server{})
	pb.RegisterStreamsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type server struct {
	pb.UnimplementedBasicsServer
	pb.UnimplementedConstructionsServer
	pb.UnimplementedStreamsServer
}

func (s *server) DoubleCall(ctx context.Context, in *pb.DoubleMes) (*pb.DoubleMes, error) {
	return in, nil
}

func (s *server) FloatCall(ctx context.Context, in *pb.FloatMes) (*pb.FloatMes, error) {
	return in, nil
}

func (s *server) Int32Call(ctx context.Context, in *pb.Int32Mes) (*pb.Int32Mes, error) {
	return in, nil
}

func (s *server) Int64Call(ctx context.Context, in *pb.Int64Mes) (*pb.Int64Mes, error) {
	return in, nil
}

func (s *server) Uint32Call(ctx context.Context, in *pb.Uint32Mes) (*pb.Uint32Mes, error) {
	return in, nil
}

func (s *server) Uint64Call(ctx context.Context, in *pb.Uint64Mes) (*pb.Uint64Mes, error) {
	return in, nil
}

func (s *server) Sint32Call(ctx context.Context, in *pb.Sint32Mes) (*pb.Sint32Mes, error) {
	return in, nil
}

func (s *server) Sint64Call(ctx context.Context, in *pb.Sint64Mes) (*pb.Sint64Mes, error) {
	return in, nil
}

func (s *server) Fixed32Call(ctx context.Context, in *pb.Fixed32Mes) (*pb.Fixed32Mes, error) {
	return in, nil
}

func (s *server) Fixed64Call(ctx context.Context, in *pb.Fixed64Mes) (*pb.Fixed64Mes, error) {
	return in, nil
}

func (s *server) Sfixed32Call(ctx context.Context, in *pb.Sfixed32Mes) (*pb.Sfixed32Mes, error) {
	return in, nil
}

func (s *server) Sfixed64Call(ctx context.Context, in *pb.Sfixed64Mes) (*pb.Sfixed64Mes, error) {
	return in, nil
}

func (s *server) BoolCall(ctx context.Context, in *pb.BoolMes) (*pb.BoolMes, error) {
	return in, nil
}

func (s *server) StringCall(ctx context.Context, in *pb.StringMes) (*pb.StringMes, error) {
	return in, nil
}

func (s *server) BytesCall(ctx context.Context, in *pb.BytesMes) (*pb.BytesMes, error) {
	return in, nil
}

// func (s *server) Unary(ctx context.Context, in *pb.Message) (*pb.Message, error) {
// 	return in, nil
// }
