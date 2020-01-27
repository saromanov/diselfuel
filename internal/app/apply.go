package app

import (
	"fmt"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/discovery"
)

// Apply provides applying of components to the servers
func (a *App) Apply(conf *config.Execution) error {
	if conf == nil {
		return fmt.Errorf("config is not defined")
	}
	if len(conf.Tasks) == 0 {
		return fmt.Errorf("tasks is not found")
	}

	for _, t := range conf.Tasks {
		if t.Tag != "" {

		}
	}
	return nil
}

func (a *App) getHosts(t config.Task) ([]string, error) {
	nodes := []string{}
	if t.Tag != "" {
		resp, err := a.serv.ListNodes(discovery.FilterNodes{
			Tag: t.Tag,
		})
		if err != nil {
			return nil, fmt.Errorf("unable to get list of nodes: %v", err)
		}

		if len(resp) > 0 {
			for _, r := range resp {
				nodes = append(nodes, r.Address)
			}
		}
	}
	return nodes, nil
}
