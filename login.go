package s60

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	s60 "github.com/zrb-channel/s60/schema"
	"github.com/zrb-channel/utils"

	log "github.com/zrb-channel/utils/logger"

	json "github.com/json-iterator/go"
)

// Login
// @param ctx
// @param conf
// @param bizData
// @date 2022-09-24 01:26:15
func Login(ctx context.Context, conf *s60.Config, bizData *s60.LoginRequest) (*s60.UserTokenResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req := s60.NewBaseRequest()

	req.SetData(bizData)

	if err := req.GenSign(conf); err != nil {
		return nil, fmt.Errorf("生成签名失败:%s", err.Error())
	}

	resp, err := utils.Request(ctx).
		SetBody(req).Post(loginAddr)

	if err != nil {
		return nil, fmt.Errorf("请求登录失败:%s", err.Error())
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("响应状态码有误:%s", resp.Status())
	}

	result := &s60.LoginResponse{}

	if err = json.ConfigCompatibleWithStandardLibrary.Unmarshal(resp.Body(), result); err != nil {
		return nil, fmt.Errorf("数据解析失败:%s", err.Error())
	}

	if result.Code != http.StatusOK {
		log.WithData(map[string]any{"addr": loginAddr, "result": result, "body": req, "data": bizData}).Error("响应code不正确")

		return nil, errors.New(result.Desc)
	}

	data := &s60.BaseResponse{}
	if err = json.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(result.Data), data); err != nil {
		return nil, fmt.Errorf("数据解析失败:%s", err.Error())
	}

	if err = data.PrivateVerify(conf.JtPublicKey); err != nil {
		return nil, fmt.Errorf("签名验证失败:%s", err.Error())
	}

	var decryptData []byte
	if decryptData, err = data.DecryptData(conf.PtPrivateKey); err != nil {
		return nil, err
	}

	user := &s60.UserTokenResponse{}
	if err = json.ConfigCompatibleWithStandardLibrary.Unmarshal(decryptData, user); err != nil {
		return nil, fmt.Errorf("数据解析失败:%s", err.Error())
	}

	return user, nil
}
