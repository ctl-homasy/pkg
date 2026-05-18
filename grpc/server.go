
// pkg/grpc/server.go
package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer represents a gRPC server
type GRPCServer struct {
	server *grpc.Server
	port   string
	name   string
}

// NewGRPCServer creates a new gRPC server
func NewGRPCServer(port, name string, interceptors ...grpc.UnaryServerInterceptor) *GRPCServer {
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors...),
	)

	return &GRPCServer{
		server: s,
		port:   port,
		name:   name,
	}
}

// RegisterService registers a service with the gRPC server
func (s *GRPCServer) RegisterService(registerFunc func(*grpc.Server)) {
	registerFunc(s.server)
}

// Start starts the gRPC server
func (s *GRPCServer) Start() {
	lis, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	reflection.Register(s.server)

	go func() {
		log.Printf("%s is running on port %s", s.name, s.port)
		if err := s.server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

// Shutdown gracefully shuts down the gRPC server
func (s *GRPCServer) Shutdown() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	log.Printf("Shutting down %s...", s.name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.server.GracefulStop()

	<-ctx.Done()
	log.Printf("%s shutdown complete", s.name)
}
