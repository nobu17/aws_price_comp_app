package aws

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// CallLambdaWithSync call lambda func with sync.
func CallLambdaWithSync(funcName string, region string, payload interface{}) (*lambda.InvokeOutput, error) {
	return callLambda("RequestResponse", funcName, region, payload)
}

// CallLambdaWithASync call lambda func with async.
func CallLambdaWithASync(funcName string, region string, payload interface{}) (*lambda.InvokeOutput, error) {
	return callLambda("Event", funcName, region, payload)
}

func callLambda(invokeType string, funcName string, region string, payload interface{}) (*lambda.InvokeOutput, error) {
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v,%v,%v", funcName, region, invokeType)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	input := &lambda.InvokeInput{FunctionName: aws.String(funcName), Payload: jsonBytes, InvocationType: aws.String(invokeType)}
	client := lambda.New(sess, &aws.Config{Region: aws.String(region)})
	result, err := client.Invoke(input)

	if result != nil && result.FunctionError != nil {
		if *result.FunctionError == "" {
			return result, errors.New("function error:" + *result.FunctionError)
		}
	}

	return result, err
}
