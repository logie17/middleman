package middleman

import (
	"net/http"
)


type wrapResponseWriter struct {
	http.ResponseWriter
	status int
}

// Writeheader Wraps ResponseWriters WriteHeader to capture status code
func (w *wrapResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

type handler func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)

var registeredHandlers []handler

// Register a set of middleware handlers
func Register(h ...handler) {
	registeredHandlers = append(registeredHandlers, h...)
}


// Wrap will take a handlerfn and wrap it
func Wrap(handlerFn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	for i := len(registeredHandlers) - 1; i >= 0; i-- {
		handlerFn = registeredHandlers[i](handlerFn)
	}
	return handlerFn
}
