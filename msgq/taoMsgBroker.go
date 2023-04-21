package msgq

import (
	"context"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"tao.exchange.com/common"
	pb "tao.exchange.com/grpc"
)

type taoBroker struct {
	common.AutoCloseable
	pb.UnimplementedTaoBrokerServer
	cord    pb.TaoCoordinatorSrvClient
	cordCon *grpc.ClientConn
}

func (brk *taoBroker) start(){

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
	taoConf := common.TaoConf{}
	taoConf.LoadTaoConf("../tao_conf.yaml")
	tbk := taoBroker{}
	tbk.start()
	
	<-ctx.Done()
	stop()
	/**资源释放**/
	common.Get().Close()
}
