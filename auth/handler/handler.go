package handler

import (
	"net/http"

	"github.com/ravenxwrx/pencilpusher-sdk/model"
)

type Handler struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

var _ model.Handler = (*Handler)(nil)

func (h *Handler) GetPath() string {
	return h.Path
}

func (h *Handler) GetMethod() string {
	return h.Method
}

func (h *Handler) GetHandlerFunc() http.HandlerFunc {
	return h.HandlerFunc
}

func IndexHandler() *Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World but in auth!"))
	}

	return &Handler{
		Path:        "",
		Method:      http.MethodGet,
		HandlerFunc: fn,
	}
}
