package main

import (
	"fmt"
	"machine"
	"rp2350-apps/lib/displays"
	"rp2350-apps/lib/sensors"
	"rp2350-apps/lib/utils"
	"time"

	"tinygo.org/x/tinyfont/freesans"
)

var lg = &freesans.Regular12pt7b
var sm = &freesans.Regular9pt7b

func main() {
	defer func() { utils.RecoverFromPanic(recover()) }()

	utils.WaitForSerial("Temperature is ready!")
	utils.BlinkLEDWhileAlive(machine.GPIO25, time.Millisecond*333)
	utils.BOOTSELOnButtonPress(machine.GPIO1)

	ssd1306 := displays.NewSSD1306(machine.I2C1, machine.GPIO6, machine.GPIO7)
	w, h := ssd1306.Size()

	thermistor := sensors.NewThermistor(machine.ADC0)

	for {

		ssd1306.ClearBuffer()

		temp := thermistor.MustReadTemperature()

		println("🔥 temp:", temp)

		utils.CenterText(ssd1306, sm, 0, 0, w, 0, "Temperature", utils.Yellow)
		utils.CenterText(ssd1306, lg, 0, 0, w, h, fmt.Sprintf("%3.1f F", temp), utils.Blue)

		ssd1306.Display()

		time.Sleep(time.Millisecond * 1000)

	}
}
