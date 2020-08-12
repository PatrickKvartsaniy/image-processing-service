package main

import (
	"context"
	"fmt"
	"github.com/PatrickKvartsaniy/image-processing-service/config"
	"github.com/PatrickKvartsaniy/image-processing-service/graph"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/resolver"
	"github.com/PatrickKvartsaniy/image-processing-service/health"
	"github.com/PatrickKvartsaniy/image-processing-service/processor"
	"github.com/PatrickKvartsaniy/image-processing-service/repository/mongo"
	"github.com/PatrickKvartsaniy/image-processing-service/storage"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	cfg := config.ReadOS()
	initLogger(cfg.LogLevel, cfg.PrettyLogOutput)

	ctx, cancel := context.WithCancel(context.Background())
	setupGracefulShutdown(cancel)
	if err := run(ctx, cfg); err != nil {
		cancel()
		logrus.WithError(err).Fatal("closing service with error")
	}
}

func run(ctx context.Context, cfg config.Config) error {
	s, err := storage.New(ctx, cfg.Storage)
	if err != nil {
		return fmt.Errorf("connecting to the storage: %w", err)
	}
	defer s.Close()

	repo, err := mongo.New(ctx, cfg.Mongo)
	if err != nil {
		return fmt.Errorf("connecting to the mongo: %w", err)
	}
	defer repo.Close()

	res := resolver.NewGraphqlResolver(s, processor.New(), repo, cfg.MaxImageSize)
	srv := graph.CreateAndRun(cfg.GraphQLPort, res)
	defer closeWithTimeout(srv.Close, time.Second*5)

	healthCheck := health.CreateAndRun(cfg.HealthCHeckPort, []health.Check{
		repo.HealthCheck,
		srv.HealthCheck,
	})
	defer closeWithTimeout(healthCheck.Close, 5*time.Second)

	<-ctx.Done()

	return nil
}

func closeWithTimeout(close func(context.Context), d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	close(ctx)
}

func setupGracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		logrus.Println("Got Interrupt signal")
		stop()
	}()
}

func initLogger(logLevel string, pretty bool) {
	if pretty {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	logrus.SetOutput(os.Stderr)

	switch strings.ToLower(logLevel) {
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}
