package main

import (
	"machine"
	"rp2350-apps/lib/utils"
	"time"
)

func main() {
	utils.WaitForSerial("Blinky is ready!")

	led := machine.GPIO25
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
