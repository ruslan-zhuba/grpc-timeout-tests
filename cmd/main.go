package main

import (
	"context"
	pb "github.com/ruslan-zhuba/grpc-timeout-tests/pb/example"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var counter int64 = 0

type server struct {
	pb.UnimplementedExampleServer
}

func (s *server) Ping(_ context.Context, _ *pb.Empty) (*pb.ResponseOk, error) {
	log.Printf("Received ping call #%v", counter)
	counter++
	return &pb.ResponseOk{Ok: true, Counter: counter}, nil
}

func (s *server) PingWithDelay(_ context.Context, in *pb.Delay) (*pb.ResponseOk, error) {
	c := counter
	log.Printf("Received ping call #%v with delay: %v ms", c, in.GetTimeout())
	counter++
	sleepTime := time.Duration(int64(time.Millisecond) * int64(in.GetTimeout()))
	time.Sleep(sleepTime)
	log.Printf("Send response #%v  after %v ms", c, in.GetTimeout())
	return &pb.ResponseOk{Ok: true, Counter: c}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	s := grpc.NewServer()
	pb.RegisterExampleServer(s, &server{})
	log.Println("Start server at port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("can't run server: %v", err)
	}

}
