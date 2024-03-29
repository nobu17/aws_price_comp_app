package controllers

// Request request common
type Request struct {
	// Method type
	Method string `json:"method"`
	// GetParam
	GetRequest GetRequest `json:"get_param"`
	// PutRequest
	PutRequest PutRequest `json:"put_param"`
}

// GetRequest input
type GetRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// 最小日付
	MinAlertDate string `json:"min_alert_date"`
}

// GetResponce output
type GetResponce struct {
	// 送付リスト
	SendAlertLogList []SendAlertLog `json:"send_alert_list"`
}

// PutRequest struct
type PutRequest struct {
	// 挿入データ
	PutAlertLogList []SendAlertLog `json:"put_alert_list"`
}

// PutResponce result
type PutResponce struct {
}

// SendAlertLog struct
type SendAlertLog struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// アラート日付
	AlertDate string `json:"alert_date"`
	// 店鋪種類
	StoreType string `json:"store_type"`
	// 商品ID
	ProductID string `json:"product_id"`
	// 価格
	Price int `json:"price"`
}

// NewSendAlertLog construcotr
func NewSendAlertLog(userID string, alertDate string, storeType string, productID string, price int) SendAlertLog {
	return SendAlertLog{
		UserID:    userID,
		AlertDate: alertDate,
		StoreType: storeType,
		ProductID: productID,
		Price:     price,
	}
}
