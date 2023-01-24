package proxy

import (
	"context"
	"net/http"
	"time"

	"github.com/clavinjune/magellan/gen"

	authenticationv1 "github.com/clavinjune/magellan/api/proto/magellan/authentication/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/clavinjune/gokit/grpcutil"
	"github.com/clavinjune/magellan/pkg/component"
	"golang.org/x/exp/slog"

	"github.com/clavinjune/gokit/cobrautil"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "proxy",
	Short: "run magellan proxy server",
}

func Init(cmd *cobra.Command) *cobra.Command {
	cmd.RunE = RunE

	cmd.Flags().IntP("grpc-port", "g", 0, "gRPC port (default random)")
	cmd.Flags().IntP("proxy-port", "p", 8000, "proxy port")
	return cmd
}

func RunE(cmd *cobra.Command, args []string) error {
	wire, cleanup, err := component.Wire(component.Option{
		Logger:       slog.Default(),
		Port:         grpcPort(cmd),
		ServerOption: nil,
		ProxyPort:    proxyPort(cmd),
		ProxyOption: []runtime.ServeMuxOption{
			grpcutil.RemoveOutgoingHeader(nil),
		},
		ProxyOpenAPIHandler: gen.MustGetOpenAPIHandler(),
	})

	if err != nil {
		return err
	}
	defer cleanup()

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	wire.Proxy.Register(func(server *grpc.Server, mux *runtime.ServeMux, endpoint string) {
		authenticationv1.RegisterServiceServer(server, wire.Authentication)
		authenticationv1.RegisterServiceHandlerFromEndpoint(
			cmd.Context(), mux, endpoint, dialOpts)
	}).Listen(cmd.Context())
	return nil
}

func grpcPort(cmd *cobra.Command) grpcutil.ServerPort {
	return grpcutil.ServerPort(cobrautil.MustGetIntFlag(cmd, "grpc-port"))
}

func proxyPort(cmd *cobra.Command) grpcutil.ProxyPort {
	return grpcutil.ProxyPort(cobrautil.MustGetIntFlag(cmd, "proxy-port"))
}

func x() runtime.ServeMuxOption {
	return runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, message proto.Message) error {
		smd, _ := runtime.ServerMetadataFromContext(ctx)

		if y := smd.HeaderMD.Get("testing"); len(y) >= 1 {
			http.SetCookie(w, &http.Cookie{
				Name:     "testing",
				Value:    y[0],
				Path:     "/",
				Expires:  time.Now().Add(time.Minute * 10),
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			})
		}
		return nil
	})
}
