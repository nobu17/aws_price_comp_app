package services

// ServiceImpl service interface
type ServiceImpl interface {
	StartObserve(req InputModel) (OutputModel, error)
}

// InputModel input
type InputModel struct {
	// ユーザーID
	UserID string
	// パスワード
	Password string
}

// NewInputModel constructor
func NewInputModel(userID, password string) InputModel {
	return InputModel{UserID: userID, Password: password}
}

// OutputModel input
type OutputModel struct {
}
