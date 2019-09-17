package common

// RequestCommon common reqparam
type RequestCommon struct {
	// Kind
	Kind string `json:"kind"`
	// メソッド
	Method string `json:"method"`
}
