package gate

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"tao.exchange.com/common"
	"tao.exchange.com/facade"
)

func gateFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()
		slog.Info("url stat:", slog.Group("perf", slog.String("url", c.FullPath()), slog.Int64("time used", latency.Milliseconds()),
			slog.Int("status", status)))
	}
}

func queryUserMarketOrder(c *gin.Context) {
	req := new(facade.QueryReq)
	req.Market = c.Query("market")
	req.OrderId = c.Query("orderId")
	req.UserId = c.Query("userId")
	if common.IsBlank(req.Market) && common.IsBlank(req.UserId) {
		slog.Info("market or userId is empty")
		rsp := facade.MakeQueryRsp(http.StatusBadRequest, "Bad args")
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	market := getMarket()
	rsp := market.query(req)
	c.JSON(http.StatusOK, rsp)
	slog.Info("OK")
}

func placeMarketOrder(c *gin.Context) {
	var order facade.OrderReq
	if err := c.BindJSON(&order); err != nil {
		rsp := facade.MakeOrderResp(http.StatusBadRequest, "args error", "", "")
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	market := getMarket()
	rsp := market.placeOrder(&order)
	c.JSON(http.StatusOK, rsp)
	slog.Info("OK")
}

func home(c *gin.Context) {
	c.String(http.StatusForbidden, "You does not have right")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}

func makeRoute(r *gin.Engine) {
	r.GET("/", home)
	r.GET("/hello", hello)
	r.GET("/order/query-user-market", queryUserMarketOrder)
	r.POST("/order/cmd/place-market-order", placeMarketOrder)
}

func StartGateService() {
	// create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.New()
	router.Use(gateFilter())
	router.Use(gin.Recovery())
	makeRoute(router)
	taoConf := common.TaoAppConf{}
	taoConf.LoadTaoConf("../tao_conf.yaml")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", taoConf.Tao.GateRestPort),
		Handler: router,
	}

	// 进行grpc 连接
	market := getMarket()
	if err := market.connect(); err != nil {
		slog.Error("connect to exchange error:", err)
		return
	}

	// Initializing the server
	go func() {
		/**进行连接**/
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen error:", err)
		}
	}()

	//Listen for the interrupt signal
	<-ctx.Done()

	stop()
	slog.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Warn("Server shutdown:", err)
	}
	market.shutdown()
	slog.Info("Server exist")
}
