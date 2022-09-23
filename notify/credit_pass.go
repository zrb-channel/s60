package s60

import (
	"context"
	json "github.com/json-iterator/go"
	s60 "github.com/zrb-channel/s60/schema"
)

// CreditPass 授信审核通过
func CreditPass(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditPassResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	return nil
}
