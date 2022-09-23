package s60

type CreditRejectResponse struct {
	CommonResponse
	ApplyStatus  string `json:"applyStatus"`  // 申请状态 6 授信拒绝
	RefuseReason string `json:"refuseReason"` // 拒绝原因
}
