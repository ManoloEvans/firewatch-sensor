package main

import (
	"time"

	"tinygo.org/x/drivers/lora"
	"tinygo.org/x/drivers/lora/lorawan"
	"tinygo.org/x/drivers/lora/lorawan/region"

	"machine"
)

func main() {
	time.Sleep(5 * time.Second)
	println("*** Firewatch 1 starting... ***")

	machine.POWER_EN3V3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.POWER_EN3V3.High()

	time.Sleep(1500 * time.Millisecond)

	// setup LoRa radio
	var err error
	radio, err = setupLora()
	if err != nil {
		failMessage(err)
	}

	// Connect LoRaWAN to use the LoRa Radio device.
	lorawan.UseRadio(radio)

	// use EU868 DR2 spreading factor for high-altitude
	settings := region.EU868()
	settings.UplinkChannel().SpreadingFactor = lora.SpreadingFactor10
	lorawan.UseRegionSettings(settings)

	// Try to connect to the LoRaWAN network
	if err := lorawanJoin(); err != nil {
		failMessage(err)
	}

	go startGPS()

	startSensors()

	for {
		println("Sleeping for", uplinkDelaySeconds, "seconds")
		time.Sleep(time.Second * uplinkDelaySeconds)

		readSensors()

		payload, err := createPayload()
		if err != nil {
			println("Payload error:", err)
			continue
		}

		if err := lorawan.SendUplink(payload, session); err != nil {
			println("Uplink error:", err)
			continue
		}

		println("Uplink complete, msglen=", len(payload))
	}
}

func failMessage(err error) {
	for {
		println("FATAL:", err)
		time.Sleep(time.Second)
	}
}
