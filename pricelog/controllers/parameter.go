package controllers

// Request request common
type Request struct {
	// Method type
	Method string `json:"method"`
	// GetParam
	GetRequest GetRequest `json:"get_param"`
	// PutParam
	PutRequest PutRequest `json:"put_param"`
	// DeleteParam
	DeleteRequest DeleteRequest `json:"delete_param"`
}

// GetRequest リクエスト入力パラメータ
type GetRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
}

// NewGetRequest constructor
func NewGetRequest(userID string, groupID string) GetRequest {
	return GetRequest{UserID: userID, GroupID: groupID}
}

// GetResponce result
type GetResponce struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
	// PriceLog List
	PriceLogList []PriceLogForGet `json:"price_log_list"`
}

// PriceLogForGet price last get
type PriceLogForGet struct {
	// store type
	StoreType string `json:"store_type"`
	// item id
	ItemID string `json:"item_id"`
	// price
	Price int `json:"price"`
	// last modified date time
	LastModifiedDatetime string `json:"last_modified_datetime"`
}

// PutRequest param
type PutRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
	// モード(put or update)
	Mode string `json:"mode"`
	// PriceLog List
	PriceLogList []PriceLogForPut `json:"price_log_list"`
}

// PutResponce result
type PutResponce struct {
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

// DeleteRequest リクエスト入力パラメータ
type DeleteRequest struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// グループID
	GroupID string `json:"group_id"`
}

// DeleteResponce result
type DeleteResponce struct {
}
