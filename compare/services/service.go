package services

import (
	"common/log"
	"common/util"
	"compare/repositories"
	"fmt"
)

// day threthold get alert
const dayThrethold = 7

// compareService service interface
type compareService struct {
	logger             log.LoggerImpl
	alertRepository    repositories.AlertImpl
	itemRepository     repositories.ItemImpl
	priceRepository    repositories.PriceImpl
	notifyRepository   repositories.NotifyImpl
	pricelogRepository repositories.PriceLogImpl
}

// NewCompareService constructor
func NewCompareService(logger log.LoggerImpl,
	alertRepository repositories.AlertImpl,
	itemRepository repositories.ItemImpl,
	priceRepository repositories.PriceImpl,
	notifyRepository repositories.NotifyImpl,
	pricelogRepository repositories.PriceLogImpl,
) ServiceImpl {
	return &compareService{logger: logger,
		alertRepository:    alertRepository,
		itemRepository:     itemRepository,
		priceRepository:    priceRepository,
		notifyRepository:   notifyRepository,
		pricelogRepository: pricelogRepository}
}

func (u *compareService) StartCompare(req InputModel) (OutputModel, error) {
	u.logger.LogWrite(log.Info, "start StartCompare")
	// get observ item list
	items, prices, err := u.getObservItemAndPriceList(req)
	if err != nil {
		u.logger.LogWrite(log.Error, "getObservItemAndPriceList is failed:"+fmt.Sprint(err))
		return OutputModel{}, err
	}

	// compare threthold
	var priceList = make([]repositories.ProductInfo, 0)
	var notifyTargets = make([]repositories.NotifyProductInfo, 0)
	for _, price := range prices.ProductInfoList {
		if !price.IsSoldOut {
			var prod = getProduct(price.ProductID, price.StoreType, items.ItemMasters)
			if prod == nil {
				u.logger.LogWrite(log.Warn, "not match product:"+price.ProductID)
				continue
			}
			// store for log
			priceList = append(priceList, price)

			// check threthoild
			if prod.ThretholdPrice >= (price.Price + price.ShippingFee) {
				target := repositories.NewNotifyProductInfo(prod.ProductID, prod.StoreType, prod.ItemName, price.Price, price.ShippingFee)
				notifyTargets = append(notifyTargets, target)
				continue
			}
			u.logger.LogWrite(log.Info, fmt.Sprintf("not over threthould: ID:%v current:%v, threthod:%v, item:%v", price.ProductID, (price.Price+price.ShippingFee), prod.ThretholdPrice, prod.ItemName))
		}
	}
	// store log
	err = u.updatePriceLog(req, priceList)
	if err != nil {
		u.logger.LogWrite(log.Error, "updatePriceLog is failed:"+fmt.Sprint(err))
		return OutputModel{}, err
	}

	if len(notifyTargets) <= 0 {
		u.logger.LogWrite(log.Info, "no notify products:")
		return OutputModel{}, nil
	}

	u.logger.LogWriteWithMsgAndObj(log.Info, "over threthold prods:", notifyTargets)
	// filter past alert items
	sendTargets, err := u.filterPastSentItems(req.UserID, notifyTargets)
	if err != nil {
		u.logger.LogWrite(log.Error, "filterPastSentItems is failed:"+fmt.Sprint(err))
		return OutputModel{}, err
	}
	if len(sendTargets) <= 0 {
		u.logger.LogWrite(log.Info, "no notify product which is not alerted")
		return OutputModel{}, nil
	}
	// send notify
	err = u.sendNotifyAndPutAlertLog(req, sendTargets)
	if err != nil {
		u.logger.LogWrite(log.Error, "sendNotifyAndPutAlertLog is failed:"+fmt.Sprint(err))
		return OutputModel{}, err
	}
	u.logger.LogWrite(log.Info, "end StartCompare")
	return OutputModel{}, nil
}

func (u *compareService) updatePriceLog(req InputModel, prices []repositories.ProductInfo) error {
	var list = make([]repositories.PriceLogForPut, 0)
	for _, price := range prices {
		allPrice := price.Price + price.ShippingFee
		pLog := repositories.PriceLogForPut{ItemID: price.ProductID, StoreType: price.StoreType, Price: allPrice}
		list = append(list, pLog)
	}
	param := repositories.PutPriceLogRequest{UserID: req.UserID, GroupID: req.GroupID, Mode: "update", PriceLogList: list}
	_, err := u.pricelogRepository.UpdatePriceLog(param)
	return err
}

