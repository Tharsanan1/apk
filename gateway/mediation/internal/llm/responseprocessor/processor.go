package responseprocessor

// Processor is an interface that defines methods for processing responses from LLMs.
type Processor interface {
	// ProcessResponse processes the response and extracts token counts.
	ProcessResponse(response string) (*Processor, error)
	// GetTotalTokens returns the total number of tokens processed.
	GetTotalTokens() int
	// GetPromptTokens returns the number of prompt tokens processed.
	GetPromptTokens() int
	// GetCompletionTokens returns the number of completion tokens processed.
	GetCompletionTokens() int
}