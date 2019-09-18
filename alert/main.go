package main

import (
	"alert/controllers"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleLambdaEvent lambda entrypoint
func HandleLambdaEvent(request controllers.Request) (interface{}, error) {
	contoller := controllers.NewAlertController()
	if request.Method == "get" {
		res, err := contoller.GetAlertLog(request.GetRequest)
		if err != nil {
			return res, err
		}
		return res, nil
	}
	if request.Method == "put" {
		res, err := contoller.PutAlertLog(request.PutRequest)
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
