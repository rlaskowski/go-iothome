package main

import (
	"github.com/rlaskowski/go-iothome/service"
)

func main() {

	service := service.New()
	service.Start()

}
