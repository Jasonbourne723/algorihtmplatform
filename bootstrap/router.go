package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"algorithmplatform/app/hubs"
	docs "algorithmplatform/docs"
	"algorithmplatform/global"
	"algorithmplatform/routes"
	"algorithmplatform/signalr"

	kitlog "github.com/go-kit/log"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 注册api路由
func setupRouter() *gin.Engine {
	if global.App.Config.App.Env == "produce" {
		//gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	//router.Use(middleware.Cors())
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)
	return router
}

// 注册swagger
func setSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// 注册signalr
func setSignalr(ctx context.Context) {

	hub := hubs.AlgorithmHub
	server, _ := signalr.NewServer(ctx,
		signalr.UseHub(hub),
		signalr.AllowOriginPatterns([]string{"*"}),
		signalr.KeepAliveInterval(2*time.Second),
		signalr.Logger(kitlog.NewLogfmtLogger(os.Stdout), false))

	router := http.NewServeMux()
	server.MapHTTP(signalr.WithHTTPServeMux(router), "/hubs/algorithmhub")
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("signalr error :" + err.(error).Error())
			}
		}()
		if err := http.ListenAndServe(fmt.Sprintf(":%d", global.App.Config.App.SignalrPort), router); err != nil {
			fmt.Println("signalr err: " + err.Error())
		}
	}()
}

// 启动http listener
func RunServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r := setupRouter()
	setSwagger(r)

	go func() {
		if err := r.Run(fmt.Sprintf(":%d", global.App.Config.App.Port)); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()

	setSignalr(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Server exiting")
}
