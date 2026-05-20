package mailer

import (
	"fmt"
	"os"

	"github.com/resend/resend-go/v2"
)

type Email struct {
	FromName     string
	FromEmail    string
	ToName       string
	ToEmail      string
	Subject      string
	PlainContent string
	HTMLContent  string
}

// SendUserEmail sends an email using Resend API.
func SendUserEmail(e Email) error {
	// Configure API key authorization
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("RESEND_API_KEY not set")
	}

	// Create API client
	client := resend.NewClient(apiKey)

	// Prepare the email
	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("%s <%s>", e.FromName, e.FromEmail),
		To:      []string{fmt.Sprintf("%s <%s>", e.ToName, e.ToEmail)},
		Subject: e.Subject,
		Html:    e.HTMLContent,
		Text:    e.PlainContent,
	}

	// Send the email
	_, err := client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}

// TestConnection tests the connection to Resend API (optional helper function)
func TestConnection() error {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("RESEND_API_KEY not set")
	}

	client := resend.NewClient(apiKey)

	_, err := client.Domains.List()
	if err != nil {
		return fmt.Errorf("error when calling Resend API: %v", err)
	}

	return nil
}
