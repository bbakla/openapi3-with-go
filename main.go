package main

import (
	"embed"
	not_strict_server "github.com/bbakla/openapi3-with-go/oapi-codegen/not-strict-server"
	"github.com/bbakla/openapi3-with-go/oapi-codegen/strictserver"
	openapigengin "github.com/bbakla/openoapi-code-generator/oapi_generator_userapi"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"

	"github.com/bbakla/openapi3-with-go/openapi-generator/openapigenonlyinterface"
	"log"
)

//go:embed swagger-ui
var swaggerContent embed.FS

func main() {

	/*	http.Handle("/api1/docs/", v5emb.New(
			"Petstore",
			"https://petstore3.swagger.io/api/v3/openapi.json",
			"/api1/docs/",
		))



		println("docs at http://localhost:8080/api1/docs/")
		_ = http.ListenAndServe("localhost:8080", http.DefaultServeMux)*/

	openpiGenerator()

	//oapicodegenNonStrict()

	//oapiCodegenStrict()
}

// Using openapi-generator generated code
func openpiGenerator() {

	routes := openapigengin.ApiHandleFunctions{UserAPI: openapigenonlyinterface.NewUserAPI()}
	log.Printf("Server started")
	router := openapigengin.NewRouter(routes)

	// add self hosted swagger ui
	fsys, _ := fs.Sub(swaggerContent, "swagger-ui")
	router.StaticFS("/swagger", http.FS(fsys))

	log.Fatal(router.Run(":8080"))
}

// using  oapi-codegen strict server version
func oapiCodegenStrict() {
	server := strictserver.NewServer()

	router := gin.Default()
	// add self hosted swagger ui
	fsys, _ := fs.Sub(swaggerContent, "swagger-ui")
	router.StaticFS("/swagger", http.FS(fsys))

	sh := strictserver.NewStrictHandler(server, nil)
	strictserver.RegisterHandlers(router, sh)
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}

// using oapi-codegen not strict server version
func oapicodegenNonStrict() {

	server := not_strict_server.NewServer()

	router := gin.Default()
	// add swagger ui
	fsys, _ := fs.Sub(swaggerContent, "swagger-ui")
	router.StaticFS("/swagger", http.FS(fsys))

	not_strict_server.RegisterHandlers(router, server)
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
