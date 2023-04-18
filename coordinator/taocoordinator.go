package coordinator

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
	"tao.exchange.com/common"
	"tao.exchange.com/common/orm"
	pb "tao.exchange.com/grpc"
)

type server struct {
	pb.UnimplementedTaoCoordinatorSrvServer
}

func (s *server) ListShards(ctx context.Context, req *pb.CommonReq) (*pb.LockShardRsp, error) {
	tlkm := orm.TaoShardMarketMapper{}
	id := int32(0)
	rsp := new(pb.LockShardRsp)
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"
	rsp.MarketList = make(map[string]*pb.MarketShardDto)
	for {
		lst, err := tlkm.BatchSelect(id)
		if err != nil {
			rsp.Status = http.StatusInternalServerError
			rsp.Msg = err.Error()
			return rsp, err
		}
		if len(lst) == 0 {
			break
		}

		for _, v := range lst {
			msdto, ok := rsp.MarketList[v.ShardId]
			if !ok {
				msdto = new(pb.MarketShardDto)
				msdto.ShardId = v.ShardId
				msdto.MarketList = make([]string, 0)
				rsp.MarketList[v.ShardId] = msdto
			}
			msdto.MarketList = append(msdto.MarketList, v.MarketId)
		}
	}
	return rsp, nil
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

	// 启动DB 连接
	dsn := "postgres://taiji:TaiJiTrading2022@@localhost:5432/taoexchange?sslmode=disable"
	db := common.GetDbCon()
	db.Connect(dsn, &ctx)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		slog.Error("Failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterTaoCoordinatorSrvServer(s, &server{})

	go func() {
		slog.Info("Server listening at:", lis.Addr().String())
		if err := s.Serve(lis); err != nil {
			slog.Error("Failed to server:", err)
			return
		}
	}()

	startTaoCoordinatorRest()

	//Listen for the interrupt signal
	<-ctx.Done()
	stop()
	slog.Info("Shutdown Server ...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.GracefulStop()
	graceFulStop(&ctx)
	common.Get().Close()
	slog.Info("Server exist")
}
