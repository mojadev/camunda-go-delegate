package main

const TEMPLATE = `
//  ^// Code generated .* DO NOT EDIT\.$
package main

import (
	"flag"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/mojadev/camunda-go-delegate/api"
	{{- range $delegate := .Delegates }}
	{{- if ne $delegate.Package "main"}}
	"{{ $.Module }}/{{ $delegate.Package }}"
	{{- end }}
	{{- end }}
)

const VERSION = "{{ .Version }}"
const DESCRIPTION = "{{ .Description }}"

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()
	delegates := make(map[string]api.HandlerFunc)
	{{- range $delegate := .Delegates }}
	{{- if eq $delegate.Package "main"}}
	delegates["{{ $delegate.Name }}"] = {{ $delegate.Handler }}
	{{- else }}
	delegates["{{ $delegate.Name }}"] = {{ $delegate.Package}}.{{ $delegate.Handler }}
	{{- end }}
	{{ end }}
	
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
`