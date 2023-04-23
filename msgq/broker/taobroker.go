package broker

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tao.exchange.com/common"
	pb "tao.exchange.com/grpc"
)

type taoSubCon struct {
	subCon *grpc.ClientConn
	subSvr pb.TaoBroker_SubServer
}

type taoSubMeta struct {
	topic   string
	groupId string
	subId   string
	offset  int64
}

type taoBroker struct {
	common.TaoCloseable
	pb.UnimplementedTaoBrokerServer
	appId    string
	cord     pb.TaoCoordinatorSrvClient
	cordCon  *grpc.ClientConn
	isMaster bool
	stop     bool
}

func (brk *taoBroker) start(conf *common.TaoAppConf) {
	taoConf := *conf
	brk.stop = false
	brk.isMaster = false
	brk.appId = uuid.NewString()
	err := brk.getCoordinator(taoConf.Tao.CoordinatorUrl)
	if err != nil {
		slog.Error("get coordinator failed")
		panic(err.Error())
	}
	go func() {
		brk.watch(conf.Tao.BrokerPort)
	}()
}

func (brk *taoBroker) AutoClose() {
	brk.stop = false
	if brk.cordCon != nil {
		err := brk.cordCon.Close()
		if err != nil {
			return
		}
	}
}

func (brk *taoBroker) watch(port int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	req := &pb.ShardReq{ShardId: common.MsqQMaster, AppId: brk.appId, Ip: "", Port: port, Role: pb.ShardRole_SR_MQ}
	_, err := brk.cord.LockShard(ctx, req)
	if err != nil {
		return
	}
}

func (brk *taoBroker) getCoordinator(dsn string) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(dsn, opts...)
	if err != nil {
		slog.Warn("Connect to coordinator error:", err.Error())
		return err
	}
	brk.cordCon = conn
	brk.cord = pb.NewTaoCoordinatorSrvClient(conn)
	return nil
}

func (brk *taoBroker) Cmd(context.Context, *pb.TaoMsgCmdReq) (*pb.TaoMsgCmdRsp, error) {
	return nil, nil
}
func (brk *taoBroker) Pub(context.Context, *pb.TaoMsgReq) (*pb.TaoMsgRsp, error) {
	return nil, nil
}
func (brk *taoBroker) Sync(context.Context, *pb.TaoMsgReq) (*pb.TaoMsgRsp, error) {
	return nil, nil
}
func (brk *taoBroker) Sub(*pb.TaoSubReq, pb.TaoBroker_SubServer) error {
	return nil
}

func StartBroker() {
	// create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	common.Get()
	taoConf := common.TaoAppConf{}
	taoConf.LoadTaoConf("../tao_conf.yaml")
	tbk := taoBroker{}
	tbk.start(&taoConf)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", taoConf.Tao.BrokerPort))
	if err != nil {
		slog.Error("Failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterTaoBrokerServer(s, &tbk)
	go func() {
		slog.Info("Server listening at:", lis.Addr().String())
		if err := s.Serve(lis); err != nil {
			slog.Error("Failed to server:", err)
			return
		}
	}()
	<-ctx.Done()
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.GracefulStop()

	stop()
	/**资源释放**/
	common.Get().Close()
	slog.Info("Server exist")
}
