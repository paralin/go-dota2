module github.com/paralin/go-dota2

go 1.18

// note: protobuf is intentionally held at 1.3.x
replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.5

require (
	github.com/faceit/go-steam v0.0.0-20190206021251-2be7df6980e1
	github.com/fatih/camelcase v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/pkg/errors v0.9.1
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	github.com/sirupsen/logrus v1.9.0
	github.com/urfave/cli v1.22.10
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.22.1 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)
