package s60

type CreditCreateResponse struct {
	CommonResponse
	ApplyStatus string `json:"applyStatus"` // 申请状态 20-订单状态完善信息中
}
