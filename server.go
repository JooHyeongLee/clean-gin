package main

import (
	"flag"
	"fmt"
	"github.com/JooHyeongLee/clean-gin/api/proto/sample"
	"github.com/JooHyeongLee/clean-gin/api/proto/user"
	"github.com/JooHyeongLee/clean-gin/bootstrap"
	"github.com/JooHyeongLee/clean-gin/internal/grpc/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	// Start gRPC server
	go func() {
		flag.Parse()
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		// grpc 서버 실행
		s := grpc.NewServer()
		// handler 등록
		user.RegisterUserServer(s, handlers.NewUserHandler())
		routeguide.RegisterRouteGuideServer(s, handlers.NewRouteGuideHandler())

		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	_ = godotenv.Load()
	// rest server
	bootstrap.RootApp.Execute()
}
