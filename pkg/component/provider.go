package component

import (
	"github.com/clavinjune/gokit/grpcutil"
	"github.com/clavinjune/magellan/api/component"
	"github.com/clavinjune/magellan/pkg/authentication"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

var (
	LocalSet = wire.NewSet(
		component.NewComponent,
		wire.FieldsOf(new(Option), "Logger",
			"Port", "ServerOption",
			"ProxyPort", "ProxyOption", "ProxyOpenAPIHandler"),
		authentication.ProviderSet,
	)

	RemoteSet = wire.NewSet(
		grpc.NewServer,
		grpcutil.ProviderSet,
	)
	ProviderSet = wire.NewSet(
		LocalSet,
		RemoteSet,
	)
)
