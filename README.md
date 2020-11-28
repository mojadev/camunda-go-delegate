# Delegate Generator (draft)

> This is a first, pre-release version and could change

## About 

This is a generator that can be used to turn a go module containing delegate functions
into a container that can be called from the [camunda-delegate](#LinkMissing).

Check out the [example](../../example) folder for an example how to use it. 

## Usage 

### General usage

`go run github.com/mojadev/camunda-go-delegate/generate --config {config} --output {output}`

Options:
- _config_="":   
  The [config file](#config-file) used for generating the delegate endpoint.
  Example: `./delegate.yaml`
- _output_="main.go":  
  The path of the entry file to generate 

### Using go generate 

Add the following line on top of your go file

`//go:generate go run github.com/mojadev/camunda-go-delegate/generate --config ../delegate.yaml --output ../main.go`

### Adding a delegate configuration

In order to know which functions should be exported, a simple `yaml` must be created.
This might not be necessary in the future, but for now this makes generation more explicit.

```yaml
name: Short name describing your delegate container
version: The version of your delegate container 
description: (optional) Longer description of your delegate container
module: github.com/<your org>/<your module name>
delegates:
  - filename: my-delegate.go
    function: MyFunctionName
    export: my-function
  - filename: my-other.delegate.go
    function: MyOtherFunction
    export: my-other-function
```
This will generate two endpoints:  
- `/my-function/execution` executing the `MyFunctionName` function in `my-delegate.go`
- `/my-other-function/execution` executing the `MyOtherFunctionName` function in `my-other-delegate.go`

