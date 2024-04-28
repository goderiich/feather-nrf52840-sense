package main

import (
	"anderssavill/btservice"
	"anderssavill/displayservice"
	"anderssavill/proxyservice"
	"time"

	"tinygo.org/x/bluetooth"
)

// TODO: use atomics to access this value.
var distance uint8 = 255

// Use custom UUID
var customUUID, _ = bluetooth.ParseUUID("7e3b7c14-3c1d-4102-b00e-96462c02df36")

func main() {

	var connected bool = false
	var previous bool = false

	_, distanceMeasurement := btservice.BtInit(distance, &connected)

	sensor := proxyservice.ProximityInit(distance, distanceMeasurement)

	display, ddev := displayservice.DisplayInit()

	for {
		if sensor.ProximityAvailable() {
			// Make the proximity value compatible with the bluetooth standard
			p := uint8(sensor.ReadProximity())

			if connected != previous {
				ddev.ClearBuffer()
			}

			if connected == true {
				display.PrintText("Connected")
			} else {
				display.PrintText("Advertising")
			}

			if distance != p {
				distance = p
				println("Proximity:", p)
				// and push the next notification
				distanceMeasurement.Write([]byte{0, p})
				// I set this timeout so that you can actually read the value on the screen
				time.Sleep(time.Millisecond * 300)
			}
		}

	}

}
