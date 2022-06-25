package graphql

import (
	"github.com/aklinker1/miasma/internal"
	"github.com/ssoroka/slice"
)

func toBoundVolumes(inputs []internal.BoundVolumeInput) []internal.BoundVolume {
	return slice.Map[internal.BoundVolumeInput, internal.BoundVolume](inputs, toBoundVolume)
}

func toBoundVolume(input internal.BoundVolumeInput) internal.BoundVolume {
	return internal.BoundVolume{
		Target: input.Target,
		Source: input.Source,
	}
}
