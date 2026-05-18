// pkg/client/patient_client.go

package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	patientpb "github.com/homasy/pkg/shared/patient-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// PatientClient is a client for the patient service
type PatientClient struct {
	client     patientpb.PatientServiceClient
	conn       *grpc.ClientConn
	serverAddr string
	mu         sync.Mutex
	connected  bool
}

// NewPatientClient creates a new patient client
func NewPatientClient(serverAddr string) *PatientClient {
	return &PatientClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the patient service
func (c *PatientClient) Connect() error {
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
		
		log.Printf("Failed to connect to patient service (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2 // Exponential backoff
	}
	
	if err != nil {
		return fmt.Errorf("failed to connect to patient service after %d attempts: %v", maxRetries, err)
	}
	
	c.conn = conn
	c.client = patientpb.NewPatientServiceClient(conn)
	c.connected = true
	
	log.Printf("Connected to patient service at %s", c.serverAddr)
	return nil
}

// Close closes the connection to the patient service
func (c *PatientClient) Close() error {
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

// GetPatient gets a patient by ID
func (c *PatientClient) GetPatient(ctx context.Context, patientID string) (*patientpb.GetPatientResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	req := &patientpb.GetPatientRequest{
		Id: patientID,
	}

	return c.client.GetPatient(ctx, req)
}

// CheckPatientExists checks if a patient exists
func (c *PatientClient) CheckPatientExists(ctx context.Context, patientID string) (bool, error) {
	patient, err := c.GetPatient(ctx, patientID)
	if err != nil {
		return false, err
	}

	return patient != nil, nil
}