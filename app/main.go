package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"lintangbs.org/lintang/template/app/di"
	"lintangbs.org/lintang/template/config"
	"lintangbs.org/lintang/template/internal/rest/middleware"
	"lintangbs.org/lintang/template/pkg/httpserver"
	"lintangbs.org/lintang/template/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func init() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	//
	if err := middleware.InitLogger(cfg.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

}

func main() {
	cfg, err := config.NewConfig()

	// init logger

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	handler := gin.New()
	httpServer := httpserver.New(handler, httpserver.Port("5033"))

	// init app
	pg := di.InitApp(cfg, handler)

	// Waiting signal

	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		zap.L().Error("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		zap.L().Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}

	// Shutdown
	err = httpServer.Shutdown()
	postgres.ClosePostgres(pg.Pool)
	if err != nil {
		zap.L().Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