// get observe item and prices
func (u *compareService) getObservItemAndPriceList(req InputModel) (*repositories.ItemGetResponce, *repositories.PriceGetResponce, error) {
	// get observ item list
	var itemReq = repositories.ItemGetRequest{UserID: req.UserID, GroupID: req.GroupID}
	items, err := u.itemRepository.GetItems(itemReq)
	if err != nil {
		u.logger.LogWrite(log.Error, "get items is failed:"+fmt.Sprint(err))
		return nil, nil, err
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "get items result:", items)

	// get price list
	var itemList = make([]repositories.ProductRequest, 0)
	for _, item := range items.ItemMasters {
		itemList = append(itemList, repositories.ProductRequest{StoreType: item.StoreType, ProductID: item.ProductID})
	}
	// get price list
	var priceReq = repositories.PriceGetRequest{ProductList: itemList}
	prices, err := u.priceRepository.GetPrices(priceReq)
	if err != nil {
		u.logger.LogWrite(log.Error, "get items is failed:"+fmt.Sprint(err))
		return nil, nil, err
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "get prices result:", prices)

	return &items, &prices, nil
}

// send notify and put alerts
func (u *compareService) sendNotifyAndPutAlertLog(input InputModel, notifies []repositories.NotifyProductInfo) error {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start sendNotifyAndPutAlertLog:input", input)
	u.logger.LogWriteWithMsgAndObj(log.Info, "start sendNotifyAndPutAlertLog:notifies", notifies)

	var user = repositories.NewNotifyUserInfo(input.UserID, input.UserName, input.Mail)
	var request = repositories.NotifyPutRequest{UserInfo: user, GroupID: input.GroupID, ProductInfoList: notifies}
	u.logger.LogWriteWithMsgAndObj(log.Info, "start SendNotify", request)
	_, err := u.notifyRepository.SendNotify(request)
	if err != nil {
		u.logger.LogWrite(log.Error, "SendNotify is failed:"+fmt.Sprint(err))
		return err
	}
	u.logger.LogWrite(log.Info, "end SendNotify")
	u.logger.LogWrite(log.Info, "start PutAlerts")
	// put alert log if successed
	var sendlist = make([]repositories.SendAlertLog, 0)
	for _, item := range notifies {
		sendlist = append(sendlist, repositories.NewSendAlertLog(input.UserID, util.GetJSTTimeStr(0), item.StoreType, item.ProductID, (item.Price+item.ShippingFee)))
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "start PutAlerts", sendlist)
	_, err = u.alertRepository.PutAlerts(repositories.AlertPutRequest{PutAlertList: sendlist})
	if err != nil {
		u.logger.LogWrite(log.Error, "PutAlerts is failed:"+fmt.Sprint(err))
		return err
	}
	u.logger.LogWrite(log.Info, "end PutAlerts")
	u.logger.LogWrite(log.Info, "end sendNotifyAndPutAlertLog")
	return nil
}

// filter past sent alert item which is expensive
func (u *compareService) filterPastSentItems(userID string, req []repositories.NotifyProductInfo) ([]repositories.NotifyProductInfo, error) {
	// get past 7 days alerts
	nowJST := util.GetJSTTimeStr(dayThrethold)

	var input = repositories.AlertGetRequest{UserID: userID, MinAlertDate: nowJST}
	u.logger.LogWriteWithMsgAndObj(log.Info, "get alert:", input)
	res, err := u.alertRepository.GetAlerts(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "get alert is failed:"+fmt.Sprint(err))
		return nil, err
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "get alert list:", res)

	var targets = make([]repositories.NotifyProductInfo, 0)
	for _, notifyItem := range req {
		// if alert item is match
		var alertedItem = getAlertProduct(notifyItem.ProductID, notifyItem.StoreType, res.SendAlertList)
		if alertedItem != nil {
			// if more cheper than before alert it should alert
			if (notifyItem.Price + notifyItem.ShippingFee) > alertedItem.Price {
				u.logger.LogWriteWithMsgAndObj(log.Info, "not get over threthold:", notifyItem)
				continue
			}
		}
		u.logger.LogWriteWithMsgAndObj(log.Info, "get over threthold:", notifyItem)
		targets = append(targets, notifyItem)
	}
	u.logger.LogWriteWithMsgAndObj(log.Info, "get over threthold list:", targets)
	return targets, nil
}

func getAlertProduct(ProductID string, storeType string, masters []repositories.SendAlertLog) *repositories.SendAlertLog {
	for _, item := range masters {
		if (item.ProductID == ProductID) && (item.StoreType == storeType) {
			return &item
		}
	}
	return nil
}

func getProduct(prdouctID string, storeType string, masters []repositories.ItemMaster) *repositories.ItemMaster {
	for _, item := range masters {
		if (item.ProductID == prdouctID) && (item.StoreType == storeType) {
			return &item
		}
	}
	return nil
}
