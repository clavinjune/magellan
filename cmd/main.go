package main

import (
	"fmt"
	"github.com/clavinjune/magellan/cmd/grpc"
	"github.com/clavinjune/magellan/cmd/proxy"
	"runtime"

	"github.com/clavinjune/gokit/cobrautil"
	"github.com/spf13/cobra"
)

const (
	AppName = "magellan"
)

var (
	Version = "dev"
)

func main() {
	version := fmt.Sprintf("%s %s/%s", Version, runtime.GOOS, runtime.GOARCH)
	root := cobrautil.New(AppName, version)
	root.AddCommand(
	//		hello.Init(hello.Cmd),
	//		playground.Init(playground.Cmd),
			grpc.Init(grpc.Cmd),
			proxy.Init(proxy.Cmd),
	)

	cobra.CheckErr(root.Execute())
}
