package main

import (
	"github.com/gin-gonic/gin"
	"github.com/openapi/api/strictserver"
	"log"
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

	server := strictserver.NewServer()

	router := gin.Default()
	sh := strictserver.NewStrictHandler(server, nil)
	strictserver.RegisterHandlers(router, sh)
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
