package handlers

import (
	pb "github.com/JooHyeongLee/clean-gin/api/proto/stream" // user proto 파일의 패키지를 임포트
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

type streamHandler struct {
	pb.UnimplementedStreamServiceServer
}

func NewStreamHandler() *streamHandler {
	return &streamHandler{}
}

func (s *streamHandler) BidirectionalStreaming(stream pb.StreamService_BidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if status.Code(err) == codes.Canceled || err == io.EOF {
				stream.Send(&pb.Response{
					Message: "끗났다!!",
				})
			}
			return err
		}

		log.Printf("Received message: %s", req.Message)

		resp := &pb.Response{
			Message: "Server received message",
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}
