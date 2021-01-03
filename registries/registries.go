package registries

import (
	"github.com/rlaskowski/go-iothome/datastore/memory"
	"github.com/rlaskowski/go-iothome/iot"
)

type Registries struct {
	DriverRepository *memory.DriverRepository
	RaspiDriver      *iot.Raspi
}

func NewRegistries() *Registries {
	return &Registries{
		DriverRepository: new(memory.DriverRepository),
		RaspiDriver:      iot.NewRaspi(),
	}
}
