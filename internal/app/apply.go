package app

import (
	"fmt"

	"github.com/saromanov/diselfuel/internal/models"
)

// Apply provides applying of components to the servers
func (a *App) Apply(conf *models.Execution) ([]*models.Exec, error) {
	if conf == nil {
		return nil, fmt.Errorf("config is not defined")
	}
	if len(conf.Tasks) == 0 {
		return nil, fmt.Errorf("tasks is not found")
	}

	hosts, err := a.getHosts(conf.Tasks)
	if err != nil {
		return nil, fmt.Errorf("unable to get hosts: %v", err)
	}
	execResult, err := a.exec(hosts, "test", "test")
	if err != nil {
		return nil, fmt.Errorf("unable to execute commands: %v", err)
	}
	return execResult, nil
}

func (a *App) getHosts(tasks []models.Task) ([]*models.Host, error) {
	nodes := []*models.Host{}
	for _, t := range tasks {
		if t.Tag != "" {
			resp, err := a.serv.ListNodes(models.FilterNodes{
				Tag: t.Tag,
			})
			if err != nil {
				return nil, fmt.Errorf("unable to get list of nodes: %v", err)
			}

			if len(resp) > 0 {
				for _, r := range resp {
					nodes = append(nodes, &models.Host{
						Address: r.Address,
						Name:    r.Name,
					})
				}
			}
		}
	}
	return nodes, nil
}
