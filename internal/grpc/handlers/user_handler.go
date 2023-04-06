package handlers

import (
	"context"
	"fmt"

	pb "github.com/JooHyeongLee/clean-gin/api/proto/user" // user proto 파일의 패키지를 임포트
)

type userHandler struct {
	pb.UnimplementedUserServer
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

// GetUser returns user message by user_id
func (u *userHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var userMessage *pb.UserMessage
	fmt.Println("hello Get User")
	return &pb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}
