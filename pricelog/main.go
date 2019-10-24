package main

import (
	"errors"
	"pricelog/controllers"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleLambdaEvent lambda entrypoint
func HandleLambdaEvent(request controllers.Request) (interface{}, error) {
	contoller := controllers.NewPriceLogController()
	if request.Method == "get" {
		res, err := contoller.GetPriceLog(request.GetRequest)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	if request.Method == "put" {
		res, err := contoller.PutPriceLog(request.PutRequest)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	if request.Method == "delete" {
		res, err := contoller.DeletePriceLog(request.DeleteRequest)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, errors.New("Not match method")
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
