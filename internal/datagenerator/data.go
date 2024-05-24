package datagenerator

import (
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/asaphin/all-databases-go/internal/utils"
)

var (
	vehicleManufacturersFile = "./testdata/datagenerator/vehicles/manufacturers.json"
	vehicleModelNamesFile    = "./testdata/datagenerator/vehicles/model-names.json"
)

var addressTypes = []domain.AddressType{domain.CustomerAddress, domain.StationAddress}

var manufacturers = func() map[domain.VehicleType][]string {
	var mnf map[domain.VehicleType][]string

	utils.MustUnmarshalJSONFromFile(vehicleManufacturersFile, &mnf)

	return mnf
}()

var vehicleTypeCumulativeDistribution = []struct {
	Type        domain.VehicleType
	Probability float64
}{
	{Type: domain.Car, Probability: 0.5},
	{Type: domain.Truck, Probability: 0.6},
	{Type: domain.Bike, Probability: 0.7},
	{Type: domain.Motorcycle, Probability: 0.8},
	{Type: domain.Boat, Probability: 0.9},
	{Type: domain.Plane, Probability: 1.0},
}

var vehicleModelNames = utils.MustMustUnmarshalJSONFromFileAsType(vehicleModelNamesFile, []string{})

var serialNumberRegexes = []string{
	`^[A-Z0-9]{10}$`,
	`^[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}$`,
	`^[A-Z]{2}[0-9]{6}[A-Z]{2}$`,
	`^[A-Z]{3}[0-9]{5}$`,
	`^[A-F0-9]{12}$`,
	`^[A-Z]{1,3}[0-9]{3,5}[A-Z]{2}$`,
	`^[A-Z0-9]{8,12}$`,
}
