module github.com/paralin/go-dota2

go 1.12

// note: protobuf is intentionally held at 1.3.x
replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.5

require (
	github.com/faceit/go-steam v0.0.0-20190206021251-2be7df6980e1
	github.com/fatih/camelcase v1.0.0
	github.com/golang/protobuf v1.4.1
	github.com/onsi/ginkgo v1.9.0 // indirect
	github.com/onsi/gomega v1.6.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	github.com/sirupsen/logrus v1.4.2
	github.com/urfave/cli v1.21.0
	google.golang.org/protobuf v1.24.0 // indirect
)
