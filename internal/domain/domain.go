package domain

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type AddressType string

const (
	CustomerAddress AddressType = "customer"
	StationAddress  AddressType = "station"
)

type Address struct {
	ID             string
	Type           AddressType
	InCareOfName   string
	Street         string
	StreetNumber   string
	Apartment      string
	Locality       string
	Region         string
	PostalCode     string
	Country        string
	AdditionalInfo map[string]string
	Latitude       float64
	Longitude      float64
}

func (a Address) String() string {
	var builder strings.Builder

	var err error

	if a.InCareOfName != "" {
		_, err = fmt.Fprintf(&builder, "In care of name %s, ", a.InCareOfName)
		if err != nil {
			log.WithError(err).Error("unable to add name info to address string")
		}
	}

	_, err = fmt.Fprintf(&builder, "%s, %s", a.StreetNumber, a.Street)
	if err != nil {
		log.WithError(err).Error("unable to add street info to address string")
	}
	if a.Apartment != "" {
		_, err = fmt.Fprintf(&builder, ", Apt %s", a.Apartment)
		if err != nil {
			log.WithError(err).Error("unable to add apartment info to address string")
		}
	}

	if len(a.AdditionalInfo) > 0 {
		builder.WriteString(" (")
		fields := make([]string, 0, len(a.AdditionalInfo))
		for key, value := range a.AdditionalInfo {
			fields = append(fields, fmt.Sprintf("%s: %s", key, value))
		}
		builder.WriteString(strings.Join(fields, ", "))
		builder.WriteString(")")
	}

	_, err = fmt.Fprintf(&builder, ", %s, %s, %s, %s", a.Locality, a.Region, a.Country, a.PostalCode)
	if err != nil {
		log.WithError(err).Error("unable to add locality info to address string")
	}

	_, err = fmt.Fprintf(&builder, " Lat: %.6f, Long: %.6f", a.Latitude, a.Longitude)
	if err != nil {
		log.WithError(err).Error("unable to add coordinates info to address string")
	}

	builder.WriteString(fmt.Sprintf(" (%s)", a.Type))

	return builder.String()
}

type Brand struct {
	ID     string
	Name   string
	Slogan string
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
