package btservice

import "tinygo.org/x/bluetooth"

//var advState = false

func BtInit(distance uint8, btconnected *bool) (*bluetooth.Adapter, bluetooth.Characteristic) {

	var adapter = bluetooth.DefaultAdapter
	// Bluetooth Init
	println("starting")
	must("enable BLE stack", adapter.Enable())
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName:    "Go Sense",
		ServiceUUIDs: []bluetooth.UUID{bluetooth.ServiceUUIDAdafruitProximity},
	}))

	adapter.SetConnectHandler(func(device bluetooth.Address, connected bool) {
		if connected {
			println("connected, not advertising...")
			*btconnected = true
		} else {
			println("disconnected, advertising...")
			*btconnected = false
		}
	})

	must("start adv", adv.Start())

	var distanceMeasurement bluetooth.Characteristic
	must("add service", adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.ServiceUUIDAdafruitProximity,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &distanceMeasurement,
				UUID:   bluetooth.CharacteristicUUIDAdafruitProximity,
				Value:  []byte{0, distance},
				Flags:  bluetooth.CharacteristicNotifyPermission,
			},
		},
	}))
	return adapter, distanceMeasurement
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
