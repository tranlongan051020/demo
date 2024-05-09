package router

import (
	"net/http"

	"demo/api/controller"
)

type routeType struct {
	method  string
	path    string
	handler func(http.ResponseWriter, *http.Request)
}

type Router struct {
	allCtrl *controller.Controller
}

func New(allCtrl *controller.Controller) *Router {
	return &Router{
		allCtrl: allCtrl,
	}
}
