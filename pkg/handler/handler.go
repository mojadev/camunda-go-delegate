package handler

import (
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	remoteApi "github.com/mojadev/camunda-go-delegate/internal/api"
	"github.com/mojadev/camunda-go-delegate/pkg/api"
)

func RegisterHandlers(router remoteApi.EchoRouter, si api.ProcessStepEndpoint) {

	remoteApi.RegisterHandlers(router, remoteApi.ProcessStepEndpoint{
		Handler: si.Handler,
		Version: si.Version,
	})
}

func SetupServer(endpoint api.ProcessStepEndpoint) *echo.Echo {
	openapiSpec, _ := remoteApi.GetSwagger()
	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(openapiSpec))
	RegisterHandlers(e, endpoint)
	return e
}