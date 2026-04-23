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

type Provider string

const (
	ProviderOpenAI  Provider = "openai"
	ProviderGroq  Provider = "groq"
	ProviderOllama Provider = "ollama"
	ProviderFallback Provider = "fallback"
)

type OpenAIClient struct {
	apiKey       string
	httpClient   *http.Client
	provider    Provider
	ollamaURL   string
	model       string
	useFreeMode  bool
}

type MessageBody struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model       string        `json:"model"`
	Messages    []MessageBody `json:"messages"`
	MaxTokens   int           `json:"max_tokens"`
	Temperature float64       `json:"temperature,omitempty"`
}

type OpenAIResponse struct {
	ID      string `json:"id"`
	Object string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Message MessageBody `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	mode := strings.ToLower(strings.TrimSpace(os.Getenv("AI_MODE")))
	provider := strings.ToLower(strings.TrimSpace(os.Getenv("AI_PROVIDER")))
	ollamaURL := strings.TrimSpace(os.Getenv("OLLAMA_BASE_URL"))
	groqAPIKey := strings.TrimSpace(os.Getenv("GROQ_API_KEY"))

	if provider == "" {
		provider = "fallback"
	}

	var selectedProvider Provider
	var model string

	switch Provider(provider) {
	case ProviderGroq:
		selectedProvider = ProviderGroq
		model = "llama-3.1-8b-instant"
	case ProviderOllama:
		selectedProvider = ProviderOllama
		model = "llama3.2"
		if ollamaURL == "" {
			ollamaURL = "http://localhost:11434/v1"
		}
	case ProviderOpenAI:
		selectedProvider = ProviderOpenAI
		model = "gpt-3.5-turbo"
	default:
		if apiKey != "" && mode == "paid" {
			selectedProvider = ProviderOpenAI
			model = "gpt-3.5-turbo"
		} else if groqAPIKey != "" {
			selectedProvider = ProviderGroq
			model = "llama-3.1-8b-instant"
		} else {
			selectedProvider = ProviderFallback
			model = "fallback"
		}
	}

	useFreeMode := selectedProvider == ProviderFallback

	return &OpenAIClient{
		apiKey:      apiKey,
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		provider:    selectedProvider,
		ollamaURL:   ollamaURL,
		model:       model,
		useFreeMode: useFreeMode,
	}
}

func (c *OpenAIClient) Chat(ctx context.Context, messages []MessageBody) (string, error) {
	if c.useFreeMode {
		return c.freeChatResponse(messages)
	}

	var endpoint string
	var authHeader string

	switch c.provider {
	case ProviderGroq:
		endpoint = "https://api.groq.com/openai/v1/chat/completions"
		authHeader = fmt.Sprintf("Bearer %s", os.Getenv("GROQ_API_KEY"))
	case ProviderOllama:
		endpoint = c.ollamaURL + "/chat/completions"
		authHeader = ""
	case ProviderOpenAI:
		endpoint = "https://api.openai.com/v1/chat/completions"
		authHeader = fmt.Sprintf("Bearer %s", c.apiKey)
	default:
		return c.freeChatResponse(messages)
	}

	requestBody := OpenAIRequest{
		Model:       c.model,
		Messages:   messages,
		MaxTokens:   250,
		Temperature: 0.7,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return c.freeChatResponse(messages)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		fmt.Printf("[AI Client] Provider %s returned status %d: %s\n", c.provider, resp.StatusCode, string(data))
		return c.freeChatResponse(messages)
	}

	var result OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return c.freeChatResponse(messages)
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

	businessName := "nosso negócio"
	for _, msg := range messages {
		if msg.Role == "system" {
			content := msg.Content
			if idx := strings.Index(content, "para o negócio "); idx != -1 {
				endIdx := strings.Index(content[idx+16:], ".")
				if endIdx > 0 {
					businessName = content[idx+16 : idx+16+endIdx]
				}
			}
			break
		}
	}

	defaultResponses := map[string]string{
		"greeting":   fmt.Sprintf("Olá! É um prazer atendê-lo em %s. Como posso ajudá-lo hoje?", businessName),
		"scheduling": fmt.Sprintf("Ficarei feliz em ajudá-lo com agendamento em %s. Qual serviço você deseja e para qual dia?", businessName),
		"services":  fmt.Sprintf("Temos diversos serviços disponíveis em %s. Qual gostaria de conhecer?", businessName),
		"hours":     fmt.Sprintf("Posso verificar os horários disponíveis para você. Em qual dia gostaria de agendar?"),
		"default":  fmt.Sprintf("Estou aqui para ajudá-lo em %s. Como posso ser útil?", businessName),
		"default2": "Posso ajudar com informações sobre serviços, horários e agendamentos. O que gostaria de saber?",
	}

	if lastUserMessage == "" {
		return defaultResponses["greeting"], nil
	}

	lower := strings.ToLower(lastUserMessage)

	switch {
	case strings.Contains(lower, "oi") || strings.Contains(lower, "olá") || strings.Contains(lower, "ola") || strings.Contains(lower, "bom dia") || strings.Contains(lower, "boa tarde") || strings.Contains(lower, "boa noite"):
		return defaultResponses["greeting"], nil

	case strings.Contains(lower, "horário") || strings.Contains(lower, "horarios") || strings.Contains(lower, "agenda") || strings.Contains(lower, "disponível") || strings.Contains(lower, "disponivel") || strings.Contains(lower, "quando"):
		return defaultResponses["scheduling"], nil

	case strings.Contains(lower, "preço") || strings.Contains(lower, "valor") || strings.Contains(lower, "custo") || strings.Contains(lower, "quanto"):
		return fmt.Sprintf("Temos diversos serviços com preços variados. Qual serviço gostaria de saber o valor?"), nil

	case strings.Contains(lower, "serviço") || strings.Contains(lower, "servicos") || strings.Contains(lower, "corte") || strings.Contains(lower, "barbeiro") || strings.Contains(lower, "salão") || strings.Contains(lower, "tratamento"):
		return defaultResponses["services"], nil

	case strings.Contains(lower, "quitar") || strings.Contains(lower, "marcar") || strings.Contains(lower, "reservar") || strings.Contains(lower, "book"):
		return defaultResponses["scheduling"], nil

	case strings.Contains(lower, "endereço") || strings.Contains(lower, "endereco") || strings.Contains(lower, "onde") || strings.Contains(lower, "local"):
		return fmt.Sprintf("Estamos localizados e atendemos com alegria. Gostaria de mais informações sobre nossa localização?"), nil

	case strings.Contains(lower, "contato") || strings.Contains(lower, "telefone") || strings.Contains(lower, "whatsapp") || strings.Contains(lower, "email"):
		return fmt.Sprintf("Pode entrar em contato conosco para mais informações. Como prefiere ser atendido?"), nil

	case strings.Contains(lower, "obrigado") || strings.Contains(lower, "valeu") || strings.Contains(lower, "tchau") || strings.Contains(lower, "até mais"):
		return fmt.Sprintf("É um prazer atendê-lo! Volte sempre que precisar. Até breve!"), nil

	default:
		return defaultResponses["default2"], nil
	}
}

func (c *OpenAIClient) GetProvider() string {
	return string(c.provider)
}

func (c *OpenAIClient) GetModel() string {
	return c.model
}

func (c *OpenAIClient) IsUsingAI() bool {
	return !c.useFreeMode
}