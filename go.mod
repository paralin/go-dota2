module github.com/paralin/go-dota2

go 1.17

replace github.com/Philipp15b/go-steam/v3/v3 => github.com/AdmiralBulldogTv/go-steam/v3 v3.0.0-20220301050433-6b949b091037

require (
	github.com/Philipp15b/go-steam/v3 v3.0.0
	github.com/fatih/camelcase v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	github.com/sirupsen/logrus v1.4.2
	github.com/urfave/cli v1.21.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/onsi/ginkgo v1.9.0 // indirect
	github.com/onsi/gomega v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894 // indirect
)
