package component

import (
	"github.com/clavinjune/gokit/grpcutil"
	authenticationv1 "github.com/clavinjune/magellan/api/proto/magellan/authentication/v1"
)

func NewComponent(
	server *grpcutil.Server,
	proxy *grpcutil.Proxy,
	authentication authenticationv1.ServiceServer,
) *Component {
	return &Component{
		Server:         server,
		Proxy:          proxy,
		Authentication: authentication,
	}
}
