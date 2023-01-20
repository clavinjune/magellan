package authentication

import (
	"sync"

	magellanv1 "github.com/clavinjune/magellan/api/proto/magellan/v1"
	"github.com/google/wire"
)

var (
	hdl     *handler
	hdlOnce sync.Once

	ProviderSet = wire.NewSet(
		New,
		wire.Bind(new(magellanv1.AuthenticationServiceServer), new(*handler)),
	)
)

func New() *handler {
	hdlOnce.Do(func() {
		hdl = &handler{}
	})

	return hdl
}
