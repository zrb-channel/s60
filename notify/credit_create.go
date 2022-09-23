package s60

import (
	"context"
	json "github.com/json-iterator/go"
	s60 "github.com/zrb-channel/s60/schema"
)

// CreditCreate 完善信息中
func CreditCreate(ctx context.Context, raw json.RawMessage) error {

	body := &s60.CreditCreateResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	return nil
}
