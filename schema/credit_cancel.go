package s60

type CreditCancelResponse struct {
	CommonResponse
	ApplyStatus  string `json:"applyStatus"`  // 申请状态 4 系统审核中
	CancelReason string `json:"cancelReason"` // 取消原因
}
