package main

import (
	"errors"
	alert "item/alert/controllers"
	"item/common"
	master "item/master/controllers"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleLambdaEvent lambda entrypoint
func HandleLambdaEvent(request common.RequestCommon) (interface{}, error) {
	if request.Kind == "master" {
		return HandleItemMaster(request)
	}
	if request.Kind == "alert" {
		return HandleAlert(request.Method, request)
	}
	return nil, errors.New("not match kind:" + request.Kind)
}

// HandleItemMaster itemmaster
func HandleItemMaster(request interface{}) (master.Responce, error) {
	contoller := master.NewItemMasterController()
	req := request.(master.Request)
	res, err := contoller.GetItemMaster(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

// HandleAlert alert
func HandleAlert(method string, request interface{}) (interface{}, error) {
	contoller := alert.NewAlertController()
	if method == "get" {
		req := request.(alert.GetRequest)
		res, err := contoller.GetAlertLog(req)
		if err != nil {
			return res, err
		}
		return res, nil
	}
	if method == "put" {
		req := request.(alert.PutRequest)
		res, err := contoller.PutAlertLog(req)
		if err != nil {
			return res, err
		}
		return res, nil
	}
	return nil, errors.New("not match method:" + method)
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
