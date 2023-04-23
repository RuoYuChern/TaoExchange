package main

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"net"
	"os/signal"
	"syscall"
	"tao.exchange.com/common"
	"tao.exchange.com/coordinator"
	pb "tao.exchange.com/grpc"
	"tao.exchange.com/infra"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 启动DB 连接
	taoConf := common.TaoAppConf{}
	taoConf.LoadTaoConf("../tao_conf.yaml")
	db := infra.GetDbCon()
	err := db.Connect(taoConf.Infra.DbDns, &ctx)
	if err != nil {
		slog.Error("db connect:", err.Error())
		panic(err)
		return
	}
	// load lock info
	coordinator.GetSs().Int()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", taoConf.Tao.CoordinatorPort))
	if err != nil {
		slog.Error("Failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterTaoCoordinatorSrvServer(s, &coordinator.TaoCoServer{})

	go func() {
		slog.Info("Server listening at:", lis.Addr().String())
		if err := s.Serve(lis); err != nil {
			slog.Error("Failed to server:", err)
			return
		}
	}()

	coordinator.StartTaoCoordinatorRest(taoConf.Tao.CoordinatorRestPort)
	//Listen for the interrupt signal
	<-ctx.Done()
	stop()
	slog.Info("Shutdown Server ...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.GracefulStop()
	coordinator.GraceFulStop(&ctx)
	common.Get().Close()
	slog.Info("Server exist")
}
