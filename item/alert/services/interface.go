package services

// ServiceImpl service interface
type ServiceImpl interface {
	GetAlertLog(req GetInputModel) (GetOutputModel, error)
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

// SendAlertLog struct
type SendAlertLog struct {
	// ユーザーID
	UserID string
	// アラート日付
	AlertDate string
	// 商品ID
	ProductID string
	// 価格
	Price int
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
