package component

import (
	"github.com/clavinjune/gokit/grpcutil"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"net/http"
)

type Option struct {
	_ struct{}

	Logger       *slog.Logger
	Port         grpcutil.ServerPort
	ServerOption []grpc.ServerOption

	ProxyPort           grpcutil.ProxyPort
	ProxyOption         []runtime.ServeMuxOption
	ProxyOpenAPIHandler http.Handler
}
