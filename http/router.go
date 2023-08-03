package http

import "net/http"

// Router main router interface, can be implemented with libraries like mux or another router
type Router interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request))
	HandleFuncWithMethod(path, method string, f func(http.ResponseWriter, *http.Request))
}
