package gate

import (
	"container/list"
	"context"
	"net/http"
	"sync"
	"time"

	"tao.exchange.com/facade"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "tao.exchange.com/grpc"
)

type marketServer struct {
	client pb.TaoExchangeSrvClient
	conn   *grpc.ClientConn
}

var instance *marketServer
var once sync.Once

var cmdMap = map[string]pb.OrderCmd{
	facade.CMD_PLACE:  pb.OrderCmd_CMD_PLACE,
	facade.CMD_MOVE:   pb.OrderCmd_CMD_MOVE,
	facade.CMD_CANCEL: pb.OrderCmd_CMD_CANCEL,
	facade.CMD_REDUCE: pb.OrderCmd_CMD_REDUCE,
}

var typeMap = map[string]pb.OrderType{
	facade.GTC:        pb.OrderType_OT_GTC,
	facade.IOC:        pb.OrderType_OT_IOC,
	facade.IOC_BUDGET: pb.OrderType_OT_IOC_BUDGET,
	facade.FOK:        pb.OrderType_OT_FOK,
	facade.FOK_BUDGET: pb.OrderType_OT_FOK_BUDGET,
}

var sideMap = map[string]pb.OrderSid{
	facade.ASK: pb.OrderSid_SD_OFFER,
	facade.BID: pb.OrderSid_SD_BID,
}

var statusMap = map[string]pb.OrderStatus{
	facade.INITIAL: pb.OrderStatus_ST_INITIAL,
	facade.ACCEPT:  pb.OrderStatus_ST_ACCEPT,
	facade.REJECT:  pb.OrderStatus_ST_REJECT,
	facade.CANCEL:  pb.OrderStatus_ST_CANCEL,
	facade.FINISH:  pb.OrderStatus_ST_FINISH,
}

func getMarket() *marketServer {
	once.Do(func() {
		instance = &marketServer{}
	})
	return instance
}

func (s *marketServer) connect() error {
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

func (s *marketServer) shutdown() {
	if s.conn != nil {
		s.conn.Close()
	}
}

func (s *marketServer) query(req *facade.QueryReq) *facade.QueryResp {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	exQReq := pb.QueryReq{}
	exQReq.UserId = req.UserId
	exQReq.Market = req.Market
	exQReq.OrderId = req.OrderId
	exRsp, err := s.client.QueryOrder(ctx, &exQReq)
	if err != nil {
		rsp := facade.MakeQueryRsp(http.StatusInternalServerError, err.Error())
		return rsp
	}
	rsp := facade.MakeQueryRsp(exRsp.GetStatus(), exRsp.GetMsg())
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

func (s *marketServer) placeOrder(req *facade.OrderReq) *facade.OrderResp {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	exOrder := pb.OrderReq{}
	exOrder.Cmd = cmdFromCMD(req.Commond)
	exOrder.Market = req.Market
	exOrder.UserId = req.UserId
	exOrder.Order = ordFromOrd(&req.Order)
	exRsp, err := s.client.DoOrderCommond(ctx, &exOrder)
	if err != nil {
		rsp := facade.MakeOrderResp(http.StatusInternalServerError, err.Error(), req.UserId, req.Market)
		return rsp
	}

	slog.Info("replyId:", exRsp.Order.ReplyId)
	rsp := facade.MakeOrderResp(http.StatusOK, "OK", req.UserId, req.Market)
	rsp.Order = *ordToOrd(exRsp.Order.Order)
	return rsp
}

func cmdFromCMD(cmd facade.OrderCmd) pb.OrderCmd {
	if v, ok := cmdMap[string(cmd)]; ok {
		return v
	}
	return pb.OrderCmd_CMD_UNKNOW
}

func typeFromType(otp facade.OrderType) pb.OrderType {
	if v, ok := typeMap[string(otp)]; ok {
		return v
	}
	return pb.OrderType_OT_UNKNOW
}

func sideFromSide(side facade.OrderSid) pb.OrderSid {
	if v, ok := sideMap[string(side)]; ok {
		return v
	}
	return pb.OrderSid_SD_UNKNOW
}

func statusFromStatus(st facade.OrderStatus) pb.OrderStatus {
	if v, ok := statusMap[string(st)]; ok {
		return v
	}
	return pb.OrderStatus_ST_INITIAL
}

func ordFromOrd(dto *facade.OrderDto) *pb.OrderDto {
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

func ordToOrd(d *pb.OrderDto) *facade.OrderDto {
	dto := new(facade.OrderDto)

	dto.Id = d.GetId()
	dto.OdType = facade.OrderType(d.GetType().String())
	dto.Side = facade.OrderSid(d.GetSide().String())
	dto.Status = facade.OrderStatus(d.GetStatus().String())
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
