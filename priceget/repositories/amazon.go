package repositories

import (
	"common/info"
	"common/log"
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// AmazonMakertRepository AmazonMarketPlace実装
type amazonMakertRepository struct {
	logger log.LoggerImpl
}

// NewAmazonMakertRepository constructor
func NewAmazonMakertRepository(logger log.LoggerImpl) GetProductPriceImpl {
	return &amazonMakertRepository{logger: logger}
}

// GetProductPrice impl
func (u *amazonMakertRepository) GetProductPrice(req Request) (ProductInfo, error) {
	if req.ProductID != "" {
		return u.getProductInfo(req.ProductID)
	}
	return NewProductInfo(req.ProductID), errors.New("no input productID")
}

func (u *amazonMakertRepository) getProductInfo(isbn string) (ProductInfo, error) {
	// Load the URL
	doc, err := goquery.NewDocument(info.GetAmazonPrdocutURL(isbn))
	if err != nil {
		u.logger.LogWrite(log.Error, "load page error:"+info.GetAmazonPrdocutURL(isbn))
		return NewProductInfo(isbn), err
	}
	// read html
	html, err := doc.Html()
	if err != nil {
		return NewProductInfo(isbn), err
	}
	// u.logger.LogWrite(log.Info, "html:"+html)
	// empty html is treated as error
	if html == "<html><head></head><body></body></html>" {
		u.logger.LogWrite(log.Error, "epmty page")
		return NewProductInfo(isbn), errors.New("no html data")
	}

	// get a price
	pRes := doc.Find("span.olpOfferPrice")
	if len(pRes.Nodes) < 1 {
		// sold out
		u.logger.LogWrite(log.Info, "span.olpOfferPrice is not existed")
		tRes := NewProductInfo(isbn)
		tRes.IsSoldOut = true
		return tRes, nil
	}
	priceText := pRes.First().Text()
	priceNum, err := u.getTrimmedPrice(priceText)
	if err != nil || priceNum <= 0 {
		u.logger.LogWrite(log.Error, "price string parse error:"+priceText)
		return NewProductInfo(isbn), errors.New("price string parse error:" + priceText)
	}

	// get a shiping price
	sRes := doc.Find("span.olpShippingPrice")
	if len(sRes.Nodes) < 1 {
		u.logger.LogWrite(log.Error, "no shipping price data")
		return NewProductInfo(isbn), errors.New("no shipping price data")
	}
	shippingPriceText := sRes.First().Text()
	shippingPriceNum, err := u.getTrimmedPrice(shippingPriceText)
	if err != nil {
		u.logger.LogWrite(log.Error, "shippingPrice string parse error")
		return NewProductInfo(isbn), errors.New("shippingPrice string parse error:" + shippingPriceText)
	}

	responce := NewProductInfo(isbn)
	responce.Price = priceNum
	responce.ShippingFee = shippingPriceNum
	responce.IsSoldOut = false

	return responce, nil
}

func (u *amazonMakertRepository) getTrimmedPrice(input string) (int, error) {
	str := strings.Replace(input, "￥", "", -1)
	str = strings.Replace(str, ",", "", -1)
	str = strings.TrimSpace(str)

	num, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	return num, nil
}
