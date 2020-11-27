module github.com/mojadev/camunda-go-delegate-example

go 1.15

replace github.com/mojadev/camunda-go-delegate => ../

replace github.com/mojadev/camunda-go-delegate/generate => ../generate

require (
	github.com/mojadev/camunda-go-delegate v0.0.0-20201127072206-e38235d6829e // indirect
	github.com/mojadev/camunda-go-delegate/generate v0.0.0-20201127072206-e38235d6829e
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.3
)
