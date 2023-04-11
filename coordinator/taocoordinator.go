package coordinator

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	pb "tao.exchange.com/grpc"
)

type server struct {
	pb.UnimplementedTaoCoordinatorSrvServer
}

func (s *server) ListShards(ctx context.Context, req *pb.CommonReq) (*pb.LockShardRsp, error) {
	return nil, nil
}
func (s *server) LockShard(ctx context.Context, req *pb.ShardReq) (*pb.LockShardRsp, error) {
	return nil, nil
}
func (s *server) UnlockShard(ctx context.Context, req *pb.ShardReq) (*pb.CommonRsp, error) {
	return nil, nil
}
func (s *server) ListConnectorInfo(ctx context.Context, req *pb.CommonReq) (*pb.ListConnectorRsp, error) {
	return nil, nil
}
func (s *server) KeepLive(ctx context.Context, req *pb.ShardReq) (*pb.CommonRsp, error) {
	return nil, nil
}

var (
	port = flag.Int("port", 59081, "The server port")
)

func StartTaoCoordinator() {
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
	pb.RegisterTaoCoordinatorSrvServer(s, &server{})

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
