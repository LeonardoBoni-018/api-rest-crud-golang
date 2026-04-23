package ai

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

type OpenAIClient struct {
	apiKey      string
	httpClient  *http.Client
	useFreeMode bool
}

type OpenAIRequest struct {
	Model     string        `json:"model"`
	Messages  []MessageBody `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

type MessageBody struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Message MessageBody `json:"message"`
	} `json:"choices"`
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	mode := strings.ToLower(strings.TrimSpace(os.Getenv("AI_MODE")))
	return &OpenAIClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		useFreeMode: apiKey == "" || mode != "paid",
	}
}

func (c *OpenAIClient) Chat(ctx context.Context, messages []MessageBody) (string, error) {
	if c.useFreeMode {
		return c.freeChatResponse(messages)
	}

	requestBody := OpenAIRequest{
		Model:     "gpt-3.5-turbo",
		Messages:  messages,
		MaxTokens: 250,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("openai status=%d body=%s", resp.StatusCode, string(data))
	}

	var result OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("openai returned no choices")
	}

	return result.Choices[0].Message.Content, nil
}

func (c *OpenAIClient) freeChatResponse(messages []MessageBody) (string, error) {
	lastUserMessage := ""
	for i := len(messages) - 1; i >= 0; i-- {
		if messages[i].Role == "user" {
			lastUserMessage = strings.TrimSpace(messages[i].Content)
			break
		}
	}

	if lastUserMessage == "" {
		return "Olá! Estou em modo gratuito de IA e posso ajudar com serviços, horários e agendamentos. Como posso ajudar você hoje?", nil
	}

	lower := strings.ToLower(lastUserMessage)

	switch {
	case strings.Contains(lower, "horário") || strings.Contains(lower, "horarios") || strings.Contains(lower, "agendamento") || strings.Contains(lower, "disponível"):
		return "Estou em modo gratuito e posso ajudar com informações sobre horários e agendamentos. Qual serviço você deseja agendar e para qual dia?", nil
	case strings.Contains(lower, "preço") || strings.Contains(lower, "valor") || strings.Contains(lower, "custo"):
		return "Estou em modo gratuito e posso fornecer valores aproximados. Qual serviço você quer saber o preço?", nil
	case strings.Contains(lower, "serviço") || strings.Contains(lower, "corte") || strings.Contains(lower, "barbeiro") || strings.Contains(lower, "salão"):
		return "No modo gratuito, posso responder sobre serviços disponíveis e suas características. Qual serviço você gostaria de conhecer?", nil
	case strings.Contains(lower, "oi") || strings.Contains(lower, "olá") || strings.Contains(lower, "ola"):
		return "Olá! Estou usando o modo gratuito de IA. Posso ajudar com serviços, horários e agendamentos para o seu negócio.", nil
	default:
		return "Estou rodando em modo gratuito. Posso ajudar com serviços, horários e agendamentos. Como posso ajudar hoje?", nil
	}
}
