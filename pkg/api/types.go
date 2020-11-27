package api

type ExecutionContext struct {
	Variables VariableScope
}

type ExecutionResult struct {
	Variables VariableScope
	Signals   []Signal
}

func CreateExecutionResult() ExecutionResult {
	return ExecutionResult{
		Variables: make(VariableScope),
		Signals: make([]Signal, 1),
	}
}

type Signal struct {
	Name             MessageName
	ProcessVariables VariableScope
}

type ProcessMessage struct {
	Name             MessageName
	BusinessKey      BusinessKey
	CorrelationKeys  VariableScope
	ProcessVariables VariableScope
}


type VariableScope = map[VariableKey]VariableValue

type VariableKey = string
type VariableValue = interface{}

type MessageName = string
type BusinessKey = *string

type SignalName = string
type SignalData = interface{}

type VersionDescriptor struct {
	Description string
	Version string
}

type HandlerFunc = func(ctx *ExecutionContext) (ExecutionResult, error)

type ProcessStepEndpoint struct {
	Version VersionDescriptor
	Handler map[string]HandlerFunc
}