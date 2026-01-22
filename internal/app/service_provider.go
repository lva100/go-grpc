package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lva100/go-grpc/internal/api/note"
	"github.com/lva100/go-grpc/internal/closer"
	"github.com/lva100/go-grpc/internal/config"
	"github.com/lva100/go-grpc/internal/config/env"
	"github.com/lva100/go-grpc/internal/repository"
	noteRepository "github.com/lva100/go-grpc/internal/repository/note"
	"github.com/lva100/go-grpc/internal/service"
	noteService "github.com/lva100/go-grpc/internal/service/note"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool         *pgxpool.Pool
	noteRepository repository.NoteRepository

	noteService service.NoteService

	noteImpl *note.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err)
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err)
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err)
		}
		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})
		s.pgPool = pool
	}
	return s.pgPool
}

func (s *serviceProvider) NoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.PgPool(ctx))
	}
	return s.noteRepository
}

func (s *serviceProvider) NoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.NoteRepository(ctx))
	}
	return s.noteService
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.NoteService(ctx))
	}
	return s.noteImpl
}
