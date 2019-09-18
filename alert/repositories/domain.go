package repositories

// AlertImpl impl
type AlertImpl interface {
	GetAlertLog(req GetRequest) (GetResponce, error)
	PutAlertLog(req PutRequest) (PutResponce, error)
}

// GetRequest リクエスト入力パラメータ
type GetRequest struct {
	// ユーザーID
	UserID string
	// 最小日付
	MinAlertDate string
}

// GetResponce result
type GetResponce struct {
	// 送付リスト
	SendAlertLogList []SendAlertLog
}

// PutRequest struct
type PutRequest struct {
	// 挿入データ
	PutAlertLogList []SendAlertLog
}

// PutResponce result
type PutResponce struct {
	// 書き込み数
	Wrote int
}

// SendAlertLog struct
type SendAlertLog struct {
	// ユーザーID
	UserID string
	// アラート日付
	AlertDate string
	// 店鋪種類
	StoreType string
	// 商品ID
	ProductID string
	// 価格
	Price int
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
