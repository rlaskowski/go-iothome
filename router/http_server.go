package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/rlaskowski/go-iothome/config"
)

type HttpServer struct {
	server  *http.Server
	mux     *http.ServeMux
	context context.Context
	cancel  context.CancelFunc
}

func NewHttpServer() *HttpServer {
	ctx, cancel := context.WithCancel(context.Background())

	mux := http.NewServeMux()

	return &HttpServer{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", config.HttpServerPort),
			ReadTimeout:  config.HttpServerReadTimeout,
			WriteTimeout: config.HttpServerWriteTimeout,
			Handler:      mux,
		},
		mux:     mux,
		context: ctx,
		cancel:  cancel,
	}
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
	h.mux.HandleFunc("/login", h.Login())
}

func (h *HttpServer) Login() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(r.Host))
	}
}
