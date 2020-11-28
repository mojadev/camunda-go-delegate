package main

const TEMPLATE = `
//  ^// Code generated .* DO NOT EDIT\.$
package main

import (
	"flag"
	"fmt"
	"github.com/mojadev/camunda-go-delegate/api"
	"github.com/mojadev/camunda-go-delegate/handler"
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
	e := handler.SetupServer(endpoint)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
`