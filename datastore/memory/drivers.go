package memory

import (
	"strings"

	"github.com/rlaskowski/go-iothome/model"
)

var drivers = []model.Driver{
	{Id: "bme280", Name: "BME280", Group: "I2C"},
	{Id: "grovelcd", Name: "Grove LCD", Group: "I2C"},
}

type DriverRepository struct {
}

func (d *DriverRepository) FindAll() []model.Driver {
	return drivers
}

func (d *DriverRepository) FindByGroup(group string) []model.Driver {
	var list []model.Driver

	for _, driver := range drivers {
		if strings.Compare(driver.Group, group) == 0 {
			list = append(list, driver)
		}
	}
	return list
}
