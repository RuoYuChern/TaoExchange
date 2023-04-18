package coordinator

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/rs/cors"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"golang.org/x/exp/slog"
	"tao.exchange.com/common"
	"tao.exchange.com/common/orm"
)

func (tqc *TaoHttpQueryCache) get(key string) string {
	return tqc.qc.Get(key)
}

func bindJson(req *bunrouter.Request, obj any) error {
	err := binding.JSON.Bind(req.Request, obj)
	if err != nil {
		slog.Warn("bindJson error")
		return err
	}
	return nil
}

func (tqc *TaoHttpQueryCache) getInt(key string) (int32, bool) {
	v := tqc.get(key)
	if common.IsBlank(v) {
		return -1, false
	}
	n, err := strconv.ParseInt(v, 10, 32)
	if err == nil {
		return int32(n), true
	} else {
		return -1, false
	}
}

func addShardMarket(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("addShardMarket")
	rsp := coordinatorRsp{
		Status: http.StatusNotFound,
		Msg:    "Not find",
	}

	var taoReq taoAddShardMarkeReq
	if err := bindJson(&req, &taoReq); err != nil {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "Bad args"
		return bunrouter.JSON(w, &rsp)
	}
	if len(taoReq.Data) == 0 {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "Bad args"
		return bunrouter.JSON(w, &rsp)
	}

	tsmm := orm.TaoShardMarketMapper{}
	tsmList := make([]orm.TaoShardMarket, 0, len(taoReq.Data))
	for _, v := range taoReq.Data {
		tsm := orm.TaoShardMarket{}
		tsm.MarketId = v.MarketId
		tsm.ShardId = v.ShardId
		tsm.CreateTime = time.Now()
		tsmList = append(tsmList, tsm)
	}
	r := tsmm.BachInsert(&tsmList)
	if r < 0 {
		rsp.Status = http.StatusInternalServerError
		rsp.Msg = "Insert error"
		return bunrouter.JSON(w, &rsp)
	}

	rsp.Status = http.StatusOK
	rsp.Msg = "OK"
	return bunrouter.JSON(w, &rsp)
}

func listShardMarket(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("listShardMarket")
	queryCache := TaoHttpQueryCache{qc: req.URL.Query()}
	id, ok := queryCache.getInt("id")
	slog.Info("id:", id)
	rsp := coordinatorRsp{
		Status: http.StatusNotFound,
		Msg:    "Not find",
	}
	if !ok {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "Bad args"
		return bunrouter.JSON(w, &rsp)
	}

	tsmm := orm.TaoShardMarketMapper{}
	tsms, err := tsmm.BatchSelect(id)
	if err != nil {
		rsp.Status = http.StatusInternalServerError
		rsp.Msg = err.Error()
		return bunrouter.JSON(w, &rsp)
	}

	data := make([]*taoShardMarketDto, 0, len(tsms))

	for _, v := range tsms {
		slog.Info("mardId:{}, shardId:{}", v.MarketId, v.ShardId)
		tsmd := new(taoShardMarketDto)
		tsmd.MarketId = v.MarketId
		tsmd.ShardId = v.ShardId
		tsmd.UpdateTime = v.CreateTime.Local().Format(time.DateTime)
		data = append(data, tsmd)
	}
	slog.Info("Get count:", len(tsms), len(data))
	rsp.Status = http.StatusOK
	rsp.Msg = "OK"
	rsp.Data = &data
	return bunrouter.JSON(w, &rsp)
}

func listTaoLock(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("listTaoLock")
	rsp := coordinatorRsp{
		Status: http.StatusNotFound,
		Msg:    "Not find",
	}
	tlm := orm.TaoLockMapper{}
	lockList, err := tlm.BatchSelect()
	if err != nil {
		rsp.Status = http.StatusInternalServerError
		rsp.Msg = err.Error()
		return bunrouter.JSON(w, &rsp)
	}

	data := make([]*taoLockDto, 0)
	for _, v := range lockList {
		tlk := new(taoLockDto)
		tlk.ShardId = v.ShardId
		tlk.AppId = v.AppIP
		tlk.AppIp = v.AppIP
		tlk.AppPort = v.AppPort
		tlk.AppRole = v.AppRole
		tlk.AppStatus = v.AppStatus
		tlk.LockTime = v.LockTime.Local().Format(time.DateTime)
		data = append(data, tlk)
	}

	rsp.Status = http.StatusOK
	rsp.Msg = "OK"
	rsp.Data = data

	return bunrouter.JSON(w, &rsp)
}

