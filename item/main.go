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
			return nil, err
		}
		return res, nil
	}
	if request.Method == "put" {
		res, err := contoller.PutItemMaster(request.PutRequest)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	if request.Method == "delete" {
		res, err := contoller.DeleteItemMaster(request.DeleteRequest)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, errors.New("Not match method")
}

func main() {
	lambda.Start(HandleItemMaster)
}
