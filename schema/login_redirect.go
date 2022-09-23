package s60

import (
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
)

type LoginRedirect struct {
	Timestamp     string `json:"t" url:"t"`
	AuthToken     string `json:"authToken" url:"authToken"`
	NeedLogin     string `json:"needLogin" url:"needLogin"`
	SourceType    string `json:"sourceType" url:"sourceType"`
	ChannelSource string `json:"channelSource" url:"channelSource"`
	SubChannel    string `json:"subChannel" url:"subChannel"`

	// ENTERPRISE:发票贷产品，用户在渠道侧完成发票数据授权，由渠道方提供发票数据
	// INVOICE:发票贷产品，用户在360侧完成发票数据授权
	// TAXATION:税贷产品，用户在渠道侧完成税务数据授权，由渠道方提供税务数据
	// INVOICE_TAXATION:税贷产品，用户在360侧完成税务数据授权
	// TOBACCO：烟草贷产品
	CurrentApplyProductType string `json:"currentApplProductType"  url:"currentApplProductType"`
}

func (r *LoginRedirect) Params() (string, error) {
	u, err := query.Values(r)
	if err != nil {
		return "", err
	}

	u.Set("serverEnv", "PRD")

	return u.Encode(), nil
}

func NewLoginRedirect() *LoginRedirect {
	return &LoginRedirect{
		NeedLogin: "Y",
		Timestamp: strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}
}
