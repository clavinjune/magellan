package authentication

import (
	"sync"

	authenticationv1 "github.com/clavinjune/magellan/api/proto/magellan/authentication/v1"
	"github.com/google/wire"
)

var (
	hdl     *handler
	hdlOnce sync.Once

	ProviderSet = wire.NewSet(
		New,
		wire.Bind(new(authenticationv1.ServiceServer), new(*handler)),
	)
)

func New() *handler {
	hdlOnce.Do(func() {
		hdl = &handler{}
	})

	return hdl
}
