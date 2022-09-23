package s60

import (
	"context"
	json "github.com/json-iterator/go"
	s60 "github.com/zrb-channel/s60/schema"
)

// CreditManual 人工审核中
func CreditManual(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditManualResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	return nil
}
