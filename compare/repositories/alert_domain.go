package repositories

// AlertImpl item master get impl
type AlertImpl interface {
	GetAlerts(req AlertGetRequest) (AlertGetResponce, error)
	PutAlerts(req AlertPutRequest) (AlertPutResponce, error)
}

// AlertGetRequest リクエスト入力パラメータ
type AlertGetRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// 最小日付
	MinAlertDate string `json:"min_alert_date"`
}

// AlertGetResponce result
type AlertGetResponce struct {
	// 送付リスト
	SendAlertList []SendAlertLog `json:"send_alert_list"`
}

// AlertPutRequest struct
type AlertPutRequest struct {
	// 挿入データ
	PutAlertList []SendAlertLog
}

// AlertPutResponce result
type AlertPutResponce struct {
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
