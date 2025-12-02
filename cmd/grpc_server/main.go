package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lva100/go-grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	note_v1.UnimplementedNoteV1Server
}

func (s *server) Get(ctx context.Context, req *note_v1.GetRequest) (*note_v1.GetResponse, error) {
	return &note_v1.GetResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	note_v1.RegisterNoteV1Server(s, &server{})

	log.Printf("Server listining at %s", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
