package docker

import "strings"

func isSwarmNotInitializedError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "This node is not a swarm manager")
}
