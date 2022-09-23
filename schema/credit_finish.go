package s60

type CreditFinishResponse struct {
	CommonResponse
	ApplyStatus string `json:"applyStatus"` // 申请状态 4 系统审核中
}
