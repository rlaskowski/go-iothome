package iot

import (
	"github.com/rlaskowski/go-iothome/iot/drivers"
	"gobot.io/x/gobot/platforms/raspi"
)

type Raspi struct {
	adaptor *raspi.Adaptor
}

func NewRaspi() *Raspi {
	return &Raspi{
		adaptor: raspi.NewAdaptor(),
	}
}

func (r *Raspi) BME280Driver() (*drivers.BME280, error) {
	driver := drivers.NewBME280(r.adaptor)
	if err := driver.Start(); err != nil {
		return nil, err
	}
	return driver, nil
}
