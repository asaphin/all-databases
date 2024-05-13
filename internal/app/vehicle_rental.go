package app

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/asaphin/all-databases-go/internal/utils"
	log "github.com/sirupsen/logrus"
)

const brandsResource = "brands"

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
	s.createBrand()
	s.cleanupBrands()
}

func (s *VehicleRentalScenarioService) createBrand() {
	newBrand := &domain.Brand{
		BrandListItem: domain.BrandListItem{
			Name: "Vintage Voyage",
		},
		Slogan: "Experience the charm of a bygone era",
	}

	id, err := s.brandsRepository.Create(context.Background(), newBrand)
	if err != nil {
		log.WithError(err).Error("unable to create brand")
		return
	}

	log.WithFields(log.Fields{"id": id, "brand": newBrand}).Info("new brand created")

	utils.LogAsWarningIfError(s.ledger.Add(NewEntity(brandsResource, id)))

	brand, err := s.brandsRepository.GetByID(context.Background(), id)
	if err != nil {
		log.WithError(err).WithField("id", id).Error("unable to get brand")
		return
	}

	log.WithFields(log.Fields{"id": id, "brand": brand}).Info("got a brand")
}

func (s *VehicleRentalScenarioService) cleanupBrands() {
	entities, err := s.ledger.GetByResource(brandsResource)
	if err != nil {
		log.WithError(err).WithField("resource", brandsResource).Error("unable to get created entities")
		return
	}

	for i := range entities {
		id := entities[i].Key.String(0)

		err = s.brandsRepository.Delete(context.Background(), id)
		if err != nil {
			log.WithError(err).WithField("id", id).Error("unable to delete brand")
			return
		}
	}

	log.Info("brands cleared")
}
