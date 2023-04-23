package msgq

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tao.exchange.com/common"
	pb "tao.exchange.com/grpc"
)

type TaoMsgQClient struct {
	common.TaoCloseable
	dsn    string
	stop   bool
	brk    pb.TaoBrokerClient
	brkCon *grpc.ClientConn

	cord      pb.TaoCoordinatorSrvClient
	cordCon   *grpc.ClientConn
	brkConDto *pb.ConnectorDto
}

var brkClient *TaoMsgQClient
var once sync.Once

func GetBrkClient() *TaoMsgQClient {
	once.Do(func() {
		brkClient = &TaoMsgQClient{stop: false}
	})
	return brkClient
}

func (blk *TaoMsgQClient) AutoClose() {
	slog.Info("broker client is close")
	blk.stop = true
	if blk.brkCon != nil {
		err := blk.brkCon.Close()
		if err != nil {
			return
		}
	}

	if blk.cordCon != nil {
		err := blk.cordCon.Close()
		if err != nil {
			return
		}
	}
}

func (blk *TaoMsgQClient) GetGrpcBroker() *pb.TaoBrokerClient {
	if blk.brk == nil {
		slog.Error("broker is nil")
		return nil
	}
	return &blk.brk
}

func (blk *TaoMsgQClient) Connect(dsn string) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	blk.dsn = dsn
	conn, err := grpc.Dial(dsn, opts...)
	if err != nil {
		slog.Warn("Connect to coordinator error:", err.Error())
		return err
	}
	blk.cord = pb.NewTaoCoordinatorSrvClient(conn)
	blk.cordCon = conn
	err = blk.setGrpcBrokerClient()
	if err != nil {
		slog.Warn("Get broker error:", err.Error())
		return err
	}
	/**注册自己**/
	common.Get().Add(blk)
	/**刷新线程**/
	go func() {
		blk.checkUpdate()
	}()
	return nil
}

func (blk *TaoMsgQClient) reset() {
	if blk.brkCon != nil {
		err := blk.brkCon.Close()
		if err != nil {
			return
		}
		blk.brkCon = nil
	}

	if blk.cordCon != nil {
		err := blk.cordCon.Close()
		if err != nil {
			return
		}
		blk.cordCon = nil
	}
}

func (blk *TaoMsgQClient) setGrpcBrokerClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	req := &pb.CommonReq{MsgId: fmt.Sprintf("msgId-%d", time.Now().UnixMilli())}
	rsp, err := blk.cord.ListConnectorInfo(ctx, req)
	if err != nil {
		slog.Warn("Connect error:", err.Error())
		return err
	}

	if rsp.Status != http.StatusOK {
		slog.Warn("Connect error:", err.Error())
		return err
	}

	dto, ok := rsp.ConnectorList[common.MsqQMaster]
	if !ok {
		slog.Warn("broker is not exist")
		return errors.New("broker is not exist")
	}

	if (blk.brkConDto != nil) && (dto.AppId == blk.brkConDto.AppId) {
		return nil
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", dto.Ip, dto.Port))
	if err != nil {
		slog.Warn("Connect to broker error:", err.Error())
	}

	blk.reset()
	blk.brkCon = conn
	blk.brk = pb.NewTaoBrokerClient(conn)
	if blk.brkConDto != nil {
		slog.Info(fmt.Sprintf("Update broker ip: %s %s, port: %d, %d", blk.brkConDto.Ip, dto.Ip, blk.brkConDto.Port, dto.Port))
	}
	blk.brkConDto = dto
	return nil
}

func (blk *TaoMsgQClient) checkUpdate() {
	for {
		if blk.stop {
			slog.Info("checkUpdate quit")
			break
		}
		time.Sleep(1500 * time.Millisecond)
		err := blk.setGrpcBrokerClient()
		if err != nil {
			return
		}
	}
}
