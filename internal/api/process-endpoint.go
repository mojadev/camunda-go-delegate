//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./types.cfg.yaml ../../endpoint-api.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./server.cfg.yaml ../../endpoint-api.yaml

package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mojadev/camunda-go-delegate/api"
	"net/http"
)

type VersionDescriptor struct {
	Description string
	Version string
}

type HandlerFunc = func(ctx *api.ExecutionContext) (api.ExecutionResult, error)

type ProcessStepEndpoint struct {
	Version VersionDescriptor
	Handler map[string]HandlerFunc
}

func (p ProcessStepEndpoint) RetrieveProcessStepInfo(ctx echo.Context, handler string) error {
	var result ProcessStepExecutorDetails
	result.Description = &p.Version.Description
	result.Version = p.Version.Version

	if p.Handler[handler] == nil {
		errMessage := "Unknown handler " + handler
		return ctx.JSON(400, ErrorResponse{Message: &errMessage })
	}
	return ctx.JSON(http.StatusOK, result)
}

func (p ProcessStepEndpoint)  ExecuteProcessStep(ctx echo.Context, handler string) error {
	var context ProcessContext
	var result ProcessExecutionResult

	err := ctx.Bind(&context)
	if err != nil {
		return err
	}
	if p.Handler[handler] == nil {
		errMessage := "Unknown handler " + handler
		return ctx.JSON(400, ErrorResponse{Message: &errMessage })
	}

	apiContext := api.ExecutionContext{Variables: *mapFromDto(context.Variables)}
	callResult, err := p.Handler[handler](&apiContext)

	if err != nil {
		errMessage := err.Error()
		return ctx.JSON(500, ErrorResponse{Message: &errMessage})
	}
	result.Variables = mapToDto(&callResult.Variables)
	return ctx.JSON(http.StatusOK, result)
}

func mapFromDto(dto *VariableScope) *api.VariableScope {
	result := make(api.VariableScope)
	for key, element := range dto.AdditionalProperties {
		result[key] = element
	}
	return &result
}

func mapToDto(domain *api.VariableScope) *VariableScope {
	result := VariableScope{
		AdditionalProperties: make(map[string]interface{}),
	}
	for key, element := range *domain {
		result.AdditionalProperties[key] = element
	}
	return &result
}