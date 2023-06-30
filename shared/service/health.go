package service

import (
	"IofIPOS/shared/httpext"
	"IofIPOS/shared/stringsext"
	"net/http"
)

func NewHealthHandler() *healthHandler {
	return new(healthHandler)
}

type healthHandler struct {
}

func (*healthHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/ping":
		res.Header().Set(httpext.ContentTypeHeader, stringsext.MIME)
		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte("pong"))
	default:
		res.Header().Set(httpext.ContentTypeHeader, stringsext.MIME)
		res.WriteHeader(http.StatusNotFound)
		_, _ = res.Write([]byte("not found"))
	}
}
