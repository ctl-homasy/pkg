package mailer

import (
	"context"
	"fmt"
	"os"

	brevo "github.com/getbrevo/brevo-go/lib"
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

// SendEmail sends an email using Brevo API.
func SendUserEmail(e Email) error {
	// Initialize context and configuration
	ctx := context.Background()
	cfg := brevo.NewConfiguration()

	// Configure API key authorization
	apiKey := os.Getenv("SENDINBLUE_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("BREVO_API_KEY not set")
	}

	cfg.AddDefaultHeader("api-key", apiKey)
	cfg.AddDefaultHeader("partner-key", apiKey)

	// Create API client
	client := brevo.NewAPIClient(cfg)

	// Prepare the email
	email := brevo.SendSmtpEmail{
		Sender: &brevo.SendSmtpEmailSender{
			Name:  e.FromName,
			Email: e.FromEmail,
		},
		To: []brevo.SendSmtpEmailTo{
			{
				Email: e.ToEmail,
				Name:  e.ToName,
			},
		},
		Subject:     e.Subject,
		HtmlContent: e.HTMLContent,
		TextContent: e.PlainContent,
	}

	// Send the email
	_, httpResp, err := client.TransactionalEmailsApi.SendTransacEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	if httpResp.StatusCode >= 400 {
		return fmt.Errorf("email sending returned status: %d", httpResp.StatusCode)
	}

	return nil
}

// TestConnection tests the connection to Brevo API (optional helper function)
func TestConnection() error {
	ctx := context.Background()
	cfg := brevo.NewConfiguration()

	apiKey := os.Getenv("SENDINBLUE_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("BREVO_API_KEY not set")
	}

	cfg.AddDefaultHeader("api-key", apiKey)
	cfg.AddDefaultHeader("partner-key", apiKey)

	client := brevo.NewAPIClient(cfg)

	_, _, err := client.AccountApi.GetAccount(ctx)
	if err != nil {
		return fmt.Errorf("error when calling AccountApi->get_account: %v", err)
	}

	return nil
}