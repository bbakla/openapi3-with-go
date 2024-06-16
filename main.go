package main

import (
	"embed"
	not_strict_server "github.com/bbakla/openapi3-with-go/oapi-codegen/not-strict-server"
	"github.com/bbakla/openapi3-with-go/oapi-codegen/strictserver"
	"github.com/gin-gonic/gin"
	"io/fs"

	//openapigengin "github.com/bbakla/openapi3-with-go/openapi-generator/generated/oapi_generator_userapi"
	openapigengin "github.com/bbakla/openoapi-code-generator/oapi_generator_userapi"
	"net/http"

	//openapigengin "github.com/bbakla/openapi3-with-go/openapi-generator/go-gen-gin/oapi_generator_userapi"
	//openapigengin "github.com/bbakla/openapi3-with-go/openapi-generator/gin-server-gen/oapi-go-codegen"
	"github.com/bbakla/openapi3-with-go/openapi-generator/openapigenonlyinterface"
	"log"
)

//go:embed swagger-ui
var content embed.FS

func main() {

	openpiGenerator()

	//oapicodegenNonStrict()

	//oapiCodegenStrict()
}

// using  oapi-codegen strict server version
func oapiCodegenStrict() {
	server := strictserver.NewServer()

	router := gin.Default()
	// add swagger ui
	fsys, _ := fs.Sub(content, "swagger-ui")
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
	fsys, _ := fs.Sub(content, "swagger-ui")
	router.StaticFS("/swagger", http.FS(fsys))

	not_strict_server.RegisterHandlers(router, server)
	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}

// Using openapi-generator generated code
func openpiGenerator() {

	routes := openapigengin.ApiHandleFunctions{UserAPI: openapigenonlyinterface.NewUserAPI()}
	log.Printf("Server started")
	router := openapigengin.NewRouter(routes)

	// add swagger-ui
	fsys, _ := fs.Sub(content, "swagger-ui")
	router.StaticFS("/swagger", http.FS(fsys))

	log.Fatal(router.Run(":8080"))
}
