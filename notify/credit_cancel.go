package s60

import (
	"context"
	json "github.com/json-iterator/go"
	s60 "github.com/zrb-channel/s60/schema"
)

// CreditCancel 申请取消通知
func CreditCancel(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditCancelResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	return nil
}
