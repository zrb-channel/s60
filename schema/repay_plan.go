package s60

import "github.com/shopspring/decimal"

type RepayPlanResponse struct {
	CommonResponse
	LoanNo          string             `json:"loanNo"`          // 放款编号
	RemainingAmount decimal.Decimal    `json:"remainingAmount"` // 剩余额度（分）
	NotifyRepayPlan []*NotifyRepayPlan `json:"notifyRepayPlan"` // 还款计划数组
}

type NotifyRepayPlan struct {
	TermNum     string          `json:"termNum"`     // 期数编号
	TermDate    string          `json:"termDate"`    // 本期应还时间 yyyy-MM-dd HH:mm:ss
	TermPrice   decimal.Decimal `json:"termPrinc"`   // 本期应还本金 (分)
	TermInter   decimal.Decimal `json:"TermInter"`   // 本期应还利息 (分)
	TermAmt     decimal.Decimal `json:"termAmt"`     // 本期应还金额 (分)
	Overdue     string          `json:"overdue"`     // 本期是否逾期 0:待还款1:已还款2:逾期
	OverdueDays string          `json:"overdueDays"` // 逾期天数
	RepayType   string          `json:"repayType"`   // 还款类型 0：正常还款1：逾期还款2：提前还款
}
