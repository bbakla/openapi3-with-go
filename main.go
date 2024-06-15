package main

import (
	not_strict_server "github.com/bbakla/openapi3-with-go/oapi-codegen/not-strict-server"
	openapigengin "github.com/bbakla/openapi3-with-go/openapi-generator/gin-server-gen/oapi-go-codegen"
	"github.com/bbakla/openapi3-with-go/openapi-generator/openapigenonlyinterface"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	// Using openapi-generator generated code
	routes := openapigengin.ApiHandleFunctions{UserAPI: openapigenonlyinterface.NewUserAPI()}
	log.Printf("Server started")
	router := openapigengin.NewRouter(routes)
	log.Fatal(router.Run(":8080"))

	// using oapi-codegen not strict server version

	server := not_strict_server.NewServer()

	router := gin.Default()
	not_strict_server.RegisterHandlers(router, server)
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())

	// using  oapi-codegen strict server version

	/*server := strictserver.NewServer()

	router := gin.Default()
	sh := strictserver.NewStrictHandler(server, nil)
	strictserver.RegisterHandlers(router, sh)
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())*/
}
