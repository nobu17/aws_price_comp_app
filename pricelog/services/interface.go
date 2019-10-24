package services

// ServiceImpl service interface
type ServiceImpl interface {
	GetPriceLog(req GetInputModel) (GetOutputModel, error)
	PutPriceLog(req PutInputModel) (PutOutputModel, error)
	UpdatePriceLog(req UpdateInputModel) (UpdateOutputModel, error)
	DeletePriceLog(req DeleteInputModel) (DeleteOutputModel, error)
}

// GetInputModel input
type GetInputModel struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
}

// NewGetInputModel constructor
func NewGetInputModel(userID string, groupID string) GetInputModel {
	return GetInputModel{UserID: userID, GroupID: groupID}
}

// GetOutputModel output
type GetOutputModel struct {
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

// PutInputModel input
type PutInputModel struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
	// PriceLog List
	PriceLogList []PriceLog
}

// PutOutputModel input
type PutOutputModel struct {
}

// UpdateInputModel input
type UpdateInputModel struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
	// PriceLog List
	PriceLogList []PriceLog
}

// UpdateOutputModel input
type UpdateOutputModel struct {
}

// DeleteInputModel input
type DeleteInputModel struct {
	// ユーザーID
	UserID string
	// グループID
	GroupID string
}

// DeleteOutputModel input
type DeleteOutputModel struct {
}
