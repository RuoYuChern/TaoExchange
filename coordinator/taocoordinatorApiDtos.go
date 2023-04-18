package coordinator

import "net/url"

type coordinatorRsp struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
}

type taoShardMarketDto struct {
	MarketId   string `json:"marketId"`
	ShardId    string `json:"shardId"`
	UpdateTime string `json:"updateTime"`
}

type taoAddShardMarkeReq struct {
	ReqId string              `json:"reqId"`
	Data  []taoShardMarketDto `json:"data"`
}

type taoLockDto struct {
	ShardId   string `json:"shardId"`
	AppId     string `json:"appId"`
	AppIp     string `json:"appIp"`
	AppRole   string `json:"appRole"`
	AppPort   int32  `json:"appPort"`
	AppStatus int32  `json:"appStatus"`
	LockTime  string `json:"lockTime"`
}

type taoLockReq struct {
	ReqId string     `json:"reqId"`
	Data  taoLockDto `json:"data"`
}

type TaoHttpQueryCache struct {
	qc url.Values
}

type taoMarketDto struct {
	Market       string `json:"market"`
	Base         string `json:"base"`
	Pair         string `json:"pair"`
	MarketStatus int32  `json:"marketStatus"`
	BasePrec     int32  `json:"basePrec"`
	PairPrec     int32  `json:"pairPrec"`
	FeePrec      int32  `json:"feePrec"`
	MinAmount    int64  `json:"minAmount"`
	MinBase      int64  `json:"minBase"`
}

type taoMarketReq struct {
	ReqId string         `json:"reqId"`
	Data  []taoMarketDto `json:"data"`
}
