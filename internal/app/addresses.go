package app

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/datagenerator"
	"github.com/asaphin/all-databases-go/internal/domain"
	log "github.com/sirupsen/logrus"
)

type AddressesScenarioService struct {
	addressesRepository AddressesRepository
	ledger              map[string]struct{}
}

func NewAddressesScenarioService(addressesRepository AddressesRepository) *AddressesScenarioService {
	return &AddressesScenarioService{
		addressesRepository: addressesRepository,
		ledger:              make(map[string]struct{}),
	}
}

func (s *AddressesScenarioService) Run() {
	err := s.createSampleAddresses()
	if err != nil {
		log.WithError(err).Error("unable to create sample addresses")
	}

	s.cleanupAddresses()
}

func (s *AddressesScenarioService) createSampleAddresses() error {
	n := 10

	addresses := make([]*domain.Address, 0, n)

	for i := 0; i < n; i++ {
		addr := datagenerator.New().VR().Address()

		addresses = append(addresses, &addr)
	}

	for i := range addresses {
		id, err := s.addressesRepository.Create(context.Background(), addresses[i])
		if err != nil {
			return err
		}

		s.ledger[id] = struct{}{}
	}

	return nil
}

func (s *AddressesScenarioService) cleanupAddresses() {
	ctx := context.Background()

	for id := range s.ledger {
		err := s.addressesRepository.Delete(ctx, id)
		if err != nil {
			log.WithError(err).WithField("addressID", id).Warning("unable to delete address")
		}
	}
}
