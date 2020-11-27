
//  ^// Code generated .* DO NOT EDIT\.$
package main

import (
	"flag"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/mojadev/camunda-go-delegate/example/delegates"
	"github.com/mojadev/camunda-go-delegate/internal/api"
)

const VERSION = "1.0.0"
const DESCRIPTION = "Convert currency"

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()
	delegates := make(map[string]api.HandlerFunc)
	delegates["convert-currency"] = delegates.ConvertCurrency
	
	
	endpoint := api.ProcessStepEndpoint{
		Version: api.VersionDescriptor{
			Version: VERSION,
			Description: DESCRIPTION,
		},
		Handler: delegates,
	}

	openapiSpec, _ := api.GetSwagger()
	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(openapiSpec))
	api.RegisterHandlers(e, endpoint)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))

}
