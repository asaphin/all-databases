package internal

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/app"
	"github.com/asaphin/all-databases-go/internal/infrastructure/ledger"
	"github.com/asaphin/all-databases-go/internal/infrastructure/postgres"
	log "github.com/sirupsen/logrus"
)

type AllDatabasesService struct {
}

func NewAllDatabasesService() *AllDatabasesService {
	return &AllDatabasesService{}
}

func (a *AllDatabasesService) Run(_ context.Context) error {
	brandsRepository, err := postgres.NewSQLXBrandsRepository()
	if err != nil {
		log.WithError(err).Error("unable to create brands repository")
		return err
	}

	entityLedger := ledger.NewInMemoryEntitiesLedger()

	vehicleRentalScenario := app.NewVehicleRentalScenarioService(brandsRepository, entityLedger)

	vehicleRentalScenario.Run()

	return nil
}

func (a *AllDatabasesService) Shutdown(_ context.Context) error {
	postgres.Shutdown()

	return nil
}
