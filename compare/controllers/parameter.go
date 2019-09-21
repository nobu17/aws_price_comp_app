package controllers

// Request リクエスト入力パラメータ
type Request struct {
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

// Responce リクエスト出力パラメータ
type Responce struct {
}