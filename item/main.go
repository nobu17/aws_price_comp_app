package main

import (
	"errors"
	"item/controllers"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleItemMaster itemmaster
func HandleItemMaster(request controllers.Request) (interface{}, error) {
	contoller := controllers.NewItemMasterController()
	if request.Method == "get" {
		res, err := contoller.GetItemMaster(request.GetRequest)
		if err != nil {
			return res, err
		}
		return res, nil
	}
	return nil, errors.New("Not match method")
}

func main() {
	lambda.Start(HandleItemMaster)
}
