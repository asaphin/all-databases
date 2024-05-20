package app

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BrandsScenarioService struct {
	brandsRepository BrandsRepository
	ledger           map[string]struct{}
}

func NewBrandsScenarioService(brandsRepository BrandsRepository) *BrandsScenarioService {
	return &BrandsScenarioService{
		brandsRepository: brandsRepository,
		ledger:           make(map[string]struct{}),
	}
}

func (s *BrandsScenarioService) Run() {
	brands, err := s.brandsRepository.List(context.Background(), 5, 5)
	if err != nil {
		log.WithError(err).Error("unable to get list of brands")
		return
	}

	log.WithField("brands", brands).Info("got list of brands")

	newBrand := &domain.Brand{
		Name:   "Vintage Voyage",
		Slogan: "Experience the charm of a bygone era",
	}

	id, err := s.brandsRepository.Create(context.Background(), newBrand)

	if err != nil {
		log.WithError(err).Error("unable to create brand")
		return
	}

	log.WithFields(log.Fields{"id": id, "brand": newBrand}).Info("new brand created")

	brand, err := s.brandsRepository.GetByID(context.Background(), id)
	if err != nil {
		log.WithError(err).WithField("id", id).Error("unable to get brand")
		return
	}

	log.WithFields(log.Fields{"id": id, "brand": brand}).Info("got a brand")

	err = s.brandsRepository.Delete(context.Background(), id)
	if err != nil {
		log.WithError(err).WithField("id", id).Error("unable to delete brand")
		return
	}

	log.WithField("id", id).Info("brand deleted")
}
