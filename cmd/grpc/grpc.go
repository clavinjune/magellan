package grpc

import (
	"github.com/clavinjune/gokit/cobrautil"
	"github.com/clavinjune/gokit/grpcutil"
	magellanv1 "github.com/clavinjune/magellan/api/proto/magellan/v1"
	"github.com/clavinjune/magellan/pkg/component"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

var Cmd = &cobra.Command{
	Use:   "grpc",
	Short: "run biko grpc server",
}

func Init(cmd *cobra.Command) *cobra.Command {
	cmd.RunE = RunE

	cmd.Flags().IntP("grpc-port", "g", 12999, "gRPC port")
	return cmd
}

func RunE(cmd *cobra.Command, args []string) error {
	wire, cleanup, err := component.Wire(component.Option{
		Logger:       slog.Default(),
		Port:         grpcPort(cmd),
		ServerOption: nil,
	})
	if err != nil {
		return err
	}
	defer cleanup()

	wire.Server.Register(func(server *grpc.Server) {
		magellanv1.RegisterAuthenticationServiceServer(server, wire.Authentication)
	}).Listen(cmd.Context())
	return nil
}

func grpcPort(cmd *cobra.Command) grpcutil.ServerPort {
	return grpcutil.ServerPort(cobrautil.MustGetIntFlag(cmd, "grpc-port"))
}
