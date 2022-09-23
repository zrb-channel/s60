package s60

import (
	"context"
	json "github.com/json-iterator/go"
	s60 "github.com/zrb-channel/s60/schema"
)

// CreditReject 授信拒绝
func CreditReject(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditRejectResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	return nil
}
