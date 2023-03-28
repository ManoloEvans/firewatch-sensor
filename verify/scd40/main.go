package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/scd4x"
)

var (
	i2c    = machine.I2C0
	sensor = scd4x.New(i2c)
)

func main() {
	time.Sleep(1500 * time.Millisecond)

	machine.PA9.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.PA9.High()

	time.Sleep(1500 * time.Millisecond)

	i2c.Configure(machine.I2CConfig{})

	time.Sleep(1500 * time.Millisecond)
	if err := sensor.Configure(); err != nil {
		println(err)
	}

	time.Sleep(1500 * time.Millisecond)

	if err := sensor.StartPeriodicMeasurement(); err != nil {
		println(err)
	}

	time.Sleep(1500 * time.Millisecond)

	for {
		co2, err := sensor.ReadCO2()
		humidity, err := sensor.ReadHumidity()
		temperature, err := sensor.ReadTemperature()
		temperature = temperature / 1000

		if err != nil {
			println(err)
		}
		println("Temperature", temperature, "°C")
		println("Humidity", humidity, "%")
		println("CO2", co2, "ppm")
		time.Sleep(time.Second)
	}
}
