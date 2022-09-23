package s60

type CreditManualResponse struct {
	CommonResponse
	ApplyStatus string `json:"applyStatus"` // 申请状态 22-订单状态人工审核中
}
