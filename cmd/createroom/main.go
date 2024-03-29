package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hogehoge-banana/sls-rtc-backend/pkg/api/createroom"
)

type proxyResponse events.APIGatewayProxyResponse

func handler(request events.APIGatewayWebsocketProxyRequest) (proxyResponse, error) {

	msg, err := createroom.CreateRoom(request)
	if err != nil {
		log.Println(msg)
		return proxyResponse{}, err
	}

	return proxyResponse{
		Body:       msg,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
