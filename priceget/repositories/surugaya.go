package repositories

import (
	"common/log"
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// surugayaRepository SurugayaRepository
type surugayaRepository struct {
	logger log.LoggerImpl
}

// NewSurugayaRepository constructor
func NewSurugayaRepository(logger log.LoggerImpl) GetProductPriceImpl {
	return &surugayaRepository{logger: logger}
}

const surugayaURLBase = "https://www.suruga-ya.jp/product/detail/"

// GetProductPrice impl
func (u *surugayaRepository) GetProductPrice(req Request) (ProductInfo, error) {
	if req.ProductID != "" {
		return u.getProductInfo(req.ProductID)
	}
	u.logger.LogWrite(log.Error, "empty productID")
	return NewProductInfo(req.ProductID), errors.New("no input productID")
}

func (u *surugayaRepository) getProductInfo(productID string) (ProductInfo, error) {
	doc, err := goquery.NewDocument(surugayaURLBase + productID)
	if err != nil {
		u.logger.LogWrite(log.Error, "load page error:"+surugayaURLBase+productID)
		return NewProductInfo(productID), err
	}
	// get a price
	pRes := doc.Find("span.mgnL10")
	if len(pRes.Nodes) < 1 {
		// sold out
		u.logger.LogWrite(log.Warn, "lnot found span.mgnL10")
		tRes := NewProductInfo(productID)
		tRes.IsSoldOut = true
		return tRes, nil
	}
	// セールの場合、複数取得される
	var currentPrice = -1
	pRes.Each(func(index int, s *goquery.Selection) {
		tempPrice, err := u.getTrimmedPrice(s.Text())
		if err == nil && tempPrice > 0 {
			if currentPrice == -1 || currentPrice > tempPrice {
				currentPrice = tempPrice
			}
		}
	})
	if currentPrice <= 0 {
		u.logger.LogWriteWithMsgAndObj(log.Error, "price get failed", pRes)
		return NewProductInfo(productID), errors.New("price get failed")
	}
	tRes := NewProductInfo(productID)
	tRes.Price = currentPrice
	tRes.ShippingFee = 0
	tRes.IsSoldOut = false
	return tRes, nil
}

func (u *surugayaRepository) getTrimmedPrice(input string) (int, error) {
	str := strings.Split(input, "円")[0]
	str = strings.Replace(str, ",", "", -1)
	str = strings.TrimSpace(str)

	num, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	return num, nil
}
