
// pkg/client/billing_client.go
package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/ctl-homasy/pkg/grpc"
	billingpb "github.com/ctl-homasy/pkg/shared/billing-service/proto"
	gogrpc "google.golang.org/grpc"
)

// BillingClient is a client for the Billing service
type BillingClient struct {
	client      billingpb.BillingServiceClient
	priceClient billingpb.PriceServiceClient
	conn        *gogrpc.ClientConn
	serverAddr  string
	mu          sync.Mutex
	connected   bool
}

// NewBillingClient creates a new billing client
func NewBillingClient(serverAddr string) *BillingClient {
	return &BillingClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the billing service
func (c *BillingClient) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connected {
		return nil
	}

	conn, err := grpc.NewGRPCConnection(c.serverAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to Billing service: %v", err)
	}

	c.conn = conn
	c.client = billingpb.NewBillingServiceClient(conn)
	c.priceClient = billingpb.NewPriceServiceClient(conn)
	c.connected = true

	return nil
}

// Close closes the connection to the Billing service
func (c *BillingClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection to billing service: %v", err)
	}

	c.connected = false
	return nil
}

func (c *BillingClient) CreateServiceRecord(ctx context.Context, req *billingpb.CreateServiceRecordRequest) (*billingpb.CreateServiceRecordResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.CreateServiceRecord(ctx, req)
}

func (c *BillingClient) UpdateServiceRecord(ctx context.Context, req *billingpb.UpdateServiceRecordStatusRequest) (*billingpb.UpdateServiceRecordStatusResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.UpdateServiceRecordStatus(ctx, req)
}

func (c *BillingClient) GetPatientPaymentStatus(ctx context.Context, patientID string) (*billingpb.GetPatientPaymentStatusResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	req := &billingpb.GetPatientPaymentStatusRequest{
		PatientId: patientID,
	}
	return c.client.GetPatientPaymentStatus(ctx, req)
}

func (c *BillingClient) GetServiceRecordsByPatient(ctx context.Context, patientID string) (*billingpb.GetServiceRecordsByPatientResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	req := &billingpb.GetServiceRecordsByPatientRequest{
		PatientId: patientID,
	}
	return c.client.GetServiceRecordsByPatient(ctx, req)
}

func (c *BillingClient) LookupWardPrice(ctx context.Context, req *billingpb.LookupWardPriceRequest) (*billingpb.LookupWardPriceResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.priceClient.LookupWardPrice(ctx, req)
}

func (c *BillingClient) LookupLabTestPrice(ctx context.Context, req *billingpb.LookupLabTestPriceRequest) (*billingpb.LookupLabTestPriceResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.priceClient.LookupLabTestPrice(ctx, req)
}

func (c *BillingClient) CreateInvoice(ctx context.Context, req *billingpb.CreateInvoiceRequest) (*billingpb.CreateInvoiceResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.CreateInvoice(ctx, req)
}

func (c *BillingClient) MarkInvoiceAsPaid(ctx context.Context, req *billingpb.MarkInvoiceAsPaidRequest) (*billingpb.MarkInvoiceAsPaidResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.MarkInvoiceAsPaid(ctx, req)
}

func (c *BillingClient) LookupServicePrice(ctx context.Context, req *billingpb.LookupServicePriceRequest) (*billingpb.LookupServicePriceResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.priceClient.LookupServicePrice(ctx, req)
}

func (c *BillingClient) GenerateInvoiceFromRecords(ctx context.Context, req *billingpb.GenerateInvoiceFromRecordsRequest) (*billingpb.CreateInvoiceResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return c.client.GenerateInvoiceFromRecords(ctx, req)
}