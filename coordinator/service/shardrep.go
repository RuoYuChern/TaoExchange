package coordinator

import (
	"container/list"
	"sync"
	"time"

	"golang.org/x/exp/slog"
	"tao.exchange.com/common"
	pb "tao.exchange.com/grpc"
	"tao.exchange.com/infra/orm"
)

type shardStatusInfo struct {
	shard        orm.TaoLock
	timeOut      int
	lastKeepTime time.Time
}

type ShardService struct {
	shardMap  map[string]*list.Element
	shardList *list.List
	lock      sync.Mutex
	stop      bool
}

var ss *ShardService
var once sync.Once

func GetSs() *ShardService {
	once.Do(func() {
		ss = &ShardService{
			shardMap:  make(map[string]*list.Element),
			shardList: list.New(),
			lock:      sync.Mutex{},
			stop:      false,
		}
	})

	return ss
}

func (s *ShardService) Int() {
	common.Get().Add(s)
	defer s.lock.Unlock()
	s.lock.Lock()
	tlm := orm.TaoLockMapper{}
	lst, err := tlm.BatchSelect()
	if err != nil {
		slog.Error("BatchSelect excpetion:", err.Error())
		panic(err.Error())
	}

	for _, shard := range lst {
		ssi := new(shardStatusInfo)
		ssi.shard = shard
		ssi.lastKeepTime = time.Now()
		ssi.timeOut = 0
		element := s.shardList.PushBack(ssi)
		s.shardMap[shard.ShardId] = element
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			if s.stop {
				slog.Info("shardService quit")
				return
			}
			s.checkTimeOut()
		}
	}()

}

func (s *ShardService) AutoClose() {
	slog.Info("shardService close")
	s.stop = true
}

func (s *ShardService) checkTimeOut() {
	defer s.lock.Unlock()
	s.lock.Lock()
	now := time.Now()
	level := 3 * time.Second
	var next *list.Element
	for c := s.shardList.Front(); c != nil; c = next {
		ssi := c.Value.(*shardStatusInfo)
		next = c.Next()
		diff := now.Sub(ssi.lastKeepTime)
		if diff > level {
			ssi.timeOut = ssi.timeOut + 1
			ssi.lastKeepTime = time.Now()
			if ssi.timeOut >= 3 {
				slog.Info("ShardId:", ssi.shard.ShardId, " ip:", ssi.shard.AppIP, " is time out")
				delete(s.shardMap, ssi.shard.ShardId)
				s.shardList.Remove(c)
			}
			continue
		} else {
			ssi.timeOut = 0
			ssi.lastKeepTime = time.Now()
		}
	}
}

func (s *ShardService) lockShard(shard *orm.TaoLock) {
	defer s.lock.Unlock()
	s.lock.Lock()
	_, ok := s.shardMap[shard.ShardId]
	if ok {
		slog.Warn("ShardId:", shard.ShardId, " is exist")
		return
	}

	ssi := new(shardStatusInfo)
	ssi.shard = *shard
	ssi.lastKeepTime = time.Now()
	ssi.timeOut = 0
	element := s.shardList.PushBack(ssi)
	s.shardMap[shard.ShardId] = element
}

func (s *ShardService) checkLock(shardId string) bool {
	s.lock.Lock()
	_, ok := s.shardMap[shardId]
	s.lock.Unlock()
	return ok
}

func (s *ShardService) unlock(shard *orm.TaoLock) int {
	defer s.lock.Unlock()
	s.lock.Lock()
	element, ok := s.shardMap[shard.ShardId]
	if ok {
		slog.Warn("ShardId:", shard.ShardId, " is exist")
		return -1
	}
	delete(s.shardMap, shard.ShardId)
	s.shardList.Remove(element)
	return 0
}

func (s *ShardService) getConnectorList() map[string]*pb.ConnectorDto {
	defer s.lock.Unlock()
	conMap := make(map[string]*pb.ConnectorDto)
	s.lock.Lock()
	for c := s.shardList.Front(); c != nil; c = c.Next() {
		ssi := c.Value.(*shardStatusInfo)
		dto := new(pb.ConnectorDto)
		dto.ShardId = ssi.shard.ShardId
		dto.AppId = ssi.shard.AppId
		dto.Ip = ssi.shard.AppIP
		dto.Port = ssi.shard.AppPort
		dto.Status = ssi.shard.AppStatus
		dto.Role = pb.ShardRole(pb.ShardRole_value[ssi.shard.AppRole])
		conMap[ssi.shard.ShardId] = dto
	}
	return conMap
}

func (s *ShardService) keepLive(req *pb.ShardReq) bool {
	defer s.lock.Unlock()
	s.lock.Lock()
	element, ok := s.shardMap[req.ShardId]
	if !ok {
		slog.Warn("ShardId:", req.ShardId, " is exist")
		return false
	}

	ssi := element.Value.(*shardStatusInfo)
	if ssi.shard.AppId != req.AppId {
		slog.Info("AppId:", ssi.shard.AppId, " is not equal to:", req.ShardId)
		return false
	}
	ssi.lastKeepTime = time.Now()
	ssi.timeOut = 0
	return true
}
