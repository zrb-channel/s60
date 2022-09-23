package s60

import (
	"context"
	json "github.com/json-iterator/go"
	s60 "github.com/zrb-channel/s60/schema"
)

// CreditFinish 用户提交授信审核
func CreditFinish(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditFinishResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	return nil
}
