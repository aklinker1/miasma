package docker

import (
	"context"
	"sort"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/samber/lo"
)

type runtimeTaskRepo struct {
	client client.APIClient
	logger server.Logger
}

func NewRuntimeTaskRepo(logger server.Logger, client *client.Client) server.RuntimeTaskRepo {
	return &runtimeTaskRepo{
		client: client,
		logger: logger,
	}
}

// GetAll implements server.RuntimeTaskRepo
func (s *runtimeTaskRepo) GetAll(ctx context.Context, filter server.RuntimeTasksFilter) ([]internal.RunningContainer, error) {
	filterArgs := []filters.KeyValuePair{}
	if filter.NodeID != nil {
		filterArgs = append(filterArgs, filters.KeyValuePair{Key: "node", Value: *filter.NodeID})
	}
	tasks, err := s.client.TaskList(ctx, types.TaskListOptions{
		Filters: filters.NewArgs(filterArgs...),
	})
	if err != nil {
		return nil, err
	}

	compareTask := func(i, j int) bool {
		return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
	}
	sort.SliceStable(tasks, compareTask)

	serviceIDs := lo.Map(tasks, func(task swarm.Task, _ int) string {
		return task.ServiceID
	})
	serviceIDSet := lo.Uniq(serviceIDs)

	finalTasks := []internal.RunningContainer{}
	for _, serviceID := range serviceIDSet {
		service, _, err := s.client.ServiceInspectWithRaw(ctx, serviceID, types.ServiceInspectOptions{})
		if err != nil {
			return nil, err
		}
		finalTasks = append(finalTasks, internal.RunningContainer{
			AppID: service.Spec.Labels[miasmaIdLabel],
			Name:  service.Spec.Name,
		})
	}
	return finalTasks, nil
}
