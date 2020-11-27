package api

// ExecutionContext provides all available information from the current process, e.g. variables.
type ExecutionContext struct {
	Variables VariableScope
}

// ExecutionResult describes the result of a process step execution and will be used to continue with the
// process.
type ExecutionResult struct {
	Variables VariableScope
	Signals   []Signal
}

// CreateExecutionResult creates a new ExecutionResult
func CreateExecutionResult() ExecutionResult {
	return ExecutionResult{
		Variables: make(VariableScope),
		Signals: make([]Signal, 0),
	}
}

// A signal that can be used to continue the process exeuction
type Signal struct {
	Name             MessageName
	ProcessVariables VariableScope
}

// A message that will be emitted in the BPMN
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

// HandlerFunc is a function for a delegate handler
type HandlerFunc = func(ctx *ExecutionContext) (ExecutionResult, error)

type ProcessStepEndpoint struct {
	Version VersionDescriptor
	Handler map[string]HandlerFunc
}