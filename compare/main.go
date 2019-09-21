package main

import (
	"compare/controllers"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest handle event.
func HandleRequest(ctx context.Context, evt events.SQSEvent) (interface{}, error) {
	for _, item := range evt.Records {
		var input = controllers.Request{}
		if err := json.Unmarshal([]byte(item.Body), &input); err != nil {
			fmt.Printf("***Unmarshal error*** %#v\n", err)
			return nil, err
		}
		var controller = controllers.NewCompareController()
		_, err := controller.StartCompare(input)
		if err != nil {
			fmt.Printf("***StartCompare error*** %#v\n", err)
			return nil, err
		}
	}
	return controllers.Responce{}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
