package volume_mapper

import (
	"docker.io/go-docker/api/types/mount"
	"github.com/aklinker1/miasma/package/models"
)

func ToDocker(model *models.RunConfigVolumesItems0) mount.Mount {
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
