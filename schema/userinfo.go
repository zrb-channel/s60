package s60

import json "github.com/json-iterator/go"

type UserInfoRequest struct {
	UserNo        string `json:"userNo"`
	ChannelSource string `json:"channelSource"`           // 渠道编号
	SourceType    string `json:"sourceType"`              // 渠道类型
	OrderNo       string `json:"orderNo,omitempty"`       // 订单编号
	CompanyName   string `json:"companyName,omitempty"`   // 企业名称
	CreditNo      string `json:"creditNo,omitempty"`      // 统一社会信用代码
	IdNumber      string `json:"idNumber,omitempty"`      // 法人身份证号
	CorporateName string `json:"corporateName,omitempty"` // 法人姓名
	MobileNo      string `json:"mobileNo,omitempty"`      // 手机号码
}

func (r *UserInfoRequest) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *UserInfoRequest) ToString() string {
	v, _ := json.Marshal(r)
	return string(v)
}

type UserInfoResponse struct {
	BizCode string `json:"bizCode"`
	Code    int    `json:"code"`
	Msg     string `json:"msg,omitempty"`
	Desc    string `json:"desc,omitempty"`
	Data    string `json:"data"`
}

type UserInfoData struct {
	UserNo string `json:"userNo"`
}
