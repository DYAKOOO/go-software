package main 

import (
	"os"
	"context"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ardanlabs/service/foundation/logger"
	// "github.com/ardanlabs/service/foundation/web"

)
var build = "develop"

// main is the entry point for the running instance.
func main()  {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT ******")
		},
	}

	traceIDFunc := func(ctx context.Context) string {
		return ""
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES-API", traceIDFunc, events)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
		return 
	}

	return 
}

func run(ctx context.Context, log *logger.Logger) error {
	// GOMAXPROCS
	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0), "build" , build)


	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	sig:= <-shutdown

	log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
	defer log.Info(ctx, "shutdown", "status", "shutdown complete", "signal", sig)
	
	return nil
}
