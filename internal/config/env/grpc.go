package env

import (
	"errors"
	"net"
	"os"

	"github.com/lva100/go-grpc/internal/config"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvFile = "GRPC_HOST"
	grpcPortEnvFile = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostEnvFile)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}
	port := os.Getenv(grpcPortEnvFile)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
