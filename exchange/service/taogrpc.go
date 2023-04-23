package exchange

import (
	"context"
	"golang.org/x/exp/slog"
	"net/http"
	pb "tao.exchange.com/grpc"
)

type TaoExServer struct {
	pb.UnimplementedTaoExchangeSrvServer
}

func (s *TaoExServer) QueryOrder(ctx context.Context, in *pb.QueryReq) (*pb.QueryRsp, error) {
	slog.Info("QueryOrder", slog.Group("Args", slog.String("userId", in.UserId), slog.String("market", in.Market), slog.String("market", in.Market),
		slog.String("orderId", in.OrderId)))
	return &pb.QueryRsp{Status: http.StatusNotImplemented, Msg: "Not supported"}, nil
}

func (s *TaoExServer) DoOrderCommond(ctx context.Context, in *pb.OrderReq) (*pb.OrderRsp, error) {
	slog.Info("QueryOrder", slog.Group("Args", slog.String("userId", in.UserId), slog.String("market", in.Market), slog.String("market", in.Market),
		slog.String("orderId", in.Order.Id)))
	return &pb.OrderRsp{Status: http.StatusNotImplemented, Msg: "Not supported"}, nil
}
