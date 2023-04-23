package coordinator

import (
	"context"
	"net/http"
	"time"

	"tao.exchange.com/common"
	pb "tao.exchange.com/grpc"
	"tao.exchange.com/infra/orm"
)

type TaoCoServer struct {
	pb.UnimplementedTaoCoordinatorSrvServer
}

func (s *TaoCoServer) ListShards(ctx context.Context, req *pb.CommonReq) (*pb.ShardListRsp, error) {
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
func (s *TaoCoServer) LockShard(ctx context.Context, req *pb.ShardReq) (*pb.CommonRsp, error) {
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
	GetSs().lockShard(&tlk)
	return rsp, nil
}
func (s *TaoCoServer) UnlockShard(ctx context.Context, req *pb.ShardReq) (*pb.CommonRsp, error) {
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

	r := GetSs().checkLock(tlk.ShardId)
	if !r {
		rsp.Status = http.StatusNotAcceptable
		rsp.Msg = "Not Acceptable"
		return rsp, nil
	}

	tlkm := orm.TaoLockMapper{}
	r2 := tlkm.ReleaseLock(&tlk)
	if r2 > 0 {
		GetSs().unlock(&tlk)
	} else {
		rsp.Status = http.StatusNotAcceptable
		rsp.Msg = "Not Acceptable"
	}
	return rsp, nil
}
func (s *TaoCoServer) ListConnectorInfo(ctx context.Context, req *pb.CommonReq) (*pb.ListConnectorRsp, error) {
	rsp := new(pb.ListConnectorRsp)
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"

	conMap := GetSs().getConnectorList()
	if len(conMap) == 0 {
		rsp.Status = http.StatusNotFound
		rsp.Msg = "Not Found"
		return rsp, nil

	}
	rsp.ConnectorList = conMap
	return rsp, nil
}

func (s *TaoCoServer) KeepLive(ctx context.Context, req *pb.ShardReq) (*pb.CommonRsp, error) {
	rsp := new(pb.CommonRsp)
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"
	r := GetSs().keepLive(req)
	if !r {
		rsp.Status = http.StatusNotAcceptable
		rsp.Msg = "Not Acceptable"
	}
	return rsp, nil
}

func (s *TaoCoServer) ListTaoMarket(ctx context.Context, req *pb.ListTaoMarketReq) (*pb.ListTaoMarketRsp, error) {
	rsp := new(pb.ListTaoMarketRsp)
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"
	return rsp, nil
}
