package server

import (
	"context"
	"fmt"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/routes"
	"net/http"
)

type WSServer struct {
	server *http.Server
}

func NewWSServer() *WSServer {
	return &WSServer{}
}

func (w *WSServer) start() {
	router := routes.NewWsRouter()

	w.server = &http.Server{
		Addr:           fmt.Sprintf(":%d", global.WsSetting.WsPort),
		Handler:        router,
		ReadTimeout:    global.WsSetting.HttpReadTimeout,
		WriteTimeout:   global.WsSetting.HttpWriteTimeout,
		MaxHeaderBytes: 1 << 32,
	}

	go func() {
		if global.WsSetting.IsTLS {
			if err := w.server.ListenAndServeTLS(global.WsSetting.TLSCertFile, global.WsSetting.TLSKeyFile); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		} else {
			if err := w.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	fmt.Printf("websocket[:%d] success \n", global.WsSetting.WsPort)
}

func (w *WSServer) Run() {
	w.start()
}

func (w *WSServer) Shutdown(ctx context.Context) error {
	fmt.Printf("websocket[:%d] shutdown \n", global.WsSetting.WsPort)

	if err := w.server.Shutdown(ctx); err != nil {
		return err
	}

	fmt.Printf("websocket[:%d] exit \n", global.WsSetting.WsPort)
	return nil
}
