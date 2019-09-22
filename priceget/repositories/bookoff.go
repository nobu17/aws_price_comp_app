package repositories

import (
	"bufio"
	"common/info"
	"common/log"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// bookoffRepository BookoffRepository
type bookoffRepository struct {
	logger log.LoggerImpl
}

// NewBookoffRepository constructor
func NewBookoffRepository(logger log.LoggerImpl) GetProductPriceImpl {
	return &bookoffRepository{logger: logger}
}

// GetProductPrice impl
func (u *bookoffRepository) GetProductPrice(req Request) (ProductInfo, error) {
	if req.ProductID != "" {
		return u.getProductInfo(req.ProductID)
	}
	u.logger.LogWrite(log.Error, "empty productID")
	return NewProductInfo(req.ProductID), errors.New("no input productID")
}

func (u *bookoffRepository) getProductInfo(productID string) (ProductInfo, error) {
	var itemURL = info.GetBookoffPrdocutURL(productID)
	u.logger.LogWrite(log.Info, "start:"+itemURL)
	res, err := http.Get(itemURL)
	if err != nil {
		u.logger.LogWrite(log.Error, "load page error:"+itemURL)
		return NewProductInfo(productID), err
	}
	u.logger.LogWrite(log.Info, "end:"+itemURL)
	defer res.Body.Close()

	// Convert the designated charset HTML to utf-8 encoded HTML.
	utfBody := transform.NewReader(bufio.NewReader(res.Body), japanese.ShiftJIS.NewDecoder())
	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(utfBody)

	// get a price
	pRes := doc.Find("td.oldprice")
	if len(pRes.Nodes) < 1 {
		// sold out
		u.logger.LogWrite(log.Warn, "not found td.oldprice")
		tRes := NewProductInfo(productID)
		tRes.IsSoldOut = true
		return tRes, nil
	}
	// get price text
	var pText = pRes.First().Text()
	u.logger.LogWrite(log.Info, "price text:"+pText)
	var currentPrice = -1
	currentPrice, err = u.getTrimmedPrice(pText)

	if err != nil {
		u.logger.LogWrite(log.Error, "getTrimmedPrice error:"+pText)
		return NewProductInfo(productID), err
	}

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

func (u *bookoffRepository) getTrimmedPrice(input string) (int, error) {
	str := strings.Split(input, "（")[0]
	str = strings.Replace(str, "￥", "", -1)
	str = strings.Replace(str, ",", "", -1)
	str = strings.TrimSpace(str)

	num, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	return num, nil
}
