package dto

// MessageContext represents the context of a message in the system.
type MessageContext struct {
	// Path is the path of the request.	
	Path string `json:"path"`
	// VHost is the virtual host of the request.
	VHost string `json:"vHost"`
	// BasePath is the base path of the request.
	BasePath string `json:"basePath"`
	// Method is the HTTP method of the request.
	Method string `json:"method"`
	// APIVersion is the version of the API.
	APIVersion string `json:"apiVersion"`
	// APIName is the name of the API.
	APIName string `json:"apiName"`
	// ClusterName is the name of the cluster.
	ClusterName string `json:"clusterName"`
	// EnableBackendBasedAIRatelimit indicates if backend-based AI rate limiting is enabled.
	EnableBackendBasedAIRatelimit string `json:"enableBackendBasedAIRatelimit"`
	// SuspendAIModel indicates if the AI model is suspended.
	SuspendAIModel string `json:"suspendAIModel"`
	// BackendBasedAIRatelimitDescriptorValue is the descriptor value for backend-based AI rate limiting.
	BackendBasedAIRatelimitDescriptorValue string `json:"backendBasedAIRatelimitDescriptorValue"`
	// RequestMethod is the HTTP method of the request.
	RequestMethod string `json:"requestMethod"`
	// Organization is the organization associated with the request.
	Organization string `json:"organization"`
	// ApplicationID is the ID of the application making the request.
	ApplicationID string `json:"applicationId"`
	// CorrelationID is the correlation ID of the request.
	CorrelationID string `json:"correlationId"`
	// EndpointBasepath is the base path of the endpoint.
	EndpointBasepath string `json:"endpointBasepath"`
	// AISchema is the schema for AI processing.
	AISchema string `json:"aiSchema"`
	// RequestHeaders is the headers of the request.
	RequestHeaders map[string]string `json:"requestHeaders"`
	// RequestBody is the body of the request.
	RequestBody string `json:"requestBody"`
	// ResponseHeaders is the headers of the response.
	ResponseHeaders map[string]string `json:"responseHeaders"`
	// ResponseBody is the body of the response.
	ResponseBody string `json:"responseBody"`
	// ResponseStatusCode is the status code of the response.
	ResponseStatusCode int `json:"responseStatusCode"`
	// ResponseStatusMessage is the status message of the response.
	ResponseStatusMessage string `json:"responseStatusMessage"`
	// ResponseTime is the time taken to process the response.
	ResponseTime int64 `json:"responseTime"`
	// ErrorMessage is the error message if any error occurred.
	ErrorMessage string `json:"errorMessage"`
	// ErrorCode is the error code if any error occurred.
	ErrorCode int `json:"errorCode"`
	// ResponseEncoding is the encoding of the response.
	ResponseEncoding string `json:"responseEncoding"`

}
