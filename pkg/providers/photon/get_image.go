package photon

import (
	"github.com/solo-io/unik/pkg/providers/common"
	"github.com/solo-io/unik/pkg/types"
)

func (p *PhotonProvider) GetImage(nameOrIdPrefix string) (*types.Image, error) {
	return common.GetImage(p, nameOrIdPrefix)
}
