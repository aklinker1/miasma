package mappers

import (
	"docker.io/go-docker/api/types/mount"
	"github.com/aklinker1/miasma/internal/server/gen/models"
)

type volume struct{}

var Volume = &volume{}

func (mapper *volume) ToDocker(model *models.AppConfigVolumesItems0) mount.Mount {
	volumeType := mount.TypeBind
	// if model.Type != nil {
	// 	volumeType = mount.Type(*model.Type)
	// }
	return mount.Mount{
		Source: model.Source,
		Target: model.Target,
		Type:   volumeType,
	}
}
