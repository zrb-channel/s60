package s60

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/aesutil"
	"github.com/zrb-channel/utils/rsautil"

	"github.com/google/go-querystring/query"
	json "github.com/json-iterator/go"
)

type BaseRequest struct {
	AppID      string  `json:"appId" url:"appId"`
	Timestamp  string  `json:"timestamp" url:"timestamp"`
	BizData    string  `json:"bizData" url:"bizData"`
	Data       BizData `json:"-" url:"-"`
	EncryptKey string  `json:"encryptKey" url:"encryptKey"`
	EncryptIV  string  `json:"encryptIV" url:"encryptIV"`
	Sign       string  `json:"sign" url:"-"`
}

type BizData interface {
	ToJson() ([]byte, error)
	ToString() string
}

func NewBaseRequest() *BaseRequest {
	return &BaseRequest{
		AppID:     "HONGYUANKEJI",
		Timestamp: strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}
}

func (r *BaseRequest) SetAppID(id string)     { r.AppID = id }
func (r *BaseRequest) SetData(data BizData)   { r.Data = data }
func (r *BaseRequest) SetBizData(data string) { r.BizData = data }
func (r *BaseRequest) SetBizDataFromByte(data []byte) {
	r.BizData = strings.ToUpper(hex.EncodeToString(data))
}

func (r *BaseRequest) GenSign(conf *Config) error {

	bizData, err := r.Data.ToJson()
	if err != nil {
		return err
	}

	var (
		key         = utils.RandString(16)
		iv          = utils.RandString(16)
		encryptData []byte
	)

	if encryptData, err = aesutil.Encrypt(bizData, []byte(key), []byte(iv)); err != nil {
		return err
	}

	r.SetBizDataFromByte(encryptData)

	publicKey, err := utils.NewPublicKey(conf.JtPublicKey)
	if err != nil {
		return err
	}
	var encryptKey []byte
	if encryptKey, err = rsautil.PublicEncrypt(publicKey, []byte(key)); err != nil {
		return err
	}

	var encryptIV []byte
	if encryptIV, err = rsautil.PublicEncrypt(publicKey, []byte(iv)); err != nil {
		return err
	}

	r.EncryptKey = base64.StdEncoding.EncodeToString(encryptKey)
	r.EncryptIV = base64.StdEncoding.EncodeToString(encryptIV)
	u, _ := query.Values(r)

	signature, _ := url.QueryUnescape(u.Encode())

	privateKey, err := rsautil.PrivateKeyFrom64(conf.PtPrivateKey)
	if err != nil {
		return err
	}

	var sign []byte
	if sign, err = rsautil.PrivateSign(privateKey, []byte(signature)); err != nil {
		return err
	}

	r.Sign = base64.StdEncoding.EncodeToString(sign)

	return nil
}

func (r *BaseRequest) String() string {
	v, _ := json.ConfigCompatibleWithStandardLibrary.Marshal(r)
	return string(v)
}

func (r *BaseRequest) ToJson() []byte {
	v, _ := json.ConfigCompatibleWithStandardLibrary.Marshal(r)
	return v
}

func (r *BaseRequest) ToReader() io.Reader {
	return bytes.NewReader(r.ToJson())
}

type BaseResponse struct {
	BizData    string `json:"bizData" url:"bizData"`
	EncryptKey string `json:"encryptKey" url:"encryptKey"`
	EncryptIV  string `json:"encryptIV" url:"encryptIV"`
	Sign       string `json:"sign" url:"-"`
	Timestamp  string `json:"timestamp" url:"timestamp,omitempty"`
	Method     string `json:"method" url:"method,omitempty"`
	AppID      string `json:"appId" url:"appId,omitempty"`
}

func (r *BaseResponse) PrivateVerify(pubKey string) error {

	sign, err := base64.StdEncoding.DecodeString(r.Sign)
	if err != nil {
		return err
	}

	var values url.Values
	if values, err = query.Values(r); err != nil {
		return err
	}

	var value string
	if value, err = url.QueryUnescape(values.Encode()); err != nil {
		return err
	}

	publicKey, err := utils.NewPublicKey(pubKey)
	if err != nil {
		return err
	}

	return rsautil.PublicVerify(publicKey, sign, []byte(value))
}

func (r *BaseResponse) DecryptData(priKey string) ([]byte, error) {
	enKey, err := base64.StdEncoding.DecodeString(r.EncryptKey)
	if err != nil {
		return nil, err
	}

	privateKey, err := rsautil.PrivateKeyFrom64(priKey)
	if err != nil {
		return nil, err
	}

	var key []byte
	if key, err = rsautil.PrivateDecrypt(privateKey, enKey); err != nil {
		return nil, fmt.Errorf("rsaKey解析失败:%s", err.Error())
	}

	var enIv []byte
	if enIv, err = base64.StdEncoding.DecodeString(r.EncryptIV); err != nil {
		return nil, err
	}

	var iv []byte
	iv, err = rsautil.PrivateDecrypt(privateKey, enIv)
	if err != nil {
		return nil, fmt.Errorf("rsaIV解析失败:%s", err.Error())
	}

	return aesutil.Decrypt(r.BizData, key, iv)
}

type CommonResponse struct {
	OrgNo    string `json:"orgNo"`    // 机构编号
	UserNo   string `json:"userNo"`   // 合作方用户号
	UserName string `json:"userName"` // 用户姓名 MD5加密
	IdNumber string `json:"idNumber"` // 身份证号 MD5加密
	FlowNo   string `json:"flowNo"`   // 流水号 由360提供
	OrderNo  string `json:"orderNo"`  // 订单编号 渠道方提供
}
