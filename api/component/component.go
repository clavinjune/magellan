package component

import (
	"github.com/clavinjune/gokit/grpcutil"
	magellanv1 "github.com/clavinjune/magellan/api/proto/magellan/v1"
)

type Component struct {
	_              struct{}
	Server         *grpcutil.Server
	Proxy          *grpcutil.Proxy
	Authentication magellanv1.AuthenticationServiceServer
}
