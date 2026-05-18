
// pkg/grpc/client.go
package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGRPCConnection creates a new gRPC client connection
func NewGRPCConnection(addr string) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error

	maxRetries := 5
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err = grpc.DialContext(
			ctx,
			addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)

		if err == nil {
			log.Printf("Connected to gRPC server at %s", addr)
			return conn, nil
		}

		log.Printf("Failed to connect to gRPC server at %s (attempt %d/%d): %v", addr, i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2 // Exponential backoff
	}

	return nil, fmt.Errorf("failed to connect to gRPC server at %s after %d attempts: %v", addr, maxRetries, err)
}
