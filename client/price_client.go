// pkg/client/price_client.go

package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	billingpb "github.com/ctl-homasy/pkg/shared/billing-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// PriceClient is a client for the price service
type PriceClient struct {
	client     billingpb.PriceServiceClient
	conn       *grpc.ClientConn
	serverAddr string
	mu         sync.Mutex
	connected  bool
}

// NewPriceClient creates a new price client
func NewPriceClient(serverAddr string) *PriceClient {
	return &PriceClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the price service
func (c *PriceClient) Connect() error {
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
		
		log.Printf("Failed to connect to price service (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2
	}
	
	if err != nil {
		return fmt.Errorf("failed to connect to price service after %d attempts: %v", maxRetries, err)
	}
	
	c.conn = conn
	c.client = billingpb.NewPriceServiceClient(conn)
	c.connected = true
	
	log.Printf("Connected to price service at %s", c.serverAddr)
	return nil
}

// Close closes the connection to the price service
func (c *PriceClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection to price service: %v", err)
	}

	c.connected = false
	return nil
}

// LookupServicePrice gets the price for a given service
func (c *PriceClient) LookupServicePrice(ctx context.Context, req *billingpb.LookupServicePriceRequest) (*billingpb.LookupServicePriceResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.LookupServicePrice(ctx, req)
}

// LookupWardPrice gets the price for a given ward
func (c *PriceClient) LookupWardPrice(ctx context.Context, req *billingpb.LookupWardPriceRequest) (*billingpb.LookupWardPriceResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.LookupWardPrice(ctx, req)
}

// LookupLabTestPrice gets the price for a given lab test
func (c *PriceClient) LookupLabTestPrice(ctx context.Context, req *billingpb.LookupLabTestPriceRequest) (*billingpb.LookupLabTestPriceResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.LookupLabTestPrice(ctx, req)
}
