package main

import (
	"context"
	_ "expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"starfission_go_api/global"
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/routers"
	"starfission_go_api/internal/task"
	"starfission_go_api/pkg/logger"
	"starfission_go_api/pkg/rabbitmq"
	"starfission_go_api/pkg/redis"
	"starfission_go_api/pkg/setting"
	"starfission_go_api/pkg/tracer"
	"starfission_go_api/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validatorV10 "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port      string
	runMode   string
	config    string
	Version   string
	isVersion bool
	GitTag    = "2000.01.01.release"
	BuildTime = "2000-01-01T00:00:00+0800"
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}

	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupRedis()
	if err != nil {
		log.Fatalf("init.setupRedis err: %v", err)
	}
	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupValidator err: %v", err)
	}

	err = setupRabbitMQ()
	if err != nil {
		log.Fatalf("init.setupRabbitMQ err: %v", err)
	}

	//err = setupTracer()
	//if err != nil {
	//	log.Fatalf("init.setupTracer err: %v", err)
	//}
}

func main() {

	if isVersion {
		fmt.Println("Version: " + Version)
		fmt.Println("Git Tag: " + GitTag)
		fmt.Println("Build Time: " + BuildTime)
		return
	}

	fmt.Println("Config Path : " + global.CONFIG)

	gin.SetMode(global.ServerSetting.RunMode)

	task.Launcher()

	validator.RegisterCustom()

	go routers.MetricsSrv()

	go routers.NewOpenServiceRouter()

	log.Println("Http API Service ListenAndServe On: ", global.ServerSetting.HttpPort, "\n")
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "????????????")
	flag.StringVar(&runMode, "mode", "", "????????????")
	flag.StringVar(&config, "config", "configs/config.yaml", "????????????????????????????????????")
	flag.BoolVar(&isVersion, "version", false, "????????????")
	v := flag.Bool("v", false, "??????????????????")
	flag.Parse()
	if *v {
		isVersion = true
	}
	return nil
}

func setupSetting() error {
	global.CONFIG = config

	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("RabbitMQ", &global.RabbitMQSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Security", &global.SecuritySetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("WechatJSAPIPay", &global.WechatJSAPIPaySetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("WechatH5Pay", &global.WechatH5PaySetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("AlipayH5Pay", &global.AlipayH5PaySetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("OSS", &global.OSSSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	})

	global.Logger = logger.NewLogger(mw, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupRedis() error {
	var err error
	global.Redis, err = redis.NewConn(global.RedisSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupRabbitMQ() error {

	var err error
	global.RabbitMQ, err = rabbitmq.NewRdabbitMQCon(global.RabbitMQSetting.Dsn)
	if err != nil {
		return err
	}
	return nil
}

func setupValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator
	uni := ut.New(en.New(), en.New(), zh.New())
	v, ok := binding.Validator.Engine().(*validatorV10.Validate)
	if ok {
		zhTran, _ := uni.GetTranslator("zh")
		enTran, _ := uni.GetTranslator("en")
		err := zh_translations.RegisterDefaultTranslations(v, zhTran)
		if err != nil {
			return err
		}
		err = en_translations.RegisterDefaultTranslations(v, enTran)
		if err != nil {
			return err
		}
	}

	global.Ut = uni

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
