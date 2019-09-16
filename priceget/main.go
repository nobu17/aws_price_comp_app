package main

import (
	"priceget/controllers"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleLambdaEvent lambda entrypoint
func HandleLambdaEvent(request controllers.Request) (controllers.Responce, error) {
	contoller := controllers.NewPriceGetController()
	res, err := contoller.GetProductPriceList(request)
	if err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
	//var prods = controllers.Request
	//contoller := controllers.NewPriceGetController()
	//req := controllers.Request{}
	//list := make([]controllers.ProductRequest, 0)
	//list = append(list, controllers.ProductRequest{StoreType: "amazon", ProductID: "4773623071000"})
	//list = append(list, controllers.ProductRequest{StoreType: "amazon", ProductID: "475111297X"})
	//list = append(list, controllers.ProductRequest{StoreType: "amazon", ProductID: "4838728980"})

	//req.ProductList = list
	//res, err := contoller.GetProductPriceList(req)
}
