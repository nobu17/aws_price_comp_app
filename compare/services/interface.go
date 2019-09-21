package services

// ServiceImpl service interface
type ServiceImpl interface {
	StartCompare(req InputModel) (OutputModel, error)
}

// InputModel input
type InputModel struct {
	// userid
	UserID string `json:"user_id"`
	// username
	UserName string `json:"user_name"`
	// mail
	Mail string `json:"mail"`
	// id
	GroupID string `json:"group_id"`
	// name
	GroupName string `json:"group_name"`
}

// NewInputModel constructor
func NewInputModel(userID, groupID, groupName string) InputModel {
	return InputModel{UserID: userID, GroupID: groupID, GroupName: groupName}
}

// OutputModel input
type OutputModel struct {
}
