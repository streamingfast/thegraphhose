package main

import (
	"context"
	"net/http"
	"time"

	"github.com/streamingfast/dmetrics"
	"go.uber.org/zap"
)

var cmdCtx = func() (ctx context.Context, cancel func()) {
	return context.WithTimeout(context.Background(), 1*time.Minute)
}

func setup() {
	setupMetrics()
	setupProfiler()
}

func setupMetrics() {
	go dmetrics.Serve(metricsListenAddr)
}

func setupProfiler() {
	go func() {
		err := http.ListenAndServe(pprofListenAddr, nil)
		if err != nil {
			zlog.Debug("unable to start profiling server", zap.Error(err), zap.String("listen_addr", pprofListenAddr))
		}
	}()
}
