package server

import (
	"context"
	"fmt"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/routes"
	"net/http"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (h *HttpServer) start() {
	router := routes.NewHttpRouter()

	h.server = &http.Server{
		Addr:           fmt.Sprintf(":%d", global.HttpSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.HttpSetting.HttpReadTimeout,
		WriteTimeout:   global.HttpSetting.HttpWriteTimeout,
		MaxHeaderBytes: 1 << 32,
	}

	go func() {
		if global.HttpSetting.IsTLS {
			if err := h.server.ListenAndServeTLS(global.HttpSetting.TLSCertFile, global.HttpSetting.TLSKeyFile); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		} else {
			if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	fmt.Printf("http[:%d] success \n", global.HttpSetting.HttpPort)
}

func (h *HttpServer) Run() {
	h.start()
}

func (h *HttpServer) Shutdown(ctx context.Context) error {
	fmt.Printf("http[:%d] shutdown \n", global.HttpSetting.HttpPort)

	if err := h.server.Shutdown(ctx); err != nil {
		return err
	}

	fmt.Printf("http[:%d] exit \n", global.HttpSetting.HttpPort)
	return nil
}
