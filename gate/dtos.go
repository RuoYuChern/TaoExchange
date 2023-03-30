package gate

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
	OFFER  = "OFFER"
	BID    = "BID"
	UN_SID = "UNSID"
)

type OrderCmd string

const (
	CMD_PLACE  = "PLACE"
	CMD_MOVE   = "MOVE"
	CMD_CANCEL = "CANCEL"
	CMD_REDUCE = "REDUCE"
	CMD_QUERY  = "QUERY"
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
	Commond         OrderCmd    `json:"commond"`
	Status          OrderStatus `json:"status"`
	Price           uint64      `json:"price"`
	ReserveBidPrice uint64      `json:"reserveBidPrice"`
	Amount          uint64      `json:"amount"`
	Filled          uint64      `json:"filled"`
	TakerFee        uint64      `json:"takerFee"`
	MakerFee        uint64      `json:"makerFee"`
	Fee             uint64      `json:"fee"`
	Market          string      `json:"market"`
	UserId          string      `json:"userId"`
	Source          string      `json:"source"`
	Version         int         `json:"version"`
	Timestamp       uint64      `json:"timestamp"`
}

type QueryReq struct {
	UserId  string `json:"userId"`
	Market  string `json:"market"`
	OrderId string `json:"orderId"`
}

type QueryResp struct {
	Status int        `json:"status"`
	Msg    string     `json:"msg"`
	Orders []OrderDto `json:"orders"`
}

type OrderReq struct {
	UserId string   `json:"userId"`
	Market string   `json:"market"`
	Order  OrderDto `json:"order"`
}

type OrderResp struct {
	UserId string   `json:"userId"`
	Market string   `json:"market"`
	Order  OrderDto `json:"order"`
}
