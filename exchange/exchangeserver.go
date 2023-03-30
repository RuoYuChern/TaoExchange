package exchange

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	pb "tao.exchange.com/grpc"
)

type server struct {
	pb.UnimplementedTaoExchangeSrvServer
}

func (s *server) QueryOrder(ctx context.Context, in *pb.QueryReq) (*pb.QueryRsp, error) {
	slog.Info("QueryOrder", slog.Group("Args", slog.String("userId", in.UserId), slog.String("market", in.Market), slog.String("market", in.Market),
		slog.String("orderId", in.OrderId)))
	return &pb.QueryRsp{Status: http.StatusNotImplemented, Msg: "Not supported"}, nil
}

func (s *server) DoOrderCommond(ctx context.Context, in *pb.OrderReq) (*pb.OrderRsp, error) {
	slog.Info("QueryOrder", slog.Group("Args", slog.String("userId", in.UserId), slog.String("market", in.Market), slog.String("market", in.Market),
		slog.String("orderId", in.Order.Id)))
	return &pb.OrderRsp{Status: http.StatusNotImplemented, Msg: "Not supported"}, nil
}

var (
	port = flag.Int("port", 58081, "The server port")
)

func StartTaoExchange() {
	// create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		slog.Error("Failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterTaoExchangeSrvServer(s, &server{})

	go func() {
		slog.Info("Server listening at:", lis.Addr())
		if err := s.Serve(lis); err != nil {
			slog.Error("Failed to server:", err)
			return
		}
	}()

	//Listen for the interrupt signal
	<-ctx.Done()
	stop()
	slog.Info("Shutdown Server ...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.GracefulStop()
	slog.Info("Server exist")
}
