module priceget

go 1.13

require (
	common v0.0.0
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/aws/aws-lambda-go v1.13.2
	github.com/aws/aws-sdk-go v1.24.1 // indirect
	github.com/djimenez/iconv-go v0.0.0-20160305225143-8960e66bd3da
	golang.org/x/text v0.3.2
)

replace common => ../common
