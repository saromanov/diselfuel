package app

import (
	"fmt"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/discovery"
	"github.com/saromanov/diselfuel/internal/models"
)

// Apply provides applying of components to the servers
func (a *App) Apply(conf *config.Execution) error {
	if conf == nil {
		return fmt.Errorf("config is not defined")
	}
	if len(conf.Tasks) == 0 {
		return fmt.Errorf("tasks is not found")
	}

	hosts, err := a.getHosts(conf.Tasks)
	if err != nil {
		return fmt.Errorf("unable to get hosts: %v", err)
	}
	if _, err := a.exec(hosts, "test", "test"); err != nil {
		return fmt.Errorf("unable to execute commands: %v", err)
	}
	return nil
}

func (a *App) getHosts(tasks []config.Task) ([]*models.Host, error) {
	nodes := []*models.Host{}
	for _, t := range tasks {
		if t.Tag != "" {
			resp, err := a.serv.ListNodes(discovery.FilterNodes{
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
