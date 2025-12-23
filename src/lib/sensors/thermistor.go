package sensors

import (
	"machine"
	"rp2350-apps/lib/utils"

	"tinygo.org/x/drivers/thermistor"
)

type Thermistor struct {
	sensor thermistor.Device
	ra     *utils.RollingAverage
}

func NewThermistor(adc machine.Pin) *Thermistor {
	t := new(Thermistor)
	machine.InitADC()
	t.sensor = thermistor.New(adc)
	t.sensor.Configure()
	t.ra = utils.NewRollingAverage(20)
	return t
}

func (t *Thermistor) MustReadTemperature() float64 {
	reading, err := t.sensor.ReadTemperature()
	if err != nil {
		panic(err.Error())
	}
	temp := ((float64(reading) / 1000) * float64(1.8)) + 32
	smoothed := t.ra.Add(temp)
	return smoothed
}
