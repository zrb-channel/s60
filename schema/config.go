package s60

type Config struct {
	// SourceType 渠道类型
	SourceType string `json:"sourceType"`

	// ChannelSource 渠道编号
	ChannelSource string `json:"channelSource"`

	SubChannel string `json:"subChannel"`

	JtPublicKey string `json:"jtPublicKey"`

	PtPrivateKey string `json:"PtPrivateKey"`
}
