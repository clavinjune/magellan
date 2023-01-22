package main

import (
	"fmt"

	"github.com/clavinjune/magellan/cmd/grpc"
	"github.com/clavinjune/magellan/cmd/proxy"

	"github.com/clavinjune/gokit/cobrautil"
	"github.com/spf13/cobra"
)

const (
	AppName = "magellan"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	version := fmt.Sprintf("%s commit=%s build_date=%s", version, commit, date)
	root := cobrautil.New(AppName, version)
	root.AddCommand(
		//		hello.Init(hello.Cmd),
		//		playground.Init(playground.Cmd),
		grpc.Init(grpc.Cmd),
		proxy.Init(proxy.Cmd),
	)

	cobra.CheckErr(root.Execute())
}
