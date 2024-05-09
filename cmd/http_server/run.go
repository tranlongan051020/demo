package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"demo/api/router"
	"demo/config"
)

type Server struct {
	r *router.Router
}

func New(r *router.Router) *Server {
	return &Server{
		r: r,
	}
}

func (sv *Server) RunHTTPServer() {
	//  init router and register all routes
	sv.r.RegisterAllRoutes()

	// init server
	server := &http.Server{
		Addr:         ":" + config.Server.Port,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
	}

	// run server
	fmt.Println("HTTP server running on port", config.Server.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

}
