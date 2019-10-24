package repositories

// PriceLogImpl interface of user repository
type PriceLogImpl interface {
	GetPriceLogs(req GetRequest) (GetResponce, error)
	PutPriceLogs(req PutRequest) (PutResponce, error)
	DeletePriceLogs(req DeleteRequest) (DeleteResponce, error)
}

// GetRequest param
type GetRequest struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
}

// GetResponce result
type GetResponce struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
	// PriceLog List
	PriceLogList []PriceLog
}

// PriceLog price last get
type PriceLog struct {
	StoreType            string
	ItemID               string
	Price                int
	LastModifiedDatetime string
}

// PutRequest param
type PutRequest struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
	// PriceLog List
	PriceLogList []PriceLog
}

// PutResponce result
type PutResponce struct {
}

// DeleteRequest param
type DeleteRequest struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
}

// DeleteResponce result
type DeleteResponce struct {
}
