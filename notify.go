package s60

import (
	"context"
	"errors"
	notify "github.com/zrb-channel/s60/notify"
	s60 "github.com/zrb-channel/s60/schema"
	log "github.com/zrb-channel/utils/logger"

	json "github.com/json-iterator/go"
)

type (
	NotifyHandleFunc func(context.Context, json.RawMessage) error

	NotifyHandlers interface {
		OnCreditFinish(ctx context.Context, req s60.CreditFinishResponse) error
	}
)

var notifyHandlers NotifyHandlers

var handlers = make(map[string]NotifyHandleFunc)

func init() {
	handlers["credit_finish"] = notify.CreditFinish // 用户提交授信审核
	handlers["credit_create"] = notify.CreditCreate // 完善信息中
	handlers["credit_cancel"] = notify.CreditCancel // 申请取消通知
	handlers["credit_manual"] = notify.CreditManual // 人工审核中通知
	handlers["credit_pass"] = notify.CreditPass     // 授信审核通过
	handlers["credit_reject"] = notify.CreditReject // 授信拒绝
	handlers["loan_success"] = notify.LoanSuccess   // 动支成功
	handlers["loan_info"] = notify.LoanInfo         // 放款信息通知
	handlers["repay_plan"] = notify.RepayPlan       // 还款计划
	handlers["overdue"] = notify.Overdue            // 逾期
	handlers["repay_success"] = notify.RepaySuccess // 还款成功
}

func RegisterNotifyHandlers(handlers NotifyHandlers) {
	notifyHandlers = handlers
}

func Notify(ctx context.Context, conf *s60.Config, resp *s60.BaseResponse) error {

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

	log.Info("处理成功")

	return nil
}
