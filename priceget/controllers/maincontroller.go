package controllers

import (
	"common/log"
	"fmt"
	"priceget/services"
)

// PriceGetController strcut
type PriceGetController struct {
	logger  log.LoggerImpl
	service services.ServiceImpl
}

// NewPriceGetController constructor
func NewPriceGetController() PriceGetController {
	container := Container{}

	controller := PriceGetController{}
	controller.logger = container.NewLogger()
	controller.service = container.NewService()

	return controller
}

// GetProductPriceList productlist
func (u *PriceGetController) GetProductPriceList(req Request) (Responce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start GetProductPriceList", req)

	data := getConvertInput(req)
	res, err := u.service.GetProductInfo(data)
	if err != nil {
		msg := fmt.Sprintf("end GetProductPriceList, error:%v", err)
		u.logger.LogWrite(log.Info, msg)
		return Responce{}, err
	}

	output := getConvertOutput(res)
	u.logger.LogWriteWithMsgAndObj(log.Info, "end GetProductPriceList", output)

	return output, nil
}

func getConvertInput(req Request) services.InputProductModel {
	prods := make([]services.ProductRequest, 0)
	for _, prod := range req.ProductList {
		prods = append(prods, services.ProductRequest{StoreType: prod.StoreType, ProductID: prod.ProductID})
	}
	return services.InputProductModel{ProductList: prods}
}

func getConvertOutput(output services.OutputProductModel) Responce {
	prods := make([]ProductInfo, 0)
	for _, prod := range output.ProductInfoList {
		prods = append(prods, ProductInfo{ProductID: prod.ProductID, StoreType: prod.StoreType, Price: prod.Price, ShippingFee: prod.ShippingFee, IsSoldOut: prod.IsSoldOut})
	}
	fprods := make([]ProductInfo, 0)
	for _, prod := range output.FailProductInfoList {
		fprods = append(fprods, ProductInfo{ProductID: prod.ProductID, StoreType: prod.StoreType, Price: prod.Price, ShippingFee: prod.ShippingFee, IsSoldOut: prod.IsSoldOut})
	}
	return Responce{ProductInfoList: prods, FailProductInfoList: fprods}
}
