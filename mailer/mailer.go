package mailer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

// mailtrapSendURL is the Mailtrap transactional sending endpoint.
const mailtrapSendURL = "https://send.api.mailtrap.io/api/send"

// SendUserEmail sends an email using the Mailtrap sending API.
func SendUserEmail(e Email) error {
	apiToken := os.Getenv("MAILTRAP_API_TOKEN")
	if apiToken == "" {
		return fmt.Errorf("MAILTRAP_API_TOKEN not set")
	}

	// Build the request payload in the shape Mailtrap expects.
	payload := map[string]interface{}{
		"from": map[string]string{
			"email": e.FromEmail,
			"name":  e.FromName,
		},
		"to": []map[string]string{
			{"email": e.ToEmail},
		},
		"subject": e.Subject,
	}
	if e.PlainContent != "" {
		payload["text"] = e.PlainContent
	}
	if e.HTMLContent != "" {
		payload["html"] = e.HTMLContent
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error encoding email payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, mailtrapSendURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("error sending email: status %d: %s", res.StatusCode, string(respBody))
	}

	return nil
}

// TestConnection tests the connection to the Mailtrap API (optional helper function)
func TestConnection() error {
	apiToken := os.Getenv("MAILTRAP_API_TOKEN")
	if apiToken == "" {
		return fmt.Errorf("MAILTRAP_API_TOKEN not set")
	}
	return nil
}
