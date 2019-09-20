module startup

go 1.13

require (
	common v0.0.0
	github.com/aws/aws-lambda-go v1.13.2
)

replace common => ../common
