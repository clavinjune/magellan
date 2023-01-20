//go:build wireinject

package component

import (
	"github.com/clavinjune/magellan/api/component"
	"github.com/google/wire"
)

func Wire(option Option) (*component.Component, func(), error) {
	panic(wire.Build(ProviderSet))
}
