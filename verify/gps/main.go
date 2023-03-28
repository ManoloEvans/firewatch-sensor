package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/gps"
)

func main() {
	println("GPS UART Example")
	// Turn on power to pins.
	machine.PA9.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.PA9.High()
	machine.UART2.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.UART2_TX_PIN, RX: machine.UART2_RX_PIN})
	ublox := gps.NewUART(machine.UART2)
	parser := gps.NewParser()
	var fix gps.Fix
	for {
		s, err := ublox.NextSentence()
		if err != nil {
			println(err)
			continue
		}

		fix, err = parser.Parse(s)
		if err != nil {
			println(err)
			continue
		}
		if fix.Valid {
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
		} else {
			println("No fix")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
