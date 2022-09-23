package s60

import (
	"context"

	json "github.com/json-iterator/go"
)

// RepaySuccess 还款成功
func RepaySuccess(ctx context.Context, raw json.RawMessage) error {
	/**
	body := &s60.RepayPlanResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}
	*/
	return nil
}
