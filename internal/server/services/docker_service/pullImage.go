package docker_service

import (
	"context"
	"io/ioutil"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func PullImage(baseImage string) error {
	log.V("docker_service.PullImage(%v)", baseImage)
	log.V("Pulling %s...", baseImage)
	ctx := context.Background()

	stream, err := docker.ImageAPIClient.ImagePull(dockerAPI, ctx, baseImage, types.ImagePullOptions{})
	if err != nil {
		log.E("Failed to pull image: %v", err)
		return err
	}
	defer stream.Close()

	// Read 1 image at a time (if pulling more than 1)
	// reader := bufio.NewReader(stream)
	// for ok := true; ok; {
	// 	line, _, err := reader.ReadLine()
	// 	log.V("%s", string(line))
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.E("Failed to read line: %v", err)
	// 		}
	// 		ok = false
	// 	}
	// }
	// Do it all
	_, err = ioutil.ReadAll(stream)
	if err != nil {
		log.E("Failed to save pulled image: %v", err)
		return err
	}
	log.V("Done!")
	return nil
}
