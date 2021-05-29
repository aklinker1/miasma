package docker_service

import (
	"context"
	"errors"
	"strings"

	"docker.io/go-docker"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func GetImageDigest(baseImage string) (string, error) {
	log.V("docker_service.GetImageDigest(%v)", baseImage)
	ctx := context.Background()

	info, _, err := docker.ImageAPIClient.ImageInspectWithRaw(dockerAPI, ctx, baseImage)
	if err != nil {
		log.E("Failed to inspect image: %v", err)
		return "", err
	}
	for _, digest := range info.RepoDigests {
		if strings.Contains(digest, "@sha256:") {
			digest := digest[strings.LastIndex(digest, "@")+1:]
			log.V("Digest: %v", digest)
			return digest, nil
		}
	}
	return "", errors.New("Could not find digest with hash instead of tag")
}
