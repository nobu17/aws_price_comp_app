module user

go 1.13

require (
	common v0.0.0
	github.com/aws/aws-lambda-go v1.13.2
	github.com/aws/aws-sdk-go v1.34.0
	github.com/guregu/dynamo v1.4.0
)

replace common => ../common
