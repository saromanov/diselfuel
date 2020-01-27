package discovery

import "github.com/saromanov/diselfuel/internal/models"

// Discovery defines interface for manging nodes
type Discovery interface {
	ListNodes(models.FilterNodes) ([]*models.Host, error)
}
