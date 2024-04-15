package domain

import (
	"time"
)

type User struct {
	UserID    string
	FirstName string
	LastName  string
	Username  string
	Email     string
	Phone     string
}

type Device struct {
	DeviceID          string
	Name              string
	OwnerID           string
	Manufacturer      string
	Model             string
	FirmwareVersion   string
	Configuration     map[string]interface{}
	Metadata          map[string]string
	LastCommunication string
}

type Status struct {
	ID         string
	StatusName string
	DeviceID   string
	Timestamp  time.Time
}

type Sensor struct {
	SensorID          string
	Type              string
	Manufacturer      string
	Model             string
	Accuracy          float64
	UnitOfMeasurement string
}

type SensorData struct {
	SensorID  string
	DeviceID  string
	Timestamp time.Time
	Value     float64
}

type Alert struct {
	AlertID   string
	DeviceID  string
	Timestamp time.Time
	Message   string
}

type Event struct {
	ID           string
	DeviceID     string
	Timestamp    time.Time
	EventType    string
	EventDetails map[string]interface{}
}
