package main

import (
	"machine"
	"rp2350-apps/lib/utils"
	"time"
)

func main() {
	utils.WaitForSerial("Interrupt is ready now!")
	utils.BlinkLEDWhileAlive(machine.GPIO25, time.Millisecond*333)
	utils.BOOTSELOnButtonPress(machine.GPIO1)

	for {
		time.Sleep(time.Minute)
	}
}
