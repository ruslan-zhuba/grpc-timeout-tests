package main

import (
	"context"
	"github.com/ruslan-zhuba/grpc-timeout-tests/handlers"
	pb "github.com/ruslan-zhuba/grpc-timeout-tests/pb/example"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	grpcServer := grpc.NewServer()
	pb.RegisterExampleServer(grpcServer, &server{})
	log.Println("Start server at port :50051")

	r := handlers.Router()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	httpServer := &http.Server{
		Addr:    ":80",
		Handler: r,
	}

	shutdown := make(chan struct{}, 1)
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			log.Printf("%v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			shutdown <- struct{}{}
			log.Printf("%v", err)
		}
	}()
	log.Print("The service is ready to listen and serve.")

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}
	case <-shutdown:
		log.Printf("Got an error...")
	}

	log.Print("The service is shutting down...")
	httpServer.Shutdown(context.Background())
	grpcServer.GracefulStop()
	log.Print("Done")

}
