package s60

import (
	json "github.com/json-iterator/go"
)

type SubmitUserCreditRequest struct {
	OrgNo       string `json:"orgNo"`
	UserNo      string `json:"userNo"`
	UserName    string `json:"userName"`
	IDNumber    string `json:"idNumber"`
	FlowNo      string `json:"flowNo"`
	OrderNo     string `json:"orderNo"`
	ApplyStatus string `json:"applyStatus"`
}

func (r *SubmitUserCreditRequest) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *SubmitUserCreditRequest) ToString() string {
	v, _ := json.Marshal(r)
	return string(v)
}
