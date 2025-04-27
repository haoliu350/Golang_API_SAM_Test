package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// Response is a simple response structure
type Response struct {
	Message string `json:"message"`
}

// Request is a simple request structure for POST
type Request struct {
	Name string `json:"name"`
}

// GetHandler handles GET requests
func GetHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing GET request: %s", request.RequestContext.RequestID)

	// Get name from query parameters, default to "World"
	name := request.QueryStringParameters["name"]
	if name == "" {
		name = "World"
	}

	// Create response
	response := Response{
		Message: "Hello, " + name + "!",
	}

	// Convert to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error creating response",
		}, nil
	}

	// Return response
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonResponse),
	}, nil
}

// PostHandler handles POST requests
func PostHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing POST request: %s", request.RequestContext.RequestID)

	// Parse request body
	var req Request
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	// Create response
	response := Response{
		Message: "Hello, " + req.Name + "! Your submission was received.",
	}

	// Convert to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error creating response",
		}, nil
	}

	// Return response
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonResponse),
	}, nil
}