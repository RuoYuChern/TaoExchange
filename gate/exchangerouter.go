package gate

import (
	"container/list"
	"context"
	"net/http"
	"sync"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "tao.exchange.com/grpc"
)

type exRouterServer struct {
	client pb.TaoExchangeSrvClient
	conn   *grpc.ClientConn
}

var instance *exRouterServer
var once sync.Once

var cmdMap = map[string]pb.OrderCmd{
	CMD_PLACE:  pb.OrderCmd_CMD_PLACE,
	CMD_MOVE:   pb.OrderCmd_CMD_MOVE,
	CMD_CANCEL: pb.OrderCmd_CMD_CANCEL,
	CMD_REDUCE: pb.OrderCmd_CMD_REDUCE,
}

var typeMap = map[string]pb.OrderType{
	GTC:        pb.OrderType_OT_GTC,
	IOC:        pb.OrderType_OT_IOC,
	IOC_BUDGET: pb.OrderType_OT_IOC_BUDGET,
	FOK:        pb.OrderType_OT_FOK,
	FOK_BUDGET: pb.OrderType_OT_FOK_BUDGET,
}

var sideMap = map[string]pb.OrderSid{
	OFFER: pb.OrderSid_SD_OFFER,
	BID:   pb.OrderSid_SD_BID,
}

var statusMap = map[string]pb.OrderStatus{
	INITIAL: pb.OrderStatus_ST_INITIAL,
	ACCEPT:  pb.OrderStatus_ST_ACCEPT,
	REJECT:  pb.OrderStatus_ST_REJECT,
	CANCEL:  pb.OrderStatus_ST_CANCEL,
	FINISH:  pb.OrderStatus_ST_FINISH,
}

func getRouter() *exRouterServer {
	once.Do(func() {
		instance = &exRouterServer{}
	})
	return instance
}

func (s *exRouterServer) connect() error {
	addr := "127.0.0.1:58081"
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		slog.Warn("Connect error")
		return err
	}
	s.client = pb.NewTaoExchangeSrvClient(conn)
	s.conn = conn
	return nil
}

func (s *exRouterServer) shutdown() {
	if s.conn != nil {
		s.conn.Close()
	}
}

func (s *exRouterServer) query(req *QueryReq) *QueryResp {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	exQReq := pb.QueryReq{}
	exQReq.UserId = req.UserId
	exQReq.Market = req.Market
	exQReq.OrderId = req.OrderId
	exRsp, err := s.client.QueryOrder(ctx, &exQReq)
	if err != nil {
		rsp := makeQueryRsp(http.StatusInternalServerError, err.Error())
		return rsp
	}
	rsp := makeQueryRsp(exRsp.GetStatus(), exRsp.GetMsg())
	if exRsp.GetOrderList() != nil {
		size := len(exRsp.GetOrderList())
		if size > 0 {
			ordList := list.New()
			for _, v := range exRsp.GetOrderList() {
				ordList.PushBack(ordToOrd(v))
			}
			rsp.Orders = ordList
		}
	}
	return rsp
}

func (s *exRouterServer) placeOrder(req *OrderReq) *OrderResp {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	exOrder := pb.OrderReq{}
	exOrder.Cmd = cmdFromCMD(req.Commond)
	exOrder.Market = req.Market
	exOrder.UserId = req.UserId
	exOrder.Order = ordFromOrd(&req.Order)
	exRsp, err := s.client.DoOrderCommond(ctx, &exOrder)
	if err != nil {
		rsp := makeOrderResp(http.StatusInternalServerError, err.Error(), req.UserId, req.Market)
		return rsp
	}

	slog.Info("replyId:", exRsp.Order.ReplyId)
	rsp := makeOrderResp(http.StatusOK, "OK", req.UserId, req.Market)
	rsp.Order = *ordToOrd(exRsp.Order.Order)
	return rsp
}

func cmdFromCMD(cmd OrderCmd) pb.OrderCmd {
	if v, ok := cmdMap[string(cmd)]; ok {
		return v
	}
	return pb.OrderCmd_CMD_UNKNOW
}

func typeFromType(otp OrderType) pb.OrderType {
	if v, ok := typeMap[string(otp)]; ok {
		return v
	}
	return pb.OrderType_OT_UNKNOW
}

func sideFromSide(side OrderSid) pb.OrderSid {
	if v, ok := sideMap[string(side)]; ok {
		return v
	}
	return pb.OrderSid_SD_UNKNOW
}

func statusFromStatus(st OrderStatus) pb.OrderStatus {
	if v, ok := statusMap[string(st)]; ok {
		return v
	}
	return pb.OrderStatus_ST_INITIAL
}

func ordFromOrd(dto *OrderDto) *pb.OrderDto {
	ord := new(pb.OrderDto)
	ord.Id = dto.Id
	ord.Type = typeFromType(dto.OdType)
	ord.Side = sideFromSide(dto.Side)
	ord.Status = statusFromStatus(dto.Status)
	ord.Version = dto.Version
	ord.Timestamp = dto.Timestamp
	ord.Price = dto.Price
	ord.ReserveBidPrice = dto.ReserveBidPrice
	ord.Amount = dto.Amount
	ord.Filled = dto.Filled
	ord.MakerFee = dto.MakerFee
	ord.TakerFee = dto.TakerFee
	ord.Fee = dto.Fee
	ord.Market = dto.Market
	ord.UserId = dto.UserId
	ord.Source = dto.Source
	return ord
}

func ordToOrd(d *pb.OrderDto) *OrderDto {
	dto := new(OrderDto)

	dto.Id = d.GetId()
	dto.OdType = OrderType(d.GetType().String())
	dto.Side = OrderSid(d.GetSide().String())
	dto.Status = OrderStatus(d.GetStatus().String())
	dto.Price = d.GetPrice()
	dto.ReserveBidPrice = d.GetReserveBidPrice()
	dto.Amount = d.GetAmount()
	dto.Filled = d.GetFilled()
	dto.TakerFee = d.GetTakerFee()
	dto.MakerFee = d.GetMakerFee()
	dto.Fee = d.GetFee()
	dto.Market = d.GetMarket()
	dto.UserId = d.GetUserId()
	dto.Source = d.GetSource()
	dto.Version = d.GetVersion()
	dto.Timestamp = d.GetTimestamp()

	return dto
}
