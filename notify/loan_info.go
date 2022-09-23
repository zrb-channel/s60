package s60

import (
	"context"
	json "github.com/json-iterator/go"
	s60 "github.com/zrb-channel/s60/schema"
)

// LoanInfo 放款信息通知
func LoanInfo(ctx context.Context, raw json.RawMessage) error {
	body := &s60.LoanInfoResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	return nil
}
