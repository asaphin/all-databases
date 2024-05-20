package app

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/asaphin/all-databases-go/internal/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const brandsResource = "brands"

const brandsFile = "./testdata/vehicle_rental/brands.csv"

type VehicleRentalScenarioService struct {
	brandsRepository BrandsRepository
	ledger           EntitiesLedger
}

func NewVehicleRentalScenarioService(brandsRepository BrandsRepository, ledger EntitiesLedger) *VehicleRentalScenarioService {
	return &VehicleRentalScenarioService{
		brandsRepository: brandsRepository,
		ledger:           ledger,
	}
}

func (s *VehicleRentalScenarioService) Run() {
	actions := []Action{
		{name: "create brands", function: s.createBrands},
		{name: "get brands", function: s.getBrands},
		{name: "get random brand", function: s.getRandomBrand},
	}

	for _, action := range actions {
		err := action.function()
		if err != nil {
			log.WithError(err).Errorf("%s action failed", action.name)
			break
		}
	}

	cleanupActions := []Action{
		{name: "cleanup brands", function: s.cleanupBrands},
	}

	for _, action := range cleanupActions {
		err := action.function()
		if err != nil {
			log.WithError(err).Errorf("%s action failed", action.name)
		}
	}
}

func (s *VehicleRentalScenarioService) createBrands() error {
	newBrands := utils.MustUnmarshalCSVFromFile(brandsFile, ';', domain.Brand{})

	for _, newBrand := range newBrands {
		id, err := s.brandsRepository.Create(context.Background(), &newBrand)
		if err != nil {
			return err
		}

		utils.LogAsWarningIfError(s.ledger.Add(NewEntity(brandsResource, id)))
	}

	log.Infof("created %d brands", len(newBrands))

	return nil
}

func (s *VehicleRentalScenarioService) getBrands() error {
	brands, err := s.brandsRepository.List(context.Background(), 0, 0)
	if err != nil {
		return err
	}

	log.Infof("retrieved %d brands", len(brands))

	return nil
}

func (s *VehicleRentalScenarioService) getRandomBrand() error {
	entities, err := s.ledger.GetByResource(brandsResource)
	if err != nil {
		return errors.Wrapf(err, "unable to get created entities for resource %s", brandsResource)
	}

	id := utils.GetRandomElement(entities).Key.String(0)

	brand, err := s.brandsRepository.GetByID(context.Background(), id)
	if err != nil {
		return errors.Wrapf(err, "unable to get created brand %s", id)
	}

	log.WithField("brand", brand).Info("got random brand")

	return nil
}

func (s *VehicleRentalScenarioService) cleanupBrands() error {
	entities, err := s.ledger.GetByResource(brandsResource)
	if err != nil {
		return errors.Wrapf(err, "unable to get created entities for resource %s", brandsResource)
	}

	for i := range entities {
		id := entities[i].Key.String(0)

		err = s.brandsRepository.Delete(context.Background(), id)
		if err != nil {
			return errors.Wrapf(err, "unable to delete brand %s", id)
		}
	}

	log.Info("brands cleared")

	return nil
}