func releaseTaoLock(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("releaseTaoLock")
	rsp := coordinatorRsp{
		Status: http.StatusNotFound,
		Msg:    "Not find",
	}
	var taoReq taoLockReq
	if err := bindJson(&req, &taoReq); err != nil {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "Bad args"
		return bunrouter.JSON(w, &rsp)
	}

	tlock := orm.TaoLock{
		ShardId:   taoReq.Data.ShardId,
		AppId:     taoReq.Data.AppId,
		AppStatus: orm.FREE_ST,
		LockTime:  time.Now(),
	}
	tlm := orm.TaoLockMapper{}
	r := tlm.ReleaseLock(&tlock)
	if r < 0 {
		rsp.Status = http.StatusInternalServerError
		rsp.Msg = "Error"
		return bunrouter.JSON(w, &rsp)
	}

	rsp.Status = http.StatusOK
	rsp.Msg = fmt.Sprintf("RowsAffected: %d", r)
	return bunrouter.JSON(w, &rsp)
}

func addTaoMarket(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("addTaoMarket")
	rsp := coordinatorRsp{
		Status: http.StatusOK,
		Msg:    "OK",
	}

	var taoReq taoMarketReq
	if err := bindJson(&req, &taoReq); err != nil {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "Bad args"
		return bunrouter.JSON(w, &rsp)
	}
	tmList := make([]orm.TaoMarket, 0, len(taoReq.Data))
	for _, v := range taoReq.Data {
		tm := orm.TaoMarket{
			Market:       v.Market,
			Base:         v.Base,
			Pair:         v.Pair,
			MarketStatus: v.MarketStatus,
			BasePrec:     v.BasePrec,
			PairPrec:     v.PairPrec,
			FeePrec:      v.FeePrec,
			MinAmount:    v.MinAmount,
			MinBase:      v.MinBase,
			CreateTime:   time.Now(),
		}
		tmList = append(tmList, tm)
	}
	tmm := orm.TaoMarketMapper{}
	_, err := tmm.BachInsert(tmList)
	if err != nil {
		rsp.Status = http.StatusInternalServerError
		rsp.Msg = err.Error()
	}
	return bunrouter.JSON(w, &rsp)
}

func listTaoMarket(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("listTaoMarket")
	queryCache := TaoHttpQueryCache{qc: req.URL.Query()}
	id, ok := queryCache.getInt("id")
	rsp := coordinatorRsp{
		Status: http.StatusOK,
		Msg:    "OK",
	}

	if !ok {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "Bad args"
		return bunrouter.JSON(w, &rsp)
	}

	tmm := orm.TaoMarketMapper{}
	tms, err := tmm.BatchSelect(id)
	if err != nil {
		rsp.Status = http.StatusInternalServerError
		rsp.Msg = err.Error()
		return bunrouter.JSON(w, &rsp)
	}

	data := make([]*taoMarketDto, 0)
	for _, v := range tms {
		tmdto := new(taoMarketDto)
		tmdto.Market = v.Market
		tmdto.Base = v.Base
		tmdto.Pair = v.Pair
		tmdto.MarketStatus = v.MarketStatus
		tmdto.BasePrec = v.BasePrec
		tmdto.PairPrec = v.PairPrec
		tmdto.FeePrec = v.FeePrec
		tmdto.MinAmount = v.MinAmount
		tmdto.MinBase = v.MinBase
		data = append(data, tmdto)
	}
	rsp.Data = data
	return bunrouter.JSON(w, rsp)
}

var httpSrv *http.Server

func startTaoCoordinatorRest() {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware(reqlog.FromEnv("bundebug"))),
	)

	router.POST("/cmd/market/add-shard-market", addShardMarket)
	router.POST("/cmd/market/add-tao-market", addTaoMarket)
	router.POST("/cmd/lock/release-tao-lock", releaseTaoLock)
	router.GET("/q/market/list-shard-market", listShardMarket)
	router.GET("/q/lock/list-tao-lock", listTaoLock)
	router.GET("/q/market/list-tao-market", listTaoMarket)

	httpLn, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	handler := http.Handler(router)
	handler = cors.Default().Handler(handler)
	httpSrv = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      handler,
	}

	slog.Info("Listen on:8090 ...")
	go func() {
		if err := httpSrv.Serve(httpLn); err != nil {
			slog.Error("listen error:", err.Error())
		}
	}()
}

func graceFulStop(ctx *context.Context) {
	if err := httpSrv.Shutdown(*ctx); err != nil {
		slog.Warn("Shutdown error:", err)
	}
}
