package domain

import (
	"time"
)

type AddressType string

const (
	CustomerAddress AddressType = "customer"
	StationAddress  AddressType = "station"
)

type Address struct {
	ID           string
	Type         string
	InCareOfName string
	Street       string
	StreetNumber string
	Apartment    string
	Suite        string
	Floor        string
	City         string
	State        string
	Province     string
	Zip          string
	PostalCode   string
	Country      string
	Latitude     float64
	Longitude    float64
}

type Brand struct {
	ID   string
	Name string
}

type Station struct {
	ID           string
	Brand        Brand
	Name         string
	Description  string
	Address      Address
	Phone        string
	Email        string
	VehicleTypes []VehicleType
	Vehicles     []*Vehicle
}

type VehicleType string

const (
	Car   VehicleType = "car"
	Truck VehicleType = "truck"
	Bike  VehicleType = "bike"
)

type VehicleStatus string

const (
	Available VehicleStatus = "available"
	Rented    VehicleStatus = "rented"
)

type Vehicle struct {
	ID           string
	Manufacturer string
	Model        string
	SerialNumber string
	Year         int
	Type         VehicleType
	Status       VehicleStatus
	Metadata     map[string]string
}

type Customer struct {
	ID            string
	FirstName     string
	LastName      string
	BirthDate     time.Time
	LicenseNumber string
	PhoneNumber   string
	Email         string
	Address       Address
}

type RentalStatus string

const (
	New    RentalStatus = "new"
	Active RentalStatus = "active"
	Closed RentalStatus = "closed"
)

type Rental struct {
	ID             string
	Vehicle        Vehicle
	Customer       Customer
	PickupStation  Station
	DropOffStation Station
	StartDate      time.Time
	EndDate        time.Time
	Status         RentalStatus
}

type SensorType string

const (
	GPS     SensorType = "gps"
	Fuel    SensorType = "fuel"
	Mileage SensorType = "mileage"
)

type Sensor struct {
	ID           string
	Vehicle      Vehicle
	Manufacturer string
	Model        string
	SerialNumber string
	Type         SensorType
	Data         []SensorData
}

type SensorData struct {
	Timestamp       time.Time
	SensorID        string
	ParameterName   string
	Value           interface{}
	MeasurementUnit string
}
