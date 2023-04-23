package msgq

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"tao.exchange.com/common"
	pb "tao.exchange.com/grpc"
)

type TaoSubConsume func(req *pb.TaoMsgReq) error

type TaoMsgPub struct {
	timeout time.Duration
	pubId   string
}

type TaoMsgSub struct {
	timeout time.Duration
	topic   string
	groupId string
	subId   string
	subFun  TaoSubConsume
	stop    bool
}

func NewPub() (*TaoMsgPub, error) {
	pub := &TaoMsgPub{
		timeout: time.Duration(500 * time.Millisecond),
		pubId:   uuid.NewString(),
	}
	return pub, nil
}

func NewSub() (*TaoMsgSub, error) {
	sub := &TaoMsgSub{
		timeout: time.Duration(500 * time.Millisecond),
		subId:   uuid.NewString(),
	}
	return sub, nil
}

func (sub *TaoMsgSub) AutoClose() {
	slog.Info("topic:", sub.topic, " sub try to close")
	sub.stop = true
}

func (sub *TaoMsgSub) Listener(topic, groupId string, subFun TaoSubConsume) error {
	sub.topic = topic
	sub.subFun = subFun
	sub.groupId = groupId
	err := sub.tryToSub()
	if err != nil {
		return err
	}
	common.Get().Add(sub)
	return nil
}

func (sub *TaoMsgSub) tryToSub() error {
	go func() {
		for {
			slog.Info("start to get sub")
			tbc := sub.getSub()
			if tbc != nil {
				slog.Info("start to listen")
				sub.notify(tbc)
			}
			if sub.stop {
				break
			}
			time.Sleep(5 * time.Second)
		}
		slog.Info("toppic:", sub.topic, " listen is stopped")
	}()
	return nil
}

func (sub *TaoMsgSub) getSub() *pb.TaoBroker_SubClient {
	brk := GetBrkClient().GetGrpcBroker()
	if brk == nil {
		slog.Info("sub error: broker is not exist")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), sub.timeout)
	req := &pb.TaoSubReq{
		Topic:   sub.topic,
		GroupId: sub.groupId,
		SubId:   sub.subId,
	}
	defer cancel()
	tbc, err := (*brk).Sub(ctx, req)
	if err != nil {
		slog.Info("sub error:", err.Error())
		return nil
	}
	return &tbc
}

func (sub *TaoMsgSub) notify(tbc *pb.TaoBroker_SubClient) {
	for {
		if sub.stop {
			break
		}
		in, err := (*tbc).Recv()
		if err == io.EOF {
			slog.Info("toppic:", sub.topic, " recv eof")
			break
		}
		if err != nil {
			slog.Info("toppic:", sub.topic, " recv error:", err.Error())
			continue
		}
		err = sub.subFun(in)
		if err != nil {
			slog.Info("toppic:", sub.topic, " notify error:", err.Error())
			continue
		}
	}
	slog.Info("toppic:", sub.topic, " notify is stopped")
	err := (*tbc).CloseSend()
	if err != nil {
		return
	}
}

func (pub *TaoMsgPub) SendMsg(msg *pb.TaoMsgReq) error {
	brk := GetBrkClient().GetGrpcBroker()
	if brk == nil {
		slog.Warn("SendMsg error: broker is not exist")
		return errors.New("broker is not exist")
	}
	ctx, cancel := context.WithTimeout(context.Background(), pub.timeout)
	defer cancel()
	rsp, err := (*brk).Pub(ctx, msg)
	if err != nil {
		slog.Warn("SendMsg error:", err.Error())
		return err
	}

	if rsp.Code != http.StatusOK {
		slog.Warn("SendMsg error:", rsp.Msg)
		return errors.New(rsp.Msg)
	}

	return nil
}
