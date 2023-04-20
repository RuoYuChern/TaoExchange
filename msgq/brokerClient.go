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

type BrokerClient struct {
	common.AutoCloseable
	dsn    string
	stop   bool
	brk    pb.TaoBrokerClient
	brkCon *grpc.ClientConn

	cord      pb.TaoCoordinatorSrvClient
	cordCon   *grpc.ClientConn
	brkConDto *pb.ConnectorDto
}

var brkClient *BrokerClient
var once sync.Once

func GetBrkClient() *BrokerClient {
	once.Do(func() {
		brkClient = &BrokerClient{stop: false}
	})
	return brkClient
}

func (blk *BrokerClient) AutoClose() {
	slog.Info("broker client is close")
	blk.stop = true
	if blk.brkCon != nil {
		blk.brkCon.Close()
	}

	if blk.cordCon != nil {
		blk.cordCon.Close()
	}
}

func (blk *BrokerClient) GetBroker() *pb.TaoBrokerClient {
	if blk.brk == nil {
		slog.Error("broker is nil")
		return nil
	}
	return &blk.brk
}

func (blk *BrokerClient) Connect(dsn string) error {
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
		return errors.New(rsp.Msg)
	}

	dto, ok := rsp.ConnectorList[common.MsqQMaster]
	if !ok {
		slog.Warn("broker is not exist")
		return errors.New("broker is not exist")
	}
	blk.brkConDto = dto
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", dto.Ip, dto.Port))
	if err != nil {
		slog.Warn("Connect to broker error:", err.Error())
		return err
	}
	blk.brkCon = conn
	blk.brk = pb.NewTaoBrokerClient(conn)
	/**注册自己**/
	common.Get().Add(blk)
	/**刷新线程**/
	go func() {
		blk.checkUpdate()
	}()
	return nil
}

func (blk *BrokerClient) reset() {
	if blk.brkCon != nil {
		blk.brkCon.Close()
		blk.brkCon = nil
	}

	if blk.cordCon != nil {
		blk.cordCon.Close()
		blk.cordCon = nil
	}
}

func (blk *BrokerClient) checkUpdate() {
	for {
		if blk.stop {
			slog.Info("checkUpdate quit")
			break
		}
		time.Sleep(1500 * time.Millisecond)
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		req := &pb.CommonReq{MsgId: fmt.Sprintf("msgId-%d", time.Now().UnixMilli())}
		rsp, err := blk.cord.ListConnectorInfo(ctx, req)
		if err != nil {
			slog.Warn("Connect error:", err.Error())
			continue
		}

		if rsp.Status != http.StatusOK {
			slog.Warn("Connect error:", err.Error())
			continue
		}

		dto, ok := rsp.ConnectorList[common.MsqQMaster]
		if !ok {
			slog.Warn("broker is not exist")
			continue
		}

		if dto.AppId == blk.brkConDto.AppId {
			continue
		}

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", dto.Ip, dto.Port))
		if err != nil {
			slog.Warn("Connect to broker error:", err.Error())
		}

		blk.reset()
		blk.brkCon = conn
		blk.brk = pb.NewTaoBrokerClient(conn)
		slog.Info(fmt.Sprintf("Update broker ip: %s %s, port: %d, %d", blk.brkConDto.Ip, dto.Ip, blk.brkConDto.Port, dto.Port))
		blk.brkConDto = dto
	}
}
