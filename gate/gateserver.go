package gate

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"tao.exchange.com/common"
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
	req := new(QueryReq)
	req.Market = c.Query("market")
	req.OrderId = c.Query("orderId")
	req.UserId = c.Query("userId")
	if common.IsBlank(req.Market) && common.IsBlank(req.UserId) {
		slog.Info("market or userId is empty")
		rsp := makeQueryRsp(http.StatusBadRequest, "Bad args")
		c.IndentedJSON(http.StatusBadRequest, rsp)
		return
	}

	rsp := makeQueryRsp(http.StatusOK, "OK")
	rsp.Orders = make([]OrderDto, 0, 1)
	order := new(OrderDto)
	order.Id = "123456"
	order.Commond = CMD_PLACE
	order.Price = 12345
	rsp.Orders = append(rsp.Orders, *order)
	c.IndentedJSON(http.StatusOK, rsp)
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
}

func StartGateService() {
	// create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.New()
	router.Use(gateFilter())
	router.Use(gin.Recovery())
	makeRoute(router)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
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
	slog.Info("Server exist")
}
