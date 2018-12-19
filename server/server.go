package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/srik1309/rndnumgenerator/rndnumgenerator"

	"math/rand"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) SendRandomNumber(ctx context.Context, in *pb.RandomNumberRequest) (*pb.RandomNumberReply, error) {
	var randomnum int32
	randomnum = rand.Int31n(1234567)
	var msg string
	msg = fmt.Sprint(randomnum)
	fmt.Printf("%v Sending %v to the client\n", time.Now(), msg)
	return &pb.RandomNumberReply{Message: msg}, nil
}

func main() {
	fmt.Printf("%v Starting server on port: %v\n", time.Now(), port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterRandomNumberGeneratorServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
