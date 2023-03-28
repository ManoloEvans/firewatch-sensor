package main

import (
	"machine"

	"time"

	"tinygo.org/x/drivers/gps"
)

var fix gps.Fix

// startGPS initializes the GPS.
func startGPS() {
	machine.UART2.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.UART2_TX_PIN, RX: machine.UART2_RX_PIN})
	ublox := gps.NewUART(machine.UART2)
	parser := gps.NewParser()
	for {
		s, err := ublox.NextSentence()
		if err != nil {
			continue
		}

		newfix, err := parser.Parse(s)
		if err != nil {
			continue
		}
		if newfix.Valid {
			fix = newfix
			print(fix.Time.Format("15:04:05"))
			print(", lat=")
			print(fix.Latitude)
			print(", long=")
			print(fix.Longitude)
			print(", altitude=", fix.Altitude)
			print(", satellites=", fix.Satellites)
			if fix.Speed != 0 {
				print(", speed=")
				print(fix.Speed)
			}
			if fix.Heading != 0 {
				print(", heading=")
				print(fix.Heading)
			}
			println()
		}
		time.Sleep(200 * time.Millisecond)
	}
}
