package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	noteAPI "github.com/lva100/go-grpc/internal/api/note"
	"github.com/lva100/go-grpc/internal/config"
	"github.com/lva100/go-grpc/internal/config/env"
	"github.com/lva100/go-grpc/internal/repository/note"
	"github.com/lva100/go-grpc/internal/service"
	"github.com/lva100/go-grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

/*type server struct {
	note_v1.UnimplementedNoteV1Server
	noteService service.NoteService
}

func (s *server) Create(ctx context.Context, req *note_v1.CreateRequest) (*note_v1.CreateResponse, error) {
	id, err := s.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("Inserted note with id: %d", id)

	return &note_v1.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Get(ctx context.Context, req *note_v1.GetRequest) (*note_v1.GetResponse, error) {
	noteObj, err := s.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &note_v1.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}*/

// func handleNullTime(tm sql.NullTime) string {
// 	if tm.Valid {
// 		return tm.Time.Format("02-01-2006")
// 	} else {
// 		return "-/-/-"
// 	}
// }

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

	noteRepo := note.NewRepository(pool)
	noteSrv := service.NoteService(noteRepo)

	s := grpc.NewServer()
	reflection.Register(s)

	note_v1.RegisterNoteV1Server(s, noteAPI.NewImplementation(noteSrv))

	log.Printf("Server listining at %s", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
