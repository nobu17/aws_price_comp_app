package main

import (
	"errors"
	"notify/controllers"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleLambdaEvent lambda entrypoint
func HandleLambdaEvent(request controllers.Request) (interface{}, error) {
	contoller := controllers.NewNotifyController()
	if request.Method == "put" {
		res, err := contoller.SendNotify(request.PutRequest)
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
