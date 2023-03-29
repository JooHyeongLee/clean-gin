package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/JooHyeongLee/clean-gin/bootstrap"
	pb "github.com/JooHyeongLee/clean-gin/proto/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedUserServer
}

// GetUser returns user message by user_id
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var userMessage *pb.UserMessage
	fmt.Println("hello Get User")
	return &pb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}

func main() {

	// Start gRPC server
	go func() {
		flag.Parse()
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterUserServer(s, &server{})
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	_ = godotenv.Load()
	// rest server
	bootstrap.RootApp.Execute()
}
