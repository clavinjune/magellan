package component

import (
	"github.com/clavinjune/gokit/grpcutil"
	magellanv1 "github.com/clavinjune/magellan/api/proto/magellan/v1"
)

func NewComponent(
	server *grpcutil.Server,
	proxy *grpcutil.Proxy,
	authentication magellanv1.AuthenticationServiceServer,
) *Component {
	return &Component{
		Server:         server,
		Proxy:          proxy,
		Authentication: authentication,
	}
}
