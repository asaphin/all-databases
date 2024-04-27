package main

import (
	"github.com/asaphin/all-databases-go/internal/app"
	"github.com/asaphin/all-databases-go/internal/infrastructure/postgres"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
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

	shutdown := func() error {
		postgres.Shutdown()

		return nil
	}

	go handleShutdown(shutdown)

	defer func() {
		shutdownErr := shutdown()
		if shutdownErr != nil {
			log.WithError(shutdownErr).Fatal("unable to shutdown correrctly")
		}
	}()

	addressesRepository, err := postgres.NewSQLXAddressRepository()
	if err != nil {
		log.WithError(err).Fatal("unable to create addresses repository")
	}

	addressesScenario := app.NewAddressesScenarioService(addressesRepository)

	addressesScenario.Run()
}

func handleShutdown(shutdown func() error) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	sig := <-sigCh
	log.WithField("signal", sig.String()).Info("received system signal, shutting down")

	err := shutdown()
	if err != nil {
		log.WithField("signal", sig).WithError(err).Error("unable to shutdown gracefully")
	}

	os.Exit(0)
}
