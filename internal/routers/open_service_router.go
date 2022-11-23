package routers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"starfission_go_api/global"
	"starfission_go_api/internal/middleware"
	"starfission_go_api/internal/routers/api"
	v1 "starfission_go_api/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 开放服务
func NewOpenServiceRouter() {
	defer func() {
		if err := recover(); err != nil {
			global.Logger.DErrorf("OpenService panic err: %v", err)
		}
	}()

	r := gin.New()
	r.Use(middleware.AppInfo())
	r.Use(middleware.DisEncrypt())
	r.Use(middleware.Sign())

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	//对404 的处理
	r.NoRoute(middleware.NoFound())

	//	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.Use(middleware.Cors())

	r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	APIV1 := r.Group("/api/v1")

	openService := v1.NewOpenService()
	APIV1.Use()
	{
		APIV1.POST("/sync/member", openService.SyncAppMember)
		APIV1.POST("/order/create", openService.CreateAppOrder)

	}

	log.Println("OpenService ListenAndServe On: ", global.ServerSetting.OpenServiceListen, "\n")
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.OpenServiceListen,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("OpenService ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("OpenService Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("OpenService forced to shutdown:", err)
	}
	log.Println("OpenService exiting")
}
