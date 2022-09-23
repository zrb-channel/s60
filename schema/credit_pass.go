package s60

import "github.com/shopspring/decimal"

type CreditPassResponse struct {
	CommonResponse
	ApplyStatus     string          `json:"applyStatus"`     // 申请状态 5 授信通过
	CreditAmount    decimal.Decimal `json:"creditAmount"`    // 授信金额 分
	CreditRate      string          `json:"creditRate"`      // 授信利率
	CreditTime      string          `json:"creditTime"`      // 授信时间
	CreditStartTime string          `json:"creditStartTime"` // 授信有效期起始时间
	CreditEndTime   string          `json:"creditEndTime"`   // 授信有效期结束时间
}
