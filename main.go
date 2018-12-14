package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	map1 := make(map[string]string)
	for key,value := range request.Headers {
		map1[key]=value
		fmt.Println(key+":"+value)
	}
	str, _ := json.Marshal(map1)
	body := bytes.NewBuffer([]byte(str))
	res,err := http.Post("http://120.79.83.107:8094/push", "application/json;charset=utf-8", body)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello AWS Lambda and Netlify",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
