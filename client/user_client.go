// pkg/client/user_client.go

package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"


	userpb "github.com/homasy/pkg/shared/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// UserClient is a client for the user service
type UserClient struct {
	client     userpb.UserServiceClient
	conn       *grpc.ClientConn
	serverAddr string
	mu         sync.Mutex
	connected  bool
}

// NewUserClient creates a new user client
func NewUserClient(serverAddr string) *UserClient {
	return &UserClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the User service
func (c *UserClient) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connected {
		return nil
	}

	// Set up connection with retry
	var err error
	var conn *grpc.ClientConn
	
	// Retry options
	maxRetries := 5
	retryDelay := 2 * time.Second
	
	for i := 0; i < maxRetries; i++ {
		// Connect with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		
		conn, err = grpc.DialContext(
			ctx,
			c.serverAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		
		if err == nil {
			break
		}
		
		log.Printf("Failed to connect to user service (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2 // Exponential backoff
	}
	
	if err != nil {
		return fmt.Errorf("failed to connect to user service after %d attempts: %v", maxRetries, err)
	}
	
	c.conn = conn
	c.client = userpb.NewUserServiceClient(conn)
	c.connected = true
	
	log.Printf("Connected to user service at %s", c.serverAddr)
	return nil
}

// Close closes the connection to the user service
func (c *UserClient) CloseUserConnection() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection to patient service: %v", err)
	}

	c.connected = false
	return nil
}

// GetUser gets a user by ID
func (c *UserClient) GetUser(ctx context.Context, userID string) (*userpb.GetUserResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	req := &userpb.GetUserRequest{
		Id: userID,
	}

	return c.client.GetUser(ctx, req)
}

// CheckUserExists checks if a user exists
func (c *UserClient) CheckUserExists(ctx context.Context, userID string) (bool, error) {
	user, err := c.GetUser(ctx, userID)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}