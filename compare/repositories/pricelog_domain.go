package repositories

// PriceLogImpl pricelog master get impl
type PriceLogImpl interface {
	UpdatePriceLog(req PutPriceLogRequest) (PutPriceLogResponce, error)
}

// PutPriceLogRequest param
type PutPriceLogRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
	// モード(put or update)
	Mode string `json:"mode"`
	// PriceLog List
	PriceLogList []PriceLogForPut `json:"price_log_list"`
}

// PutPriceLogResponce result
type PutPriceLogResponce struct {
}

// PriceLogForPut price last get
type PriceLogForPut struct {
	// store type
	StoreType string `json:"store_type"`
	// item id
	ItemID string `json:"item_id"`
	// price
	Price int `json:"price"`
}
