package router

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/rlaskowski/go-iothome/config"
	"github.com/rlaskowski/go-iothome/registries"
)

type HttpServer struct {
	server     *http.Server
	router     *pat.PatternServeMux
	context    context.Context
	cancel     context.CancelFunc
	registries *registries.Registries
}

type Router struct {
	method string
	host   string
	name   http.HandlerFunc
}

func NewHttpServer(registries *registries.Registries) *HttpServer {
	ctx, cancel := context.WithCancel(context.Background())

	h := &HttpServer{
		context:    ctx,
		cancel:     cancel,
		router:     pat.New(),
		registries: registries,
	}

	h.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", config.HttpServerPort),
		ReadTimeout:  config.HttpServerReadTimeout,
		WriteTimeout: config.HttpServerWriteTimeout,
		Handler:      h,
	}

	return h
}

func (h *HttpServer) Start() error {
	go func() {
		h.configureEndpoints()

		log.Printf("Starting REST API on http://localhost:%d", config.HttpServerPort)

		if err := h.server.ListenAndServe(); err != nil {
			log.Fatalf("Caught error while starting server: %s", err.Error())
		}
	}()

	return nil
}

func (h *HttpServer) Stop() error {
	h.cancel()

	log.Print("Stopping REST API")

	return h.server.Close()
}

func (h *HttpServer) configureEndpoints() {
	h.Get("/bme280", h.BME280)
	h.Get("/drivers", h.Drivers)
}

func (h *HttpServer) Get(path string, handler http.HandlerFunc) {
	h.router.Get(path, handler)
}

func (h *HttpServer) Post(path string, handler http.HandlerFunc) {
	h.router.Post(path, handler)
}

func (h *HttpServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(rw, r)
}

func (h *HttpServer) BME280(rw http.ResponseWriter, r *http.Request) {
	driver, err := h.registries.RaspiDriver.BME280Driver()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	b, err := json.Marshal(driver.Stat())
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	rw.Header().Add("Content-Type", "application/json")

	rw.Write(b)
}

func (h *HttpServer) Drivers(rw http.ResponseWriter, r *http.Request) {
	drivers, err := json.Marshal(h.registries.DriverRepository.FindAll())
	if err != nil {
		rw.WriteHeader(http.StatusNoContent)
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(drivers)
}

func (h *HttpServer) json(rw http.ResponseWriter, i interface{}) {
	rw.Header().Add("Content-Type", "application/json")

	marshal, err := json.Marshal(i)
	if err != nil {
		rw.Write([]byte(err.Error()))
	}
	rw.Write(marshal)
}
