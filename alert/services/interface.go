package services

// ServiceImpl service interface
type ServiceImpl interface {
	GetAlertLog(req GetInputModel) (GetOutputModel, error)
	PutAlertLog(req PutInputModel) error
}

// GetInputModel input
type GetInputModel struct {
	// ユーザーID
	UserID string
	// 最小日付
	MinAlertDate string
}

// GetOutputModel output
type GetOutputModel struct {
	// ItemMasters 商品リスト
	// 送付リスト
	SendAlertLogList []SendAlertLog
}

// PutInputModel struct
type PutInputModel struct {
	// 挿入データ
	PutAlertLogList []SendAlertLog
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
