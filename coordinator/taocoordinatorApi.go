package coordinator

import (
	"context"
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
	rsp.Data = taoReq.Data
	return bunrouter.JSON(w, &rsp)
}

func listShardMarket(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("listShardMarket")
	queryCache := TaoHttpQueryCache{qc: req.URL.Query()}
	offset, ok := queryCache.getInt("offset")
	slog.Info("offset:", offset)
	rsp := coordinatorRsp{
		Status: http.StatusNotFound,
		Msg:    "Not find",
	}

	if !ok {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "Bad args"
		return bunrouter.JSON(w, &rsp)
	}

	return bunrouter.JSON(w, &rsp)
}

func listTaoLock(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("listTaoLock")
	rsp := coordinatorRsp{
		Status: http.StatusNotFound,
		Msg:    "Not find",
	}
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

	return bunrouter.JSON(w, &rsp)
}

func addTaoMarket(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("addTaoMarket")
	rsp := coordinatorRsp{
		Status: http.StatusNotFound,
		Msg:    "Not find",
	}

	var taoReq taoMarketReq
	if err := bindJson(&req, &taoReq); err != nil {
		rsp.Status = http.StatusBadRequest
		rsp.Msg = "Bad args"
		return bunrouter.JSON(w, &rsp)
	}

	return bunrouter.JSON(w, &rsp)
}

func listTaoMarket(w http.ResponseWriter, req bunrouter.Request) error {
	slog.Info("listTaoMarket")
	rsp := coordinatorRsp{
		Status: http.StatusNotFound,
		Msg:    "Not find",
	}
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
			slog.Error("listen error:", err)
		}
	}()
}

func graceFulStop(ctx *context.Context) {
	if err := httpSrv.Shutdown(*ctx); err != nil {
		slog.Warn("Shutdown error:", err)
	}
}
