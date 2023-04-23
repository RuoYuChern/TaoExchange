package facade

import (
	"container/list"
)

type OrderType string

const (
	GTC        = "GTC"
	IOC        = "IOC"
	IOC_BUDGET = "IOC_BUDGET"
	FOK        = "FOK"
	FOK_BUDGET = "FOK_BUDGET"
	UN_TYPE    = "UNTYPE"
)

type OrderSid string

const (
	ASK    = "OFFER"
	BID    = "BID"
	UN_SID = "UNSID"
)

type OrderCmd string

const (
	CMD_PLACE  = "PLACE"
	CMD_MOVE   = "MOVE"
	CMD_CANCEL = "CANCEL"
	CMD_REDUCE = "REDUCE"
	UN_CMD     = "UNCMD"
)

type OrderStatus string

const (
	INITIAL = "INITIAL"
	ACCEPT  = "ACCEPT"
	REJECT  = "REJECT"
	CANCEL  = "CANCEL"
	FINISH  = "FINISH"
)

type OrderDto struct {
	Id              string      `json:"id"`
	OdType          OrderType   `json:"type"`
	Side            OrderSid    `json:"side"`
	Status          OrderStatus `json:"status"`
	Price           int64       `json:"price"`
	ReserveBidPrice int64       `json:"reserveBidPrice"`
	Amount          int64       `json:"amount"`
	Filled          int64       `json:"filled"`
	TakerFee        int64       `json:"takerFee"`
	MakerFee        int64       `json:"makerFee"`
	Fee             int64       `json:"fee"`
	Market          string      `json:"market"`
	UserId          string      `json:"userId"`
	Source          string      `json:"source"`
	Version         int32       `json:"version"`
	Timestamp       int64       `json:"timestamp"`
}

type QueryReq struct {
	UserId  string `json:"userId"`
	Market  string `json:"market"`
	OrderId string `json:"orderId"`
}

type QueryResp struct {
	Status int32      `json:"status"`
	Msg    string     `json:"msg"`
	Orders *list.List `json:"orders"`
}

type OrderReq struct {
	UserId  string   `json:"userId"`
	Market  string   `json:"market"`
	Order   OrderDto `json:"order"`
	Commond OrderCmd `json:"commond"`
}

type OrderResp struct {
	Status int32    `json:"status"`
	Msg    string   `json:"msg"`
	UserId string   `json:"userId"`
	Market string   `json:"market"`
	Order  OrderDto `json:"order"`
}

func MakeQueryRsp(status int32, msg string) *QueryResp {
	rsp := new(QueryResp)
	rsp.Status = status
	rsp.Msg = msg
	return rsp
}

func MakeOrderResp(status int32, msg, userId, market string) *OrderResp {
	rsp := new(OrderResp)
	rsp.Status = status
	rsp.Msg = msg
	rsp.Market = market
	rsp.UserId = userId
	return rsp
}
