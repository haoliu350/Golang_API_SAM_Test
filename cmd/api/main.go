package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/haozi/Go-Api-Sam-Test/internal/handlers"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing request data for request %s\n", request.RequestContext.RequestID)
	log.Printf("Method: %s, Resource: %s, Path: %s\n", request.HTTPMethod, request.Resource, request.Path)

	// Route the request to the appropriate handler
	switch {
	case request.HTTPMethod == "GET" && request.Resource == "/hello":
		return handlers.GetHandler(ctx, request)
	case request.HTTPMethod == "POST" && request.Resource == "/submit":
		return handlers.PostHandler(ctx, request)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}
}

func main() {
	lambda.Start(handleRequest)
}