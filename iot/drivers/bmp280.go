package drivers

import (
	"fmt"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

type BME280 struct {
	i2cDriver *i2c.BME280Driver
}

type BME280Metrics struct {
	Temperature string `json:"temperature"`
	Pressure    string `json:"pressure"`
	Altitude    string `json:"altitude"`
	Humidity    string `json:"humidity"`
}

func NewBME280(raspiAdaptor *raspi.Adaptor) *BME280 {
	return &BME280{
		i2cDriver: i2c.NewBME280Driver(raspiAdaptor),
	}
}

func (b *BME280) String() string {
	return b.i2cDriver.Name()
}

func (b *BME280) Start() error {
	return b.i2cDriver.Start()
}

func (b *BME280) Stop() error {
	return b.i2cDriver.Halt()
}

func (b *BME280) Temperature() string {
	temp, err := b.i2cDriver.Temperature()
	if err != nil {
		temp = 0
	}
	return fmt.Sprintf("%.2f", temp)
}

func (b *BME280) Pressure() string {
	press, err := b.i2cDriver.Pressure()
	if err != nil {
		press = 0
	}
	return fmt.Sprintf("%.2f", press)
}

func (b *BME280) Altitude() string {
	alt, err := b.i2cDriver.Altitude()
	if err != nil {
		alt = 0
	}
	return fmt.Sprintf("%.2f", alt)
}

func (b *BME280) Humidity() string {
	hum, err := b.i2cDriver.Humidity()
	if err != nil {
		hum = 0
	}
	return fmt.Sprintf("%.2f", hum)
}

func (b *BME280) Stat() BME280Metrics {
	m := BME280Metrics{
		Temperature: b.Temperature(),
		Pressure:    b.Pressure(),
		Altitude:    b.Altitude(),
		Humidity:    b.Humidity(),
	}
	return m
}
