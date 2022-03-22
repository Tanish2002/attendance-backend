package handlers

import (
	"context"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-lambda-go/events"
)

var Router *gin.Engine
var ProxyResponse events.APIGatewayProxyResponse
var ProxyRequest events.APIGatewayProxyRequest

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var ginLambda *ginadapter.GinLambda
	if ginLambda == nil {
		ginLambda = ginadapter.New(Router)
	}
	ProxyResponse, _ = ginLambda.Proxy(req)
	ProxyRequest = req
	return ginLambda.ProxyWithContext(ctx, req)
}
