package controllers

// Request リクエスト入力パラメータ
type Request struct {
	// ユーザーID
	UserID string `json:"user_id"`
	// パスワード
	Password string `json:"password"`
}

// Responce リクエスト出力パラメータ
type Responce struct {
}