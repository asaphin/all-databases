package main

import (
	"github.com/asaphin/all-databases-go/internal"
	"github.com/asaphin/runner"
	log "github.com/sirupsen/logrus"
	"time"
)

func init() {
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
	log.SetReportCaller(true)

	log.Trace("initialization done")
}

func main() {
	log.Trace("main() called")

	svc := internal.NewAllDatabasesService()

	runner.Run(svc)
}
