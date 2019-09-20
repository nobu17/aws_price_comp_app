package main

import (
	"startup/controllers"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleLambdaEvent lambda entrypoint
func HandleLambdaEvent(request controllers.Request) (interface{}, error) {
	contoller := controllers.NewStartupController()
	res, err := contoller.StartObserv(request)
	return res, err
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
