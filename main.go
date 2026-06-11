package main

import (
	"app/components/database"
	"app/components/idgen"
	"app/components/logger"
	"app/config"
	"app/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	logManager := logger.NewLoggerManager(cfg.Log)

	db, err := database.Connect(cfg.DB)
	if err != nil {
		panic(err)
	}

	// 初始化Id生成器
	gen := idgen.NewSnowflake(1)

	// gin 创建服务 并优雅的关闭
	engine := router.NewRouter(db, logManager, gen)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      engine.Handler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logManager.App.Sugar().Errorf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logManager.App.Sugar().Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logManager.App.Sugar().Errorf("Server Shutdown: %s", err)
	}

	// 关闭数据库连接
	if err := database.Close(db); err != nil {
		logManager.App.Sugar().Errorf("Database Close: %s", err)
	}

	logManager.App.Sugar().Infof("Server exiting")
}
