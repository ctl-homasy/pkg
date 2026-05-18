// pkg/client/supply_chain_client.go

package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	supplypb "github.com/ctl-homasy/pkg/shared/supply-chain-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// SupplyChainClient is a client for the supply-chain service
type SupplyChainClient struct {
	client     supplypb.SupplyChainServiceClient
	conn       *grpc.ClientConn
	serverAddr string
	mu         sync.Mutex
	connected  bool
}

// NewSupplyChainClient creates a new supply-chain client
func NewSupplyChainClient(serverAddr string) *SupplyChainClient {
	return &SupplyChainClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the supply-chain service
func (c *SupplyChainClient) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connected {
		return nil
	}

	var err error
	var conn *grpc.ClientConn
	
	maxRetries := 5
	retryDelay := 2 * time.Second
	
	for i := 0; i < maxRetries; i++ {
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
		
		log.Printf("Failed to connect to supply-chain service (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2
	}
	
	if err != nil {
		return fmt.Errorf("failed to connect to supply-chain service after %d attempts: %v", maxRetries, err)
	}
	
	c.conn = conn
	c.client = supplypb.NewSupplyChainServiceClient(conn)
	c.connected = true
	
	log.Printf("Connected to supply-chain service at %s", c.serverAddr)
	return nil
}

// Close closes the connection to the supply-chain service
func (c *SupplyChainClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection to supply-chain service: %v", err)
	}

	c.connected = false
	return nil
}

// GetMedicalScheme gets a medical scheme by ID
func (c *SupplyChainClient) GetMedicalScheme(ctx context.Context, schemeID int32) (*supplypb.GetMedicalSchemeResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	req := &supplypb.GetMedicalSchemeRequest{
		Id: schemeID,
	}

	return c.client.GetMedicalScheme(ctx, req)
}

// GetMedicalSchemeByName gets a medical scheme by name
func (c *SupplyChainClient) GetMedicalSchemeByName(ctx context.Context, name string) (*supplypb.GetMedicalSchemeResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	req := &supplypb.GetMedicalSchemeByNameRequest{
		Name: name,
	}

	return c.client.GetMedicalSchemeByName(ctx, req)
}
