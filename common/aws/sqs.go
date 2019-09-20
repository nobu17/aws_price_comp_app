package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SendMessageToSQS sendmessage specifiedURL SQS
func SendMessageToSQS(queURL, messageBody, region string) (messageID string, err error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess, &aws.Config{Region: aws.String(region)})
	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(messageBody),
		QueueUrl:     aws.String(queURL),
		DelaySeconds: aws.Int64(1),
	}

	sqsRes, err := svc.SendMessage(params)
	if err != nil {
		return "", err
	}
	messageID = *sqsRes.MessageId
	return messageID, err
}
