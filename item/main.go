package main

import (
	"item/master/controllers"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleLambdaEvent lambda entrypoint
func HandleLambdaEvent(request controllers.Request) (controllers.Responce, error) {
	contoller := controllers.NewItemMasterController()
	res, err := contoller.GetItemMaster(request)
	if err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
