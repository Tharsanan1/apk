/*
 *  Copyright (c) 2025, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package extproc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"time"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_service_proc_v3 "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
	"github.com/wso2/apk/gateway/mediation/internal/config"
	"github.com/wso2/apk/gateway/mediation/internal/dto"
	"github.com/wso2/apk/gateway/mediation/internal/llm/responseprocessor"
	"github.com/wso2/apk/gateway/mediation/internal/logging"
	"github.com/wso2/apk/gateway/mediation/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	// "google.golang.org/grpc/health"
	// "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	pathAttribute                                   string = "path"
	vHostAttribute                                  string = "vHost"
	basePathAttribute                               string = "basePath"
	methodAttribute                                 string = "method"
	apiVersionAttribute                             string = "version"
	apiNameAttribute                                string = "name"
	clusterNameAttribute                            string = "clusterName"
	enableBackendBasedAIRatelimitAttribute          string = "enableBackendBasedAIRatelimit"
	backendBasedAIRatelimitDescriptorValueAttribute string = "backendBasedAIRatelimitDescriptorValue"
	customOrgMetadataKey                            string = "customorg"
	endpointBasepath                                string = "endpointBasepath"
	aiSchema                                        string = "aischema"
	suspendAIModelValueAttribute                    string = "ai:suspendmodel"
	externalProessingMetadataContextKey             string = "envoy.filters.http.ext_proc"
	subscriptionMetadataKey                         string = "ratelimit:subscription"
	usagePolicyMetadataKey                          string = "ratelimit:usage-policy"
	organizationMetadataKey                         string = "ratelimit:organization"
	orgAndRLPolicyMetadataKey                       string = "ratelimit:organization-and-rlpolicy"
	extractTokenFromMetadataKey                     string = "extracttokenfrom"
	promptTokenIDMetadataKey                        string = "prompttokenid"
	completionTokenIDMetadataKey                    string = "completiontokenid"
	totalTokenIDMetadataKey                         string = "totaltokenid"
	matchedAPIMetadataKey                           string = "request:matchedapi"
	matchedResourceMetadataKey                      string = "request:matchedresource"
	matchedSubscriptionMetadataKey                  string = "request:matchedsubscription"
	matchedApplicationMetadataKey                   string = "request:matchedapplication"

	modelMetadataKey string = "aitoken:model"
	metadataNamespace = "com.wso2.bijira.ai_gateway"
)

// ExternalProcessingServer represents a server for handling external processing requests.
// It contains a logger for logging purposes.
type ExternalProcessingServer struct {
	log logging.Logger
}

// StartExternalProcessingServer initializes and starts the external processing server.
// It creates a gRPC server using the provided configuration and registers the external
// processor server with it.
//
// Parameters:
//   - cfg: A pointer to the Server configuration which includes paths to the mediation's
//     public and private keys, and a logger instance.
//
// If there is an error during the creation of the gRPC server, the function will panic.
func StartExternalProcessingServer(cfg *config.Server) {
	kaParams := keepalive.ServerParameters{
		Time:    time.Duration(cfg.ExternalProcessingKeepAliveTime) * time.Hour, // Ping the client if it is idle for 2 hours
		Timeout: 20 * time.Second,
	}
	server, err := util.CreateGRPCServer(cfg.MediationServerPublicKeyPath,
		cfg.MediationServerPrivateKeyPath,
		grpc.MaxRecvMsgSize(cfg.ExternalProcessingMaxMessageSize),
		grpc.MaxHeaderListSize(uint32(cfg.ExternalProcessingMaxHeaderLimit)),
		grpc.KeepaliveParams(kaParams))
	if err != nil {
		panic(err)
	}

	// grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	cfg.Logger.Info("Health check added.....")
	envoy_service_proc_v3.RegisterExternalProcessorServer(server,
		&ExternalProcessingServer{cfg.Logger,})
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.ExternalProcessingPort))
	if err != nil {
		cfg.Logger.Error(err, fmt.Sprintf("Failed to listen on port: %s", cfg.ExternalProcessingPort))
	}
	cfg.Logger.Info(fmt.Sprintf("Starting to serve external processing server on port: %s", cfg.ExternalProcessingPort))
	if err := server.Serve(listener); err != nil {
		cfg.Logger.Error(err, "Failed to serve grpc server")
	}
}

// Process handles the external processing server stream. It continuously receives
// requests from the stream, processes them, and sends back appropriate responses.
// The function supports different types of processing requests including request headers,
// response headers, request body, and response body.
//
// Parameters:
// - srv: The stream server for processing external requests.
//
// Returns:
// - error: Returns an error if the context is done or if there is an issue receiving or sending the stream request.
//
// The function processes the following request types:
// - envoy_service_proc_v3.ProcessingRequest_RequestHeaders: Logs and processes request headers.
// - envoy_service_proc_v3.ProcessingRequest_ResponseHeaders: Logs and processes response headers.
// - envoy_service_proc_v3.ProcessingRequest_RequestBody: Logs and processes request body.
// - envoy_service_proc_v3.ProcessingRequest_ResponseBody: Logs and processes response body.
//
// If an unknown request type is received, it logs the unknown request type.
func (s *ExternalProcessingServer) Process(srv envoy_service_proc_v3.ExternalProcessor_ProcessServer) error {
	s.log.Sugar().Error("**************")
	ctx := srv.Context()
	messageContext := &dto.MessageContext{}
	for {
		s.log.Sugar().Error("**************1")
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		s.log.Sugar().Error("**************2")
		req, err := srv.Recv()
		s.log.Sugar().Error("**************3")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			s.log.Sugar().Error(err)
			return status.Errorf(codes.Unknown, "cannot receive stream request: %v", err)
		}

		resp := &envoy_service_proc_v3.ProcessingResponse{}
		// log req.Attributes
		s.log.Sugar().Debug(fmt.Sprintf("Attributes: %+v", req.Attributes))
		s.processgXDSRouteMetadataAttributes(req.Attributes, messageContext)
		switch v := req.Request.(type) {
		case *envoy_service_proc_v3.ProcessingRequest_RequestHeaders:
			rhq := &envoy_service_proc_v3.HeadersResponse{
				Response: &envoy_service_proc_v3.CommonResponse{
					HeaderMutation: &envoy_service_proc_v3.HeaderMutation{
					},
					// This is necessary if the remote server modified headers that are used to calculate the route.
					ClearRouteCache: true,
				},
			}
			resp.Response = &envoy_service_proc_v3.ProcessingResponse_RequestHeaders{
				RequestHeaders: rhq,
			}
			s.log.Sugar().Debug("Request Header Flow")
			s.processRequestHeader(messageContext, req.GetRequestHeaders().Headers)
		case *envoy_service_proc_v3.ProcessingRequest_RequestBody:
			resp.Response = &envoy_service_proc_v3.ProcessingResponse_RequestBody{
				RequestBody: &envoy_service_proc_v3.BodyResponse{
					Response: &envoy_service_proc_v3.CommonResponse{},
				},
			}
			s.log.Sugar().Debug("Request Body Flow")
			s.processRequestBody(messageContext, req.GetRequestBody())
		case *envoy_service_proc_v3.ProcessingRequest_ResponseHeaders:
			rhq := &envoy_service_proc_v3.HeadersResponse{
				Response: &envoy_service_proc_v3.CommonResponse{},
			}
			resp = &envoy_service_proc_v3.ProcessingResponse{
				Response: &envoy_service_proc_v3.ProcessingResponse_ResponseHeaders{
					ResponseHeaders: rhq,
				},
			}
			s.log.Sugar().Debug("Response Headers Flow")
			s.processResponseHeader(messageContext, req.GetResponseHeaders().Headers)
		case *envoy_service_proc_v3.ProcessingRequest_ResponseBody:
			resp.Response = &envoy_service_proc_v3.ProcessingResponse_ResponseBody{
				ResponseBody: &envoy_service_proc_v3.BodyResponse{
					Response: &envoy_service_proc_v3.CommonResponse{},
				},
			}
			s.log.Sugar().Debug("Response Body Flow")
			resp, err = s.processResponseBody(messageContext, req.GetResponseBody())
			if err != nil {
				s.log.Sugar().Error(err)
			}
		default:
			s.log.Sugar().Debug(fmt.Sprintf("Unknown Request type %v\n", v))
		}
		if err := srv.Send(resp); err != nil {
			s.log.Sugar().Debug(fmt.Sprintf("send error %v", err))
		}
	}
}

// processgXDSRouteMetadataAttributes extracts the external processing attributes from the given data.
func (s *ExternalProcessingServer) processgXDSRouteMetadataAttributes(data map[string]*structpb.Struct, messageContext *dto.MessageContext) (error) {

	// Get the fields from the map
	extProcData, exists := data["envoy.filters.http.ext_proc"]
	if !exists {
		return fmt.Errorf("key envoy.filters.http.ext_proc not found")
	}

	// Extract the "fields" and iterate over them
	fields := extProcData.Fields

	if field, ok := fields["request.method"]; ok {
		method := field.GetStringValue()
		messageContext.RequestMethod = method}

	// We need to navigate through the nested fields to get the actual values
	if field, ok := fields["xds.route_metadata"]; ok {

		filterMetadata := field.GetStringValue()
		var structData corev3.Metadata
		err := prototext.Unmarshal([]byte(filterMetadata), &structData)
		if err != nil {
			return fmt.Errorf("failed to parse Protobuf text: %v", err)
		}

		// Extract values for predefined keys
		extractedValues := make(map[string]string)

		keysToExtract := []string{
			aiSchema,
			pathAttribute,
			vHostAttribute,
			basePathAttribute,
			methodAttribute,
			apiVersionAttribute,
			apiNameAttribute,
			clusterNameAttribute,
			enableBackendBasedAIRatelimitAttribute,
			backendBasedAIRatelimitDescriptorValueAttribute,
			suspendAIModelValueAttribute,
			endpointBasepath,
		}

		if fieldEnvoyGateway, exists := structData.FilterMetadata["envoy-gateway"]; exists {
			if fieldEnvoyGatewayResources, exists := fieldEnvoyGateway.Fields["resources"]; exists {
				// `resources` is expected to be a list
				if resourcesList := fieldEnvoyGatewayResources.GetListValue(); resourcesList != nil {
					for _, resource := range resourcesList.Values {
						resourceStruct := resource.GetStructValue()
						if resourceStruct == nil {
							continue
						}
						// Check if annotations exist
						if annotationsField, exists := resourceStruct.Fields["annotations"]; exists {
							if annotationsStruct := annotationsField.GetStructValue(); annotationsStruct != nil {
								
								for _, key := range keysToExtract {
									s.log.Sugar().Debugf("Key: %s", key)
									if field, exists := annotationsStruct.Fields[key]; exists {
										s.log.Sugar().Debugf("Field: %s", field)
										extractedValues[key] = field.GetStringValue()
										switch key {
										case pathAttribute:
											messageContext.Path = extractedValues[key]
										case vHostAttribute:
											messageContext.VHost = extractedValues[key]
										case basePathAttribute:
											messageContext.BasePath = extractedValues[key]
										case methodAttribute:
											messageContext.Method = extractedValues[key]
										case apiVersionAttribute:
											messageContext.APIVersion = extractedValues[key]
										case apiNameAttribute:
											messageContext.APIName = extractedValues[key]
										case clusterNameAttribute:
											messageContext.ClusterName = extractedValues[key]
										case enableBackendBasedAIRatelimitAttribute:
											messageContext.EnableBackendBasedAIRatelimit = extractedValues[key]
										case backendBasedAIRatelimitDescriptorValueAttribute:
											messageContext.BackendBasedAIRatelimitDescriptorValue = extractedValues[key]
										case suspendAIModelValueAttribute:
											messageContext.SuspendAIModel = extractedValues[key]
										case endpointBasepath:
											messageContext.EndpointBasepath = extractedValues[key]
										case aiSchema:
											messageContext.AISchema = extractedValues[key]
										}
									}
								}
								
							}
						}
					}
				}
			}
		}
		s.log.Sugar().Debugf("Message Context: %+v", messageContext)
		// Return the populated struct
		return nil
	}

	// Key not found
	return fmt.Errorf("key xds.route_metadata not found")
}

// processRequestHeader processes the request headers from the given message context.
func (s *ExternalProcessingServer) processRequestHeader(messageContxt *dto.MessageContext, headers *corev3.HeaderMap) {
	// Process the request headers
	s.log.Sugar().Debug("Processing Request Headers")
	messageContxt.RequestHeaders = make(map[string]string)
	for _, header := range headers.GetHeaders() {
		key := header.GetKey()
		value := string(header.GetRawValue())
		messageContxt.RequestHeaders[key] = value
		s.log.Sugar().Debug(fmt.Sprintf("Header: %s, Value: %s", key, value))
	}

	// Add your processing logic here
}

// processResponseHeader processes the response headers from the given message context.
func (s *ExternalProcessingServer) processResponseHeader(messageContxt *dto.MessageContext, headers *corev3.HeaderMap) {
	// Process the request headers
	s.log.Sugar().Debug("Processing Request Headers")
	messageContxt.ResponseHeaders = make(map[string]string)
	for _, header := range headers.GetHeaders() {
		key := header.GetKey()
		value := string(header.GetRawValue())
		if key == "Content-Encoding" {
			messageContxt.ResponseEncoding = value
		}
		messageContxt.ResponseHeaders[key] = value
		s.log.Sugar().Debug(fmt.Sprintf("Header: %s, Value: %s", key, value))
	}

	// Add your processing logic here
}

// processRequestBody processes the request body from the given message context.
func (s *ExternalProcessingServer) processRequestBody(messageContxt *dto.MessageContext, body *envoy_service_proc_v3.HttpBody) {
	// Process the request headers
	s.log.Sugar().Debug("Processing Request Headers")
	messageContxt.RequestBody = string(body.Body)
	s.log.Sugar().Debug(fmt.Sprintf("Request Body: %s", messageContxt.RequestBody))

	// Add your processing logic here
}

// processResponseBody processes the response body from the given message context.
func (s *ExternalProcessingServer) processResponseBody(messageContxt *dto.MessageContext, body *envoy_service_proc_v3.HttpBody) (res *envoy_service_proc_v3.ProcessingResponse, err error) {
	// Process the request headers
	s.log.Sugar().Debug("Processing Request Headers")
	messageContxt.ResponseBody = string(body.Body)
	s.log.Sugar().Debug(fmt.Sprintf("Response Body: %s", messageContxt.ResponseBody))
	resp := &envoy_service_proc_v3.ProcessingResponse{
		Response: &envoy_service_proc_v3.ProcessingResponse_ResponseBody{
			ResponseBody: &envoy_service_proc_v3.BodyResponse{
				Response: &envoy_service_proc_v3.CommonResponse{
				},
			},
		},
	}
	// Check whether we need to do AI processing
	if messageContxt.AISchema == "" {
		s.log.Sugar().Debug("AI processing is not required")
		return resp, nil
	}
	metadata := make(map[string]*structpb.Value)
	if messageContxt.AISchema == "openai" {
		var br io.Reader
		var err error
		switch messageContxt.ResponseEncoding {
		case "gzip":
			br, err = gzip.NewReader(bytes.NewReader(body.Body))
			if err != nil {
				return nil, fmt.Errorf("failed to decode gzip: %w", err)
			}
		default:
			br = bytes.NewReader(body.Body)
		}
		p, err := responseprocessor.NewOpenAIResponseProcessor().ProcessResponse(br)
		if err != nil {
			return nil, fmt.Errorf("failed to process response: %w", err)
		}
		totalTokens := p.TotalTokens
		promptTokens := p.PromptTokens
		completionTokens := p.CompletionTokens
		
		metadata[promptTokenIDMetadataKey] = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(promptTokens)}}
		metadata[completionTokenIDMetadataKey] = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(completionTokens)}}
		metadata[totalTokenIDMetadataKey] = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(totalTokens)}}
	}
	resp.DynamicMetadata = &structpb.Struct{
		Fields: map[string]*structpb.Value{
			metadataNamespace: {
				Kind: &structpb.Value_StructValue{
					StructValue: &structpb.Struct{Fields: metadata},
				},
			},
		},
	}
	s.log.Sugar().Debug(fmt.Sprintf("Prepared metadata: %s", metadata))
	return resp, nil
}


