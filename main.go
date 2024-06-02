package main

import (
	openapigen_gin "github.com/bbakla/openapi3-with-go/open-oapi-codegen"
	"github.com/bbakla/openapi3-with-go/openapi-generator/openapigenonlyinterface"
	"log"
)

func main() {

	// Using openapi-generator generated code
	routes := openapigen_gin.ApiHandleFunctions{UserAPI: openapigenonlyinterface.NewUserAPI()}
	log.Printf("Server started")
	router := openapigen_gin.NewRouter(routes)
	log.Fatal(router.Run(":8080"))

	// using oapi-codegen not strict server version
	/*
			server := not_strict_server.NewServer()

			router := gin.Default()
			not_strict_server.RegisterHandlers(router, server)
			s := &http.Server{
				Handler: router,
				Addr:    "0.0.0.0:8080",
			}

		log.Fatal(s.ListenAndServe())
	*/

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
