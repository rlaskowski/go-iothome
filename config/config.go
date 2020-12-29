package config

import (
	"time"
)

var (
	ExecutableName         = "iothome"
	HttpServerPort         = 8080
	HttpServerReadTimeout  = 30 * time.Second
	HttpServerWriteTimeout = 30 * time.Second
)
