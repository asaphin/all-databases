package datagenerator

import (
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/jaswdr/faker"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
)

type DataGenerator struct {
	faker.Faker
}

var dataGeneratorInstance *DataGenerator
var dataGeneratorSync = sync.Once{}

func New() *DataGenerator {
	dataGeneratorSync.Do(func() {
		dataGeneratorInstance = &DataGenerator{Faker: faker.New()}
	})

	return dataGeneratorInstance
}

type VehicleRentalDataGenerator struct {
	dg *DataGenerator
}

func (dg *DataGenerator) VR() *VehicleRentalDataGenerator {
	return &VehicleRentalDataGenerator{dg: New()}
}

func (vr *VehicleRentalDataGenerator) UnitedStatesAddress() domain.Address {
	return domain.Address{
		ID:           "",
		Type:         randomElement(addressTypes),
		InCareOfName: vr.dg.Person().Name(),
		Street:       vr.dg.Address().StreetName(),
		StreetNumber: vr.dg.Address().BuildingNumber(),
		Apartment:    strconv.Itoa(vr.dg.IntBetween(1, 100)),
		Suite:        "",
		Floor:        optional(strconv.Itoa(vr.dg.IntBetween(1, 25)), 0.5),
		City:         vr.dg.Address().City(),
		State:        vr.dg.Address().State(),
		Province:     "",
		Zip:          vr.dg.Address().PostCode(),
		PostalCode:   "",
		Country:      vr.dg.Address().Country(),
		Latitude:     vr.dg.Float64(16, -90, 90),
		Longitude:    vr.dg.Float64(16, -180, 180),
	}
}

func randomElement[T any](s []T) T {
	return s[rand.Intn(len(s))]
}

func optional[T any](v T, appearProbability float64) T {
	if rand.Float64() < appearProbability {
		return v
	}
	return zeroValue(v)
}

func zeroValue[T any](v T) T {
	return reflect.Zero(reflect.TypeOf(v)).Interface().(T)
}
