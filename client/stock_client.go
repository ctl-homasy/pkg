// pkg/client/user_client.go

package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	supplypb "github.com/homasy/pkg/shared/supply-chain-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// StockClient is a client for the supply chain service
type StockClient struct {
	client     supplypb.SupplyChainServiceClient
	conn       *grpc.ClientConn
	serverAddr string
	mu         sync.Mutex
	connected  bool
}

// NewStockClient creates a new stock client
func NewStockClient(serverAddr string) *StockClient {
	return &StockClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the Supply Chain service
func (c *StockClient) Connect() error {
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

		log.Printf("Failed to connect to supply chain service (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2 // Exponential backoff
	}

	if err != nil {
		return fmt.Errorf("failed to connect to supply chain service after %d attempts: %v", maxRetries, err)
	}

	c.conn = conn
	c.client = supplypb.NewSupplyChainServiceClient(conn)
	c.connected = true

	log.Printf("Connected to supply chain service at %s", c.serverAddr)
	return nil
}

// Close closes the connection to the supply chain service
func (c *StockClient) CloseSupplyChainConnection() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection to supply chain service: %v", err)
	}

	c.connected = false
	return nil
}

// Reduce Stock Item
func (c *StockClient) ReduceStockItem(ctx context.Context, req *supplypb.ReduceStockItemRequest) (*supplypb.ReduceStockItemResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.ReduceStockItem(ctx, req)
}

func (c *StockClient) ReduceStockFromStore(ctx context.Context, req *supplypb.ReduceStockFromStoreRequest) (*supplypb.ReduceStockFromStoreResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.ReduceStockFromStore(ctx, req)
}

func (c *StockClient) ManualStockUpdate(ctx context.Context, req *supplypb.ManualStockUpdateRequest) (*supplypb.ManualStockUpdateResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.ManualStockUpdate(ctx, req)
}

// ListStockItems lists stock items
func (c *StockClient) ListStockItems(ctx context.Context, req *supplypb.ListStockItemsRequest) (*supplypb.ListStockItemsResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.ListStockItems(ctx, req)
}

func (c *StockClient) DeleteStockItem(ctx context.Context, req *supplypb.DeleteStockItemRequest) (*supplypb.DeleteStockItemResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.DeleteStockItem(ctx, req)
}

func (c *StockClient) CreateStockItem(ctx context.Context, req *supplypb.CreateStockItemRequest) (*supplypb.CreateStockItemResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.CreateStockItem(ctx, req)
}

func (c *StockClient) GetStockItem(ctx context.Context, req *supplypb.GetStockItemRequest) (*supplypb.GetStockItemResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.GetStockItem(ctx, req)
}

func (c *StockClient) UpdateStockItem(ctx context.Context, req *supplypb.UpdateStockItemRequest) (*supplypb.UpdateStockItemResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.UpdateStockItem(ctx, req)
}

func (c *StockClient) CreateRequisition(ctx context.Context, req *supplypb.CreateRequisitionRequest) (*supplypb.CreateRequisitionResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.CreateRequisition(ctx, req)
}

// GetMedicalScheme retrieves a medical scheme by ID
func (c *StockClient) GetMedicalScheme(ctx context.Context, id int32) (*supplypb.MedicalScheme, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	req := &supplypb.GetMedicalSchemeRequest{Id: id}
	resp, err := c.client.GetMedicalScheme(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get medical scheme: %v", err)
	}

	return resp.GetMedicalScheme(), nil
}

// GetStore retrieves a store by its ID.
func (c *StockClient) GetStore(ctx context.Context, req *supplypb.GetStoreRequest) (*supplypb.GetStoreResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.GetStore(ctx, req)
}

// ListStores retrieves a list of all stores.
func (c *StockClient) ListStores(ctx context.Context, req *supplypb.ListStoresRequest) (*supplypb.ListStoresResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.ListStores(ctx, req)
}

// GetItemQuantityInStore gets the quantity of an item in a specific store
func (c *StockClient) GetItemQuantityInStore(ctx context.Context, req *supplypb.GetItemQuantityInStoreRequest) (*supplypb.GetItemQuantityInStoreResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.GetItemQuantityInStore(ctx, req)
}

// GetStoreStockQuantities gets all stock quantities for a specific store
func (c *StockClient) GetStoreStockQuantities(ctx context.Context, req *supplypb.GetStoreStockQuantitiesRequest) (*supplypb.GetStoreStockQuantitiesResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.GetStoreStockQuantities(ctx, req)
}
