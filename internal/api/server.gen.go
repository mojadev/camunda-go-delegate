// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /{handler}/execution)
	RetrieveProcessStepInfo(ctx echo.Context, handler string) error

	// (POST /{handler}/execution)
	ExecuteProcessStep(ctx echo.Context, handler string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// RetrieveProcessStepInfo converts echo context to params.
func (w *ServerInterfaceWrapper) RetrieveProcessStepInfo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "handler" -------------
	var handler string

	err = runtime.BindStyledParameter("simple", false, "handler", ctx.Param("handler"), &handler)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter handler: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RetrieveProcessStepInfo(ctx, handler)
	return err
}

// ExecuteProcessStep converts echo context to params.
func (w *ServerInterfaceWrapper) ExecuteProcessStep(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "handler" -------------
	var handler string

	err = runtime.BindStyledParameter("simple", false, "handler", ctx.Param("handler"), &handler)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter handler: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ExecuteProcessStep(ctx, handler)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/:handler/execution", wrapper.RetrieveProcessStepInfo)
	router.POST(baseURL+"/:handler/execution", wrapper.ExecuteProcessStep)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xWyY7jNhD9lQITIMlAkD1ZLj4lmTSQPgzQaBu5DOZAkyWbHW7h4ozR0L8HRcptyVYc",
	"IOiZm7hVvfdq0zMTznhn0abIVs8sij0aXj7vQnDhEaN3NiJt+OA8hqSwHHMpVVLOcn1vOxcMpwUdqITm",
	"4sbD6GUKGRuWjh7ZirntE4rE+pcNHgI/0tpgjHxX/A5HMQVld6zvZ14/BCcwxnfOJvyUrsEeeFB8q+vi",
	"64AdW7GvFmfqi4H34o/h4lo4jzd93X1CkYneI8asZ3xGtbNcl0+JUQTlq0JsXQ+AWwl4IPeQ9jxB3Lus",
	"JWwR0KiUUIKykPYIvnpkzVnbWxyq/TlRP4MM64S+SuHCb5i4qoynUkz4XwW0YQcMcf6sb1jAv7IKKNnq",
	"w8vFjzOABtpXzv/E46zTf8+wk/HLwG32CHQCroMa3cvARbSyhVI5Ed7U/TfA9d/8GEtgi1IoQfOYSgKY",
	"HBM4q4/ghMgBnBXYsoahzYYYvx9QvvBramGOJBjp+P/jO1a52J2TePrsP0p8qh2VJlcWA1Rz0LlwSmw4",
	"475yStiU7Rx5E2RDlFJDw5VmK2bcE5d4+HlH61Y4QzAvo6YiKCo3yNZ1nRJcw++bzQOgld4pW7EIbrKV",
	"HCKGgxIFilYCh95nuSFU7+83JT1U0rR8N7x5ROMSwt1gj40ymr1tl+2SHjmPlnvFVuyHdtm+ZQ3zPO1L",
	"qBbPe26lxtAv8NRWaH+H6ToLHzEFhQcEde67wLcup3GzgJjQgzJeo0GbancuIEL5vpcjU6NapmZeoAVu",
	"MGGIbPXhmSlyTHBZc5JiQMyaYWbMFe9Hyqs6QArP75fLUxzRFmrce61EQbR4ipX12d6tFL7Rf0rSTEUb",
	"bsO6qDIaWOWud3FG6HXeGpWAw0tMgKoEY00YEnunDmgnkl9pXMGNJX5deQugX508vrayp2HaT9sDFXf/",
	"+eN6OV5nYrrOgm52WY8i5LqrKqDi+/EVEU5/jGaAbUYABI0A+02adH+ZEZIDZWOmfqTQJhBVbvgW210L",
	"RsWo7O7cGb8jFj99SRa/WEC6UwdTQU15RyyyONdOzMbwcKRMH7fTUFvipBudo1TGpuCWVBFca5TQBWcm",
	"fzt93/8TAAD//24DUraeCgAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}