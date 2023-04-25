package main

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"tao.exchange.com/common"
	exchange "tao.exchange.com/exchange/service"
	pb "tao.exchange.com/grpc"
)

func main() {
	// create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	taoConf := common.TaoAppConf{}
	taoConf.LoadTaoConf("../tao_conf.yaml")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", taoConf.Tao.ExchangePort))
	if err != nil {
		slog.Error("Failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterTaoExchangeSrvServer(s, &exchange.TaoExServer{})

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
