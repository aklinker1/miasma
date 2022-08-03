package docker

import (
	"context"
	"sort"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
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
func (s *runtimeTaskRepo) GetAll(ctx context.Context, filter server.RuntimeTasksFilter) ([]internal.AppTask, error) {
	filterArgs := []filters.KeyValuePair{}
	if filter.NodeID != nil {
		filterArgs = append(filterArgs, filters.KeyValuePair{Key: "node", Value: *filter.NodeID})
	}
	if filter.ServiceID != nil {
		filterArgs = append(filterArgs, filters.KeyValuePair{Key: "service", Value: *filter.ServiceID})
	}
	if filter.State != nil {
		filterArgs = append(filterArgs, filters.KeyValuePair{Key: "desired-state", Value: string(*filter.State)})
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

	finalTasks := []internal.AppTask{}
	for _, task := range tasks {
		service, _, err := s.client.ServiceInspectWithRaw(ctx, task.ServiceID, types.ServiceInspectOptions{})
		if err != nil {
			return nil, err
		}
		finalTasks = append(finalTasks, internal.AppTask{
			AppID:        service.Spec.Labels[miasmaIdLabel],
			Name:         service.Spec.Name,
			Message:      task.Status.Message,
			State:        string(task.Status.State),
			DesiredState: string(task.DesiredState),
			Timestamp:    task.Status.Timestamp,
			Error:        lo.Ternary(task.Status.Err == "", nil, &task.Status.Err),
			ExitCode: lo.If[*int](task.Status.ContainerStatus == nil, nil).ElseF(func() *int {
				return &task.Status.ContainerStatus.ExitCode
			}),
		})
	}
	return finalTasks, nil
}
