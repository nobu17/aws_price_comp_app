package lambda

import (
	"encoding/json"

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

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String(region)})
	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String(funcName), Payload: jsonBytes, InvocationType: aws.String(invokeType)})
	return result, err
}
