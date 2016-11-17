package main

import (
	"net/http"
)

// Route for the router
type Route struct {
	Name    string
	Path    string
	Handler func(x http.ResponseWriter, r *http.Request)
}

// Router interface to support routing
type Router interface {
	GetRoutes() []Route
	AddRoute(route Route)
	RemoveRoute(name string)
}

type defaultRouter struct {
	Routes map[string]Route
}

// NewRouter creates a new router
func NewRouter() Router {
	return &defaultRouter{Routes: make(map[string]Route)}
}

// Bind binds route handlers to the http.HandlerFunc
func Bind(r Router) {
	for _, rt := range r.GetRoutes() {
		http.HandleFunc(rt.Path, rt.Handler)
	}
}

// GetRoutes returns a list of routes
func (r *defaultRouter) GetRoutes() []Route {
	rts := make([]Route, len(r.Routes))

	count := 0
	for _, rt := range r.Routes {
		rts[count] = rt
		count++
	}

	return rts
}

func (r *defaultRouter) AddRoute(route Route) {
	r.Routes[route.Name] = route
}

func (r *defaultRouter) RemoveRoute(name string) {
	delete(r.Routes, name)
}
