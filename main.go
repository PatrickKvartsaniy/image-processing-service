package main

import (
	"context"
	"github.com/PatrickKvartsaniy/image-processing-service/config"
	"github.com/PatrickKvartsaniy/image-processing-service/graph"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/resolver"
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
	initlogger(cfg.LogLevel, cfg.PrettyLogOutput)

	ctx, cancel := context.WithCancel(context.Background())
	setupGracefulShutdown(cancel)

	s, err := storage.New(ctx, cfg.Storage)
	if err != nil {
		logrus.WithError(err).Error("connecting to storage")
	}
	defer s.Close()

	repo, err := mongo.New(ctx, cfg.Mongo)
	if err != nil {
		logrus.WithError(err).Error("connecting to mongo")
	}
	defer repo.Close()

	p := processor.New()
	res := resolver.NewGraphqlResolver(s, p, repo, cfg.MaxImageSize)
	srv := graph.CreateAndRun(cfg.GraphQLPort, res)
	defer closeWithTimeout(srv.Close, time.Second*5)
	<-ctx.Done()
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

func initlogger(logLevel string, pretty bool) {
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
