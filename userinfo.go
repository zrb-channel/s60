package s60

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zrb-channel/utils"

	s60 "github.com/zrb-channel/s60/schema"

	json "github.com/json-iterator/go"
)

// UserInfo
// @param ctx
// @param conf
// @param info
// @date 2022-09-24 01:28:12
func UserInfo(ctx context.Context, conf *s60.Config, info *s60.UserInfoRequest) (*s60.UserInfoData, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req := s60.NewBaseRequest()
	req.SetData(info)

	if err := req.GenSign(conf); err != nil {
		return nil, fmt.Errorf("生成签名失败:%s", err.Error())
	}

	resp, err := utils.Request(ctx).
		SetBody(req).
		Post(userInfoAddr)

	if err != nil {
		return nil, fmt.Errorf("提交失败:%s", err.Error())
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf(resp.Status())
	}

	body := &s60.UserInfoResponse{}
	if err = json.Unmarshal(resp.Body(), body); err != nil {
		return nil, err
	}

	if body.Code != http.StatusOK {
		return nil, fmt.Errorf("提交用户信息失败:%s | desc:%s", body.Msg, body.Desc)
	}

	data := &s60.BaseResponse{}
	if err = json.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(body.Data), data); err != nil {
		return nil, fmt.Errorf("数据解析失败:%s", err.Error())
	}

	if err = data.PrivateVerify(conf.JtPublicKey); err != nil {
		return nil, fmt.Errorf("签名验证失败:%s", err.Error())
	}

	var decryptData []byte
	if decryptData, err = data.DecryptData(conf.PtPrivateKey); err != nil {
		return nil, err
	}

	user := &s60.UserInfoData{}
	if err = json.ConfigCompatibleWithStandardLibrary.Unmarshal(decryptData, user); err != nil {
		return nil, fmt.Errorf("数据解析失败:%s", err.Error())
	}

	return user, nil
}
