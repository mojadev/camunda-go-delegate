module github.com/mojadev/camunda-go-delegate-example

go 1.15

replace github.com/mojadev/camunda-go-delegate => ../

replace github.com/mojadev/camunda-go-delegate/generate => ../generate

require (
	github.com/go-critic/go-critic v0.5.2 // indirect
	github.com/go-toolsmith/pkgload v1.0.1 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mojadev/camunda-go-delegate v0.0.0-20201127072206-e38235d6829e
	github.com/mojadev/camunda-go-delegate/generate v0.0.0-20201127072206-e38235d6829e
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.3
	github.com/quasilyte/go-ruleguard v0.2.1 // indirect
	github.com/quasilyte/regex/syntax v0.0.0-20200805063351-8f842688393c // indirect
	golang.org/x/tools v0.0.0-20201125231158-b5590deeca9b // indirect
)
