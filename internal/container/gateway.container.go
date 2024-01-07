package container

import (
	"m2m/internal/infra"
	"m2m/internal/shared"
)

type ContainerGateway struct {
	acquirersGateway shared.AcquiresGateway
}

func NewContainerGateway() *ContainerGateway {
	return &ContainerGateway{
		acquirersGateway: infra.NewAcquiresGateway(),
	}
}
