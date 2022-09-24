package s60

import (
	"context"
	"errors"
	json "github.com/json-iterator/go"
	s60 "github.com/zrb-channel/s60/schema"
	log "github.com/zrb-channel/utils/logger"
)

type (
	NotifyHandleFunc func(context.Context, json.RawMessage) error

	NotifyHandlers interface {
		// OnCreditFinish 用户提交授信审核
		OnCreditFinish(ctx context.Context, req *s60.CreditFinishResponse) error

		// OnCreditCreate 完善信息中
		OnCreditCreate(ctx context.Context, req *s60.CreditCreateResponse) error

		// OnCreditCancel 申请取消通知
		OnCreditCancel(ctx context.Context, req *s60.CreditCancelResponse) error

		// OnCreditManual 人工审核中
		OnCreditManual(ctx context.Context, req *s60.CreditManualResponse) error

		// OnCreditPass 授信审核通过
		OnCreditPass(ctx context.Context, req *s60.CreditPassResponse) error

		// OnCreditReject 授信拒绝
		OnCreditReject(ctx context.Context, req *s60.CreditRejectResponse) error

		// OnLoanSuccess 动支成功
		OnLoanSuccess(ctx context.Context, req *s60.LoanSuccessResponse) error

		// OnLoanInfo 放款信息通知
		OnLoanInfo(ctx context.Context, req *s60.LoanInfoResponse) error

		// OnRepayPlan 还款计划
		OnRepayPlan(ctx context.Context, req *s60.RepayPlanResponse) error

		// OnOverdue 逾期
		OnOverdue(ctx context.Context, req *s60.RepayPlanResponse) error

		// OnRepaySuccess 还款成功
		OnRepaySuccess(ctx context.Context, req *s60.RepayPlanResponse) error
	}
)

var notifyHandlers NotifyHandlers

var handlers = make(map[string]NotifyHandleFunc)

func init() {
	handlers["credit_finish"] = CreditFinish // 用户提交授信审核
	handlers["credit_create"] = CreditCreate // 完善信息中
	handlers["credit_cancel"] = CreditCancel // 申请取消通知
	handlers["credit_manual"] = CreditManual // 人工审核中通知
	handlers["credit_pass"] = CreditPass     // 授信审核通过
	handlers["credit_reject"] = CreditReject // 授信拒绝
	handlers["loan_success"] = LoanSuccess   // 动支成功
	handlers["loan_info"] = LoanInfo         // 放款信息通知
	handlers["repay_plan"] = RepayPlan       // 还款计划
	handlers["overdue"] = Overdue            // 逾期
	handlers["repay_success"] = RepaySuccess // 还款成功
}

func RegisterNotifyHandlers(handlers NotifyHandlers) {
	notifyHandlers = handlers
}

func Notify(ctx context.Context, conf *s60.Config, req []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	resp := &s60.BaseResponse{}
	if err := json.Unmarshal(req, resp); err != nil {
		return err
	}

	if err := resp.PrivateVerify(conf.JtPublicKey); err != nil {
		log.WithError(err).Error("签名验证失败")
		return err
	}

	data, err := resp.DecryptData(conf.PtPrivateKey)
	if err != nil {
		log.WithError(err).Error("数据解密失败")
		return err
	}

	handler, ok := handlers[resp.Method]
	if !ok {
		log.WithData(map[string]any{"method": resp.Method}).Warn("不存在的方法")
		return errors.New("method not exits")
	}

	if err = handler(ctx, data); err != nil {
		log.WithError(err).Error("处理失败")
		return err
	}

	return nil
}

// CreditFinish 用户提交授信审核
func CreditFinish(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditFinishResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}
	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnCreditFinish(ctx, body)
}

// CreditCreate 完善信息中
func CreditCreate(ctx context.Context, raw json.RawMessage) error {

	body := &s60.CreditCreateResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnCreditCreate(ctx, body)
}

// CreditCancel 申请取消通知
func CreditCancel(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditCancelResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnCreditCancel(ctx, body)
}

// CreditManual 人工审核中
func CreditManual(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditManualResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}
	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnCreditManual(ctx, body)
}

// CreditPass 授信审核通过
func CreditPass(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditPassResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnCreditPass(ctx, body)
}

// CreditReject 授信拒绝
func CreditReject(ctx context.Context, raw json.RawMessage) error {
	body := &s60.CreditRejectResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnCreditReject(ctx, body)
}

// LoanSuccess 动支成功
func LoanSuccess(ctx context.Context, raw json.RawMessage) error {
	body := &s60.LoanSuccessResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}
	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnLoanSuccess(ctx, body)
}

// LoanInfo 放款信息通知
func LoanInfo(ctx context.Context, raw json.RawMessage) error {
	body := &s60.LoanInfoResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnLoanInfo(ctx, body)
}

// RepayPlan 还款计划
func RepayPlan(ctx context.Context, raw json.RawMessage) error {
	body := &s60.RepayPlanResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}
	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnRepayPlan(ctx, body)
}

// Overdue 逾期
func Overdue(ctx context.Context, raw json.RawMessage) error {
	body := &s60.RepayPlanResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}
	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnOverdue(ctx, body)
}

// RepaySuccess 还款成功
func RepaySuccess(ctx context.Context, raw json.RawMessage) error {
	body := &s60.RepayPlanResponse{}
	if err := json.Unmarshal(raw, body); err != nil {
		return err
	}
	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnRepaySuccess(ctx, body)
}
