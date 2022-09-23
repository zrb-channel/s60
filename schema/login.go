package s60

import (
	json "github.com/json-iterator/go"
)

type LoginRequest struct {
	SourceType    string `json:"sourceType"`        // 渠道类型
	UserNo        string `json:"userNo"`            // 合作方用户号
	SubChannel    string `json:"subChannel"`        // 子渠道
	MobileNo      string `json:"mobileNo"`          // 手机号码
	ChannelSource string `json:"channelSource"`     // 渠道编号
	HostApp       string `json:"hostApp,omitempty"` // 宿主信息
}

func (r *LoginRequest) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *LoginRequest) ToString() string {
	v, _ := json.Marshal(r)
	return string(v)
}

type LoginResponse struct {
	BizCode string `json:"bizCode"`
	Code    int    `json:"code"`
	Msg     string `json:"msg,omitempty"`
	Desc    string `json:"desc,omitempty"`
	Data    string `json:"data"`
}

type UserTokenResponse struct {
	ExpireTime string `json:"expireTime"`
	Token      string `json:"token"`
}
