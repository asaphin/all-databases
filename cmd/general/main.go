package main

import (
	"github.com/asaphin/all-databases-go/internal"
	"github.com/asaphin/runner"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetReportCaller(true)

	log.Trace("initialization done")
}

func main() {
	log.Trace("main() called")

	svc := internal.NewAllDatabasesService()

	runner.Run(svc)
}
