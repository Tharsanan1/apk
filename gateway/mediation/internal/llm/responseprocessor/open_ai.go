package responseprocessor

import (
	"encoding/json"
	"io"

	"github.com/wso2/apk/gateway/mediation/internal/llm/apischema/openai"
)

// OpenAIResponseProcessor is a struct that processes OpenAI responses and extracts token counts.
type OpenAIResponseProcessor struct {
	// TotalTokens is the total number of tokens processed.
	TotalTokens       int
	// PromptTokens is the number of prompt tokens processed.
	PromptTokens      int
	// CompletionTokens is the number of completion tokens processed.
	CompletionTokens  int
	// ProcessedResponse is the processed response string.
	ProcessedResponse string
}

// NewOpenAIResponseProcessor creates a new instance of OpenAIResponseProcessor.
func NewOpenAIResponseProcessor() *OpenAIResponseProcessor {
	return &OpenAIResponseProcessor{
		TotalTokens:       0,
		PromptTokens:      0,
		CompletionTokens:  0,
		ProcessedResponse: "",
	}
}

// ProcessResponse processes the OpenAI response and extracts token counts.
func (p *OpenAIResponseProcessor) ProcessResponse(response io.Reader) (*OpenAIResponseProcessor, error) {
	// Simulate processing the response and extracting token counts
	// buf, err := io.ReadAll(response)
	// if err != nil {
	// 	return p, fmt.Errorf("failed to read response: %w", err)
	// }
	// p.ProcessedResponse = string(buf)

	var resp openai.ChatCompletionResponse
	if err := json.NewDecoder(response).Decode(&resp); err != nil {
		return p, err
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	p.ProcessedResponse = string(bytes)

	p.TotalTokens = resp.Usage.TotalTokens
	p.PromptTokens = resp.Usage.PromptTokens
	p.CompletionTokens = resp.Usage.CompletionTokens
	return p, nil
}

// GetTotalTokens returns the total number of tokens processed.
func (p *OpenAIResponseProcessor) GetTotalTokens() int {
	return p.TotalTokens
}

// GetPromptTokens returns the number of prompt tokens processed.
func (p *OpenAIResponseProcessor) GetPromptTokens() int {
	return p.PromptTokens
}

// GetCompletionTokens returns the number of completion tokens processed.
func (p *OpenAIResponseProcessor) GetCompletionTokens() int {
	return p.CompletionTokens
}

// GetProcessedResponse returns the processed response.
func (p *OpenAIResponseProcessor) GetProcessedResponse() string {
	return p.ProcessedResponse
}