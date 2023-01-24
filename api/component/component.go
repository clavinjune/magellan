package component

import (
	"github.com/clavinjune/gokit/grpcutil"
	authenticationv1 "github.com/clavinjune/magellan/api/proto/magellan/authentication/v1"
)

type Component struct {
	_              struct{}
	Server         *grpcutil.Server
	Proxy          *grpcutil.Proxy
	Authentication authenticationv1.ServiceServer
}
