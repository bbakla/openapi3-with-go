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

	openpiGenerator()

	//oapicodegenNonStrict()

	//oapiCodegenStrict()
}

// Using openapi-generator generated code
func openpiGenerator() {

	routes := openapigengin.ApiHandleFunctions{UserAPI: openapigenonlyinterface.NewUserAPI()}
	log.Printf("NonStrictServer started")
	router := openapigengin.NewRouter(routes)

	addSwaggerEndpoint(router)

	log.Fatal(router.Run(":8080"))
}

// using  oapi-codegen strict server version
func oapiCodegenStrict() {
	server := strictserver.NewServer()

	router := gin.Default()
	addSwaggerEndpoint(router)

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
	addSwaggerEndpoint(router)

	not_strict_server.RegisterHandlers(router, server)
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}

func addSwaggerEndpoint(router *gin.Engine) {
	// add swagger ui
	fsys, _ := fs.Sub(swaggerContent, "swagger-ui")
	router.StaticFS("/swagger", http.FS(fsys))
}
