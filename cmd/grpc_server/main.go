package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lva100/go-grpc/internal/config"
	"github.com/lva100/go-grpc/internal/config/env"
	"github.com/lva100/go-grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	note_v1.UnimplementedNoteV1Server
}

func (s *server) Get(ctx context.Context, req *note_v1.GetRequest) (*note_v1.GetResponse, error) {
	return &note_v1.GetResponse{}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Could`t load config file: %s", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %s", err)
	}
	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %s", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)

	note_v1.RegisterNoteV1Server(s, &server{})

	log.Printf("Server listining at %s", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
