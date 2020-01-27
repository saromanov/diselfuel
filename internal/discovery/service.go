package discovery

import "github.com/saromanov/diselfuel/internal/models"

// Discovery defines interface for manging nodes
type Discovery interface {
	ListNodes(FilterNodes) ([]*models.Host, error)
}

// FilterNodes provides filtering of the nodes by params
type FilterNodes struct {
	Tag string
}
