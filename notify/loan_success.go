package s60

import (
	"context"

	json "github.com/json-iterator/go"
)

// LoanSuccess 动支成功
func LoanSuccess(ctx context.Context, raw json.RawMessage) error {
	/**
	body := &s60.LoanSuccessResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}
	*/

	return nil
}
