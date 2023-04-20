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

func (s *server) ListShards(ctx context.Context, req *pb.CommonReq) (*pb.ShardListRsp, error) {
	tlkm := orm.TaoShardMarketMapper{}
	id := int32(0)
	rsp := new(pb.ShardListRsp)
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
func (s *server) LockShard(ctx context.Context, req *pb.ShardReq) (*pb.CommonRsp, error) {
	rsp := new(pb.CommonRsp)
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"
	if common.IsBlank(req.AppId) || common.IsBlank(req.Ip) || common.IsBlank(req.ShardId) {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "BadRequest"
		return rsp, nil
	}
	tlkm := orm.TaoLockMapper{}
	tlk := orm.TaoLock{
		ShardId:   req.ShardId,
		AppId:     req.AppId,
		AppIP:     req.Ip,
		AppRole:   pb.ShardRole_name[int32(req.Role)],
		AppPort:   req.Port,
		AppStatus: orm.LOCK_ST,
		LockTime:  time.Now(),
	}

	r := tlkm.Insert(&tlk)
	if r < 0 {
		rsp.Status = http.StatusInternalServerError
		rsp.Msg = "Internal Server Error"
		return rsp, nil
	}

	if r == 0 {
		rsp.Status = http.StatusNotAcceptable
		rsp.Msg = "Not Acceptable"
		return rsp, nil
	}
	getSs().lockShard(&tlk)
	return rsp, nil
}
func (s *server) UnlockShard(ctx context.Context, req *pb.ShardReq) (*pb.CommonRsp, error) {
	rsp := new(pb.CommonRsp)
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"

	if common.IsBlank(req.AppId) || common.IsBlank(req.Ip) || common.IsBlank(req.ShardId) {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "BadRequest"
		return rsp, nil
	}

	tlk := orm.TaoLock{
		ShardId:   req.ShardId,
		AppId:     req.AppId,
		AppIP:     req.Ip,
		AppRole:   pb.ShardRole_name[int32(req.Role)],
		AppPort:   req.Port,
		AppStatus: orm.FREE_ST,
		LockTime:  time.Now(),
	}

	r := getSs().checkLock(tlk.ShardId)
	if !r {
		rsp.Status = http.StatusNotAcceptable
		rsp.Msg = "Not Acceptable"
		return rsp, nil
	}

	tlkm := orm.TaoLockMapper{}
	r2 := tlkm.ReleaseLock(&tlk)
	if r2 > 0 {
		getSs().unlock(&tlk)
	} else {
		rsp.Status = http.StatusNotAcceptable
		rsp.Msg = "Not Acceptable"
	}
	return rsp, nil
}
func (s *server) ListConnectorInfo(ctx context.Context, req *pb.CommonReq) (*pb.ListConnectorRsp, error) {
	rsp := new(pb.ListConnectorRsp)
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"

	conMap := getSs().getConnectorList()
	if len(conMap) == 0 {
		rsp.Status = http.StatusNotFound
		rsp.Msg = "Not Found"
		return rsp, nil

	}
	rsp.ConnectorList = conMap
	return rsp, nil
}
func (s *server) KeepLive(ctx context.Context, req *pb.ShardReq) (*pb.CommonRsp, error) {
	rsp := new(pb.CommonRsp)
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"
	r := getSs().keepLive(req)
	if !r {
		rsp.Status = http.StatusNotAcceptable
		rsp.Msg = "Not Acceptable"
	}
	return rsp, nil
}

var (
	port = flag.Int("port", 59081, "The server port")
)

func StartTaoCoordinator() {
	// create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 启动DB 连接
	taoConf := common.TaoConf{}
	taoConf.LoadTaoConf("../tao_conf.yaml")
	db := common.GetDbCon()
	db.Connect(taoConf.DbDns, &ctx)
	// load lock info
	getSs().int()

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
