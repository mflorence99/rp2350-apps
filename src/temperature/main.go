package main

import (
	"fmt"
	"image/color"
	"machine"
	"rp2350-apps/lib/utils"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/drivers/thermistor"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var lg = &freesans.Regular12pt7b
var sm = &freesans.Regular9pt7b

func main() {
	defer func() { utils.RecoverFromPanic(recover()) }()

	utils.WaitForSerial("Temperature is ready!")
	utils.BlinkLEDWhileAlive(machine.GPIO25, time.Millisecond*333)
	utils.BOOTSELOnButtonPress(machine.GPIO1)
	display := utils.ConfigureSSD1306Display(machine.I2C1, machine.GPIO6, machine.GPIO7)

	w, h := display.Size()

	machine.InitADC()
	sensor := thermistor.New(machine.ADC0)
	sensor.HighSide = true
	sensor.Configure()

	ra := utils.NewRollingAverage(20)

	for {

		reading, err := sensor.ReadTemperature()
		if err != nil {
			println("🔥", err.Error())
			break
		}

		temp := ((float64(reading) / 1000) * float64(1.8)) + 32
		smoothed := ra.Add(temp)

		display.ClearBuffer()

		center(display, sm, 0, 0, w, 0, "Temperature", utils.Yellow)
		center(display, lg, 0, 0, w, h, fmt.Sprintf("%3.1f F", smoothed), utils.Blue)

		display.Display()

		time.Sleep(time.Millisecond * 1000)

	}
}

func center(display *ssd1306.Device, f tinyfont.Fonter, x, y, w, h int16, str string, c color.RGBA) {
	_, lw := tinyfont.LineWidth(f, str)
	lh := int16(tinyfont.GetGlyph(f, rune(str[0])).Info().Height)
	lx := max(w-int16(lw), 0) / 2
	ly := max(h-lh, 0) / 2
	// 👇 y coord for text is baseline
	tinyfont.WriteLine(display, f, lx, ly+lh, str, c)
}
