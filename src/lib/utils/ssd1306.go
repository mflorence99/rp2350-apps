package utils

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

func ConfigureSSD1306Display(I2C *machine.I2C, SDA, SCL machine.Pin) *ssd1306.Device {
	I2C.Configure(machine.I2CConfig{SDA: SDA, SCL: SCL, Frequency: 400 * machine.KHz})
	// 🔥 who knows why?
	time.Sleep(time.Second * 1)
	display := ssd1306.NewI2C(machine.I2C1)
	display.Configure(ssd1306.Config{Width: 128, Height: 64, Address: ssd1306.Address_128_32, VccState: ssd1306.SWITCHCAPVCC})
	display.Sleep(false)
	display.ClearBuffer()
	display.ClearDisplay()
	return display
}
