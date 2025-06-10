package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	r := gin.Default()

	r.POST("/addTrack", AddTrack)
	r.GET("/getAllTracks", GetAllTracks)
	r.GET("/getRandomTrack", GetRandomTrack)

	ginLambda = ginadapter.NewV2(r)
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
