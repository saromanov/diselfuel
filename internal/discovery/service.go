package discovery

// Discovery defines interface for manging nodes
type Discovery interface {
	ListNodes() ([]string, error)
}
