package main

import (
	"machine"

	"errors"

	"tinygo.org/x/drivers/lora"
	"tinygo.org/x/drivers/sx126x"
)

var (
	spi = machine.SPI3

	loraRadio *sx126x.Device
)

// do sx126x setup here
func setupLora() (lora.Radio, error) {
	loraRadio = sx126x.New(spi)
	loraRadio.SetDeviceType(sx126x.DEVICE_TYPE_SX1262)

	loraRadio.SetRadioController(sx126x.NewRadioControl())

	if state := loraRadio.DetectDevice(); !state {
		return nil, errors.New("LoRa radio not found")
	}

	return loraRadio, nil
}
