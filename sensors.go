package main

import (
	"machine"

	"tinygo.org/x/drivers/scd4x"
)

var (
	// devices
	co2Sensor *scd4x.Device

	// readings
	co2SensorReading int32
	temperature      int32
	humidity         int32
)

// startSensors initializes the sensors.
func startSensors() {
	machine.I2C0.Configure(machine.I2CConfig{})

	co2Sensor = scd4x.New(machine.I2C0)
	co2Sensor.Configure()

	if err := co2Sensor.StartPeriodicMeasurement(); err != nil {
		println(err)
	}
}

// readSensors reads the sensors.
func readSensors() {
	temperature, _ = co2Sensor.ReadTemperature()
	co2SensorReading, _ = co2Sensor.ReadCO2()
	humidity, _ = co2Sensor.ReadHumidity()

	println("Temperature", temperature, "°C")
	println("Humidity", humidity, "%")
	println("CO2", co2SensorReading, "ppm")

}
