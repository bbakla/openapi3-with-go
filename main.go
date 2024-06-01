package main

import (
	"github.com/gin-gonic/gin"
	strict_server "github.com/openapi/api/strict-server"
	"net/http"
)

func main() {
	/*
		server := not_strict_server.NewServer()

		router := gin.Default()
		not_strict_server.RegisterHandlers(router, server)
		s := &http.Server{
			Handler: router,
			Addr:    "0.0.0.0:8080",
		}*/

	server := strict_server.NewServer()

	router := gin.Default()
	sh := strict_server.NewStrictHandler(server, nil)
	strict_server.RegisterHandlers(router, sh)
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
