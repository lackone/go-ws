package client

import "sync"

var WsClientHandler = NewClientHandler()

type HandlerFunc func(client *Client, data []byte) ClientResponse

type ClientHandler struct {
	handlers map[string]HandlerFunc
	lock     sync.RWMutex
}

func NewClientHandler() *ClientHandler {
	return &ClientHandler{
		handlers: make(map[string]HandlerFunc),
		lock:     sync.RWMutex{},
	}
}

func (h *ClientHandler) Register(key string, fn HandlerFunc) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.handlers[key] = fn
}

func (h *ClientHandler) UnRegister(key string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.handlers, key)
}

func (h *ClientHandler) GetHandler(key string) (HandlerFunc, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	fn, ok := h.handlers[key]
	return fn, ok
}
