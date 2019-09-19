package repositories

import (
	"common/info"
	"common/log"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// deafult dynamoregion
const envRegionKey = "REGION"
const defaultRegion = "us-east-1"

// ses const
const envSendFromKey = "SEND_FROM"
const defaultSendFrom = "no-replay@fake.com"
const subject = "監視通知"

// sesRepository interface
type sesRepository struct {
	logger log.LoggerImpl
}

// NewSesRepository constructor.
func NewSesRepository(logger log.LoggerImpl) SendNotifyImpl {
	return &sesRepository{logger: logger}
}

func (u *sesRepository) SendNotify(req PutRequest) (PutResponce, error) {
	ses, err := u.getSES()
	if err != nil {
		return PutResponce{}, err
	}
	// Create a mail body
	mailBody := u.getMailBody(req.GroupID, req.ProductInfoList)
	input := u.getEmailInput(req.UserInfo.Mail, mailBody)
	_, err = ses.SendEmail(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "ses SendMail Error"+fmt.Sprint(err))
	}
	return PutResponce{}, err
}

func (u *sesRepository) getSES() (*ses.SES, error) {
	region := os.Getenv(envRegionKey)
	if region == "" {
		region = defaultRegion
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}

	// Create an SES session.
	svc := ses.New(sess)

	return svc, nil
}

func (u *sesRepository) getEmailInput(sendMail string, body string) *ses.SendEmailInput {
	sendFrom := os.Getenv(envSendFromKey)
	if sendFrom == "" {
		sendFrom = defaultSendFrom
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(sendMail),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("utf-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sendFrom),
	}
	return input
}

func (u *sesRepository) getMailBody(groupID string, infoList []ProductInfo) string {
	str := "<h1>指し値よりも安い商品があります。</h1></br>"
	str += fmt.Sprintf("<h2>グループ名:%v</2></br></br>", groupID)

	for _, item := range infoList {
		str += fmt.Sprintf("<p>商品名:%v</p></br>", item.Name)
		str += fmt.Sprintf("<p>価格(送料):%v (%v)</p></br>", item.Price, item.ShippingFee)

		if info.IsStoreTypeAmazon(item.StoreType) {
			str += fmt.Sprintf("<a href='%v'>Link</a></br></br>", info.GetAmazonPrdocutURL(item.ProductID))
		}
		if info.IsStoreTypeSurugaya(item.StoreType) {
			str += fmt.Sprintf("<a href='%v'>Link</a></br></br>", info.GetSurugayaPrdocutURL(item.ProductID))
		}
	}

	return str
}
