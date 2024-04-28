package proxyservice

import (
	"machine"

	"tinygo.org/x/bluetooth"
	"tinygo.org/x/drivers/apds9960"
)

func ProximityInit(distance uint8, distanceMeasurement bluetooth.Characteristic) apds9960.Device {
	// Proximit Init
	// Init Proxy Sensor
	// If I ever need it 0×39
	machine.I2C0.Configure(machine.I2CConfig{Frequency: 400000})

	sensor := apds9960.New(machine.I2C0)

	// use default settings
	sensor.Configure(apds9960.Configuration{})

	if !sensor.Connected() {
		println("APDS-9960 not connected!")
	}

	sensor.EnableProximity() // enable proximity engine

	//Finished Proxy Sensor Init
	return sensor
}
