module github.com/paralin/go-dota2

go 1.19

// note: protobuf is intentionally held at 1.3.x
replace github.com/golang/protobuf => github.com/aperturerobotics/go-protobuf-1.3.x v0.0.0-20200726220404-fa7f51c52df0 // aperture-1.3.x

require (
	github.com/fatih/camelcase v1.0.0
	github.com/golang/protobuf v1.5.3
	github.com/paralin/go-steam v0.0.0-20230323124938-41e2a40a56ff
	github.com/pkg/errors v0.9.1
	github.com/serenize/snaker v0.0.0-20201027110005-a7ad2135616e
	github.com/sirupsen/logrus v1.9.0
	github.com/urfave/cli v1.22.12
	github.com/urfave/cli/v2 v2.25.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.22.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)
