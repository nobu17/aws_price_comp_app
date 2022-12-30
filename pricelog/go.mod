module pricelog

go 1.13

require (
	common v0.0.0
	github.com/aws/aws-lambda-go v1.13.2
	github.com/aws/aws-sdk-go v1.33.0
	github.com/guregu/dynamo v1.4.1
)

replace common => ../common
