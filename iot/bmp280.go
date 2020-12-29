package iot

import (
	"fmt"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

type BMP280 struct {
	i2cDriver *i2c.BMP280Driver
}

func NewBMP280() *BMP280 {
	raspiAdaptor := raspi.NewAdaptor()

	return &BMP280{
		i2cDriver: i2c.NewBMP280Driver(raspiAdaptor),
	}
}

func (b *BMP280) String() string {
	return b.i2cDriver.Name()
}

func (b *BMP280) Temperature() string {
	temp, err := b.i2cDriver.Temperature()
	if err != nil {
		temp = 0
	}
	return fmt.Sprintf("%.2f", temp)
}
