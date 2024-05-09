package router

import (
	"fmt"
	"net/http"
)

func (r *Router) RegisterAllRoutes() {
	allRoutes := [][]routeType{
		r.userRoutes(),
	}

	fmt.Println("================= API =================")
	for _, domainRoutes := range allRoutes {
		for _, route := range domainRoutes {
			pattern := fmt.Sprintf("%s %s", route.method, route.path)
			fmt.Println(pattern)
			http.HandleFunc(pattern, route.handler)
		}
	}
	fmt.Println("=======================================")
}
