package controllers

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

// SendAlertLog struct
type SendAlertLog struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// アラート日付
	AlertDate string `json:"alert_date"`
	// 商品ID
	ProductID string `json:"product_id"`
	// 価格
	Price int `json:"price"`
}

// NewSendAlertLog construcotr
func NewSendAlertLog(userID string, alertDate string, productID string, price int) SendAlertLog {
	return SendAlertLog{
		UserID:    userID,
		AlertDate: alertDate,
		ProductID: productID,
		Price:     price,
	}
}
