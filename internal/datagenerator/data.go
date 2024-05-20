package datagenerator

import (
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/asaphin/all-databases-go/internal/utils"
)

var manufacturersFile = "./testdata/datagenerator/vehicles/manufacturers.json"

var addressTypes = []domain.AddressType{domain.CustomerAddress, domain.StationAddress}

var manufacturers = func() map[domain.VehicleType][]string {
	var mnf map[domain.VehicleType][]string

	utils.MustUnmarshalJSONFromFile(manufacturersFile, &mnf)

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
