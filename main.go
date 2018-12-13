package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	for key,value := range request.Headers {
		fmt.Println(key+":"+value)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello AWS Lambda and Netlify",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
