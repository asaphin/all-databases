package internal

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/app"
	"github.com/asaphin/all-databases-go/internal/infrastructure/postgres"
	log "github.com/sirupsen/logrus"
)

type AllDatabasesService struct {
}

func NewAllDatabasesService() *AllDatabasesService {
	return &AllDatabasesService{}
}

func (a *AllDatabasesService) Run(_ context.Context) error {
	addressesRepository, err := postgres.NewSQLXAddressRepository()
	if err != nil {
		log.WithError(err).Error("unable to create addresses repository")
		return err
	}

	addressesScenario := app.NewAddressesScenarioService(addressesRepository)

	addressesScenario.Run()

	return nil
}

func (a *AllDatabasesService) Shutdown(_ context.Context) error {
	postgres.Shutdown()

	return nil
}
