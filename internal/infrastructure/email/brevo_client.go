package email

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type EmailProvider interface {
	SendEmail(to, toName, subject, htmlContent string) error
}

type BrevoClient struct {
	apiKey     string
	fromEmail string
	fromName  string
	httpClient *http.Client
}

type BrevoSendEmailRequest struct {
	Sender   BrevoSender   `json:"sender"`
	To      []BrevoRecipient `json:"to"`
	Subject string       `json:"subject"`
	HtmlContent string    `json:"htmlContent"`
}

type BrevoSender struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type BrevoRecipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type BrevoResponse struct {
	MessageID string `json:"messageId"`
}

func NewBrevoClient() *BrevoClient {
	return &BrevoClient{
		apiKey:     strings.TrimSpace(os.Getenv("BREVO_API_KEY")),
		fromEmail:  strings.TrimSpace(os.Getenv("BREVO_FROM_EMAIL")),
		fromName:  strings.TrimSpace(os.Getenv("BREVO_FROM_NAME")),
		httpClient: &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *BrevoClient) SendEmail(to, toName, subject, htmlContent string) error {
	if c.apiKey == "" {
		fmt.Println("[Brevo] API key not configured, skipping email send")
		return fmt.Errorf("brevo API key not configured")
	}

	if c.fromEmail == "" {
		c.fromEmail = "noreply@agendamentos.com"
	}
	if c.fromName == "" {
		c.fromName = "Sistema de Agendamentos"
	}

	request := BrevoSendEmailRequest{
		Sender: BrevoSender{
			Email: c.fromEmail,
			Name:  c.fromName,
		},
		To: []BrevoRecipient{
			{
				Email: to,
				Name:  toName,
			},
		},
		Subject:     subject,
		HtmlContent: htmlContent,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"https://api.brevo.com/v3/smtp/email",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", c.apiKey)
	req.Header.Set("content-type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("brevo API error: status=%d body=%s", resp.StatusCode, string(respBody))
	}

	var result BrevoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("[Brevo] Email sent but failed to parse response: %v\n", err)
	}

	fmt.Printf("[Brevo] Email sent successfully to %s, messageId: %s\n", to, result.MessageID)
	return nil
}

func (c *BrevoClient) IsConfigured() bool {
	return c.apiKey != ""
}

func GetEmailProvider() EmailProvider {
	provider := strings.ToLower(strings.TrimSpace(os.Getenv("EMAIL_PROVIDER")))
	
	switch provider {
	case "brevo":
		return NewBrevoClient()
	default:
		brevoClient := NewBrevoClient()
		if brevoClient.IsConfigured() {
			return brevoClient
		}
		fmt.Println("[EmailProvider] No provider configured, using mock")
		return &MockEmailProvider{}
	}
}

type MockEmailProvider struct{}

func (m *MockEmailProvider) SendEmail(to, toName, subject, htmlContent string) error {
	fmt.Printf("[MOCK EMAIL] To: %s <%s>, Subject: %s\n", toName, to, subject)
	return nil
}