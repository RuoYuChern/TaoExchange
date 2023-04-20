package msgq

import (
	"time"

	pb "tao.exchange.com/grpc"
)

type TaoMsgPub struct {
	timeout time.Duration
	retries int
}

func (pub *TaoMsgPub) Connect() error {
	var err error
	return err
}

func NewPub() (*TaoMsgPub, error) {
	pub := &TaoMsgPub{
		timeout: time.Duration(3000 * time.Millisecond),
		retries: 3,
	}
	return pub, nil
}

func (pub *TaoMsgPub) Sendmsg(msg *pb.TaoMsgReq) error {
	return nil
}
