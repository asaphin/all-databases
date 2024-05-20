package datagenerator

import (
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/jaswdr/faker"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
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

func (vr *VehicleRentalDataGenerator) Address() domain.Address {
	addressType := randomElement(addressTypes)

	var inCareOfName string

	if addressType == domain.CustomerAddress {
		inCareOfName = optional(vr.dg.Person().Name(), 0.3)
	}

	country := randomChioce("United States", vr.dg.Address().Country(), 0.4)
	var region string
	if country == "United States" {
		region = vr.dg.Address().State()
	} else {
		region = cases.Title(language.English, cases.Compact).String(vr.dg.Lorem().Word())
	}

	additionalInfo := make(map[string]string)

	optionalCall(func() {
		additionalInfo["floor"] = strconv.Itoa(vr.dg.IntBetween(0, 10))
	}, 0.3)

	optionalCall(func() {
		additionalInfo["block"] = strconv.Itoa(vr.dg.IntBetween(0, 5))
	}, 0.3)

	if len(additionalInfo) == 0 {
		additionalInfo = nil
	}

	return domain.Address{
		ID:             "",
		Type:           addressType,
		InCareOfName:   inCareOfName,
		Street:         vr.dg.Address().StreetName(),
		StreetNumber:   strings.TrimPrefix(vr.dg.Address().BuildingNumber(), "%"),
		Apartment:      optional(strconv.Itoa(vr.dg.IntBetween(1, 100)), 0.5),
		Locality:       vr.dg.Address().City(),
		Region:         region,
		PostalCode:     vr.dg.Address().PostCode(),
		Country:        country,
		AdditionalInfo: additionalInfo,
		Latitude:       vr.dg.Float64(16, -90, 90),
		Longitude:      vr.dg.Float64(16, -180, 180),
	}
}

func (vr *VehicleRentalDataGenerator) Vehicle() domain.Vehicle {
	probabilityValue := rand.Float64()

	var vehicleType domain.VehicleType

	for _, typeProbability := range vehicleTypeCumulativeDistribution {
		vehicleType = typeProbability.Type

		if typeProbability.Probability > probabilityValue {
			break
		}
	}

	return domain.Vehicle{
		Manufacturer: randomElement(manufacturers[vehicleType]),
		Type:         vehicleType,
	}
}

func randomElement[T any](s []T) T {
	return s[rand.Intn(len(s))]
}

func randomChioce[T any](e1, e2 T, firstElementProbability float64) T {
	if rand.Float64() < firstElementProbability {
		return e1
	}

	return e2
}

func optional[T any](v T, appearProbability float64) T {
	if rand.Float64() < appearProbability {
		return v
	}
	return zeroValue(v)
}

func optionalCall(f func(), callProbability float64) {
	if rand.Float64() < callProbability {
		f()
	}
}

func zeroValue[T any](v T) T {
	return reflect.Zero(reflect.TypeOf(v)).Interface().(T)
}
