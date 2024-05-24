package datagenerator

import (
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/brianvoe/gofakeit/v7/source"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
)

type DataGenerator struct {
	faker *gofakeit.Faker
}

var dataGeneratorInstance *DataGenerator
var dataGeneratorSync = sync.Once{}

func New() *DataGenerator {
	dataGeneratorSync.Do(func() {
		dataGeneratorInstance = &DataGenerator{faker: gofakeit.NewFaker(source.NewCrypto(), true)}
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
		inCareOfName = optional(vr.dg.faker.Person().FirstName+vr.dg.faker.Person().LastName, 0.3)
	}

	address := vr.dg.faker.Address()

	country := randomChoice("United States", address.Country, 0.4)
	var region string
	if country == "United States" {
		region = address.State
	} else {
		region = cases.Title(language.English, cases.Compact).String(vr.dg.faker.LoremIpsumWord())
	}

	additionalInfo := make(map[string]string)

	optionalCall(func() {
		additionalInfo["floor"] = strconv.Itoa(vr.dg.faker.IntRange(0, 10))
	}, 0.3)

	optionalCall(func() {
		additionalInfo["block"] = strconv.Itoa(vr.dg.faker.IntRange(0, 5))
	}, 0.3)

	if len(additionalInfo) == 0 {
		additionalInfo = nil
	}

	return domain.Address{
		ID:             "",
		Type:           addressType,
		InCareOfName:   inCareOfName,
		Street:         address.Street,
		StreetNumber:   strconv.Itoa(vr.dg.faker.IntRange(1, 10000)),
		Apartment:      strconv.Itoa(vr.dg.faker.IntRange(1, 1000)),
		Locality:       address.City,
		Region:         region,
		PostalCode:     address.Zip,
		Country:        country,
		AdditionalInfo: additionalInfo,
		Latitude:       address.Latitude,
		Longitude:      address.Longitude,
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
		Manufacturer: vr.dg.faker.RandomString(manufacturers[vehicleType]),
		Type:         vehicleType,
		Model:        vr.dg.faker.RandomString(vehicleModelNames) + " " + vr.dg.faker.Regex(`^[A-Z1-9]{2,3}$`),
		SerialNumber: vr.dg.faker.Regex(vr.dg.faker.RandomString(serialNumberRegexes)),
		Year:         vr.dg.faker.IntRange(2004, 2024),
		Status:       domain.Available,
		Metadata:     vr.dg.faker.Map(),
	}
}

func randomElement[T any](s []T) T {
	return s[rand.Intn(len(s))]
}

func randomChoice[T any](e1, e2 T, firstElementProbability float64) T {
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
