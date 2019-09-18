package main

import (
	"errors"
	"user/controllers"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleLambdaEvent lambda entrypoint
func HandleLambdaEvent(request controllers.Request) (interface{}, error) {
	contoller := controllers.NewUserController()
	if request.Method == "get" {
		res, err := contoller.GetUserInfo(request.GetRequest)
		if err != nil {
			return res, err
		}
		return res, nil
	}
	return nil, errors.New("Not match method")
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
