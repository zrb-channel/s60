package s60

import (
	json "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

type LoanInfoResponse struct {
	CommonResponse
	LoanNo          string          `json:"loanNo"`          // 放款编号
	Amount          decimal.Decimal `json:"amount"`          // 放款金额（分）
	Term            string          `json:"term"`            //  期限（月)
	LoanTime        string          `json:"loanTime"`        // 放款时间
	RemainingAmount decimal.Decimal `json:"remainingAmount"` // 剩余额度（分）
	SettleFlag      string          `json:"settleFlag"`      // 结算标识
}

func (l *LoanInfoResponse) String() string {
	v, _ := json.Marshal(l)
	return string(v)
}
