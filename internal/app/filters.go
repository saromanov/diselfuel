package app

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/saromanov/diselfuel/internal/models"
)

type filterRequest struct {
	query  string
	prefix string
}

// filterNodes provides filtering of nodes by the query
// for example:
// node=test1
// node=test*
func filterNodes(query string, hosts []*models.Host) ([]*models.Host, error) {
	if query == "*" {
		return hosts, nil
	}
	query = prepareQuery(query)
	filteredHosts, err := func(f []filterRequest, hosts []*models.Host) ([]*models.Host, error) {
		for _, m := range f {
			hosts, err := generalFilter(m.query, m.prefix, hosts)
			if err != nil {
				return nil, err
			}
			if len(hosts) > 0 {
				return hosts, nil
			}
		}
		return nil, nil
	}([]filterRequest{filterRequest{
		query:  query,
		prefix: "node",
	},
		filterRequest{
			query:  query,
			prefix: "address",
		},
		filterRequest{
			query:  query,
			prefix: "user",
		},
		filterRequest{
			query:  query,
			prefix: "tags",
		}}, hosts)
	return filteredHosts, err
}

func prepareQuery(s string) string {
	if strings.Contains(s, "=*") {
		return strings.Replace(s, "=*", "=.", 1)
	}
	return s
}

func generalFilter(query, prefix string, hosts []*models.Host) ([]*models.Host, error) {
	if !strings.HasPrefix(query, "node") {
		return nil, nil
	}
	rawResult := strings.Split(query, "=")
	if len(rawResult) < 2 {
		return nil, fmt.Errorf("invalid expression")
	}
	result := []*models.Host{}
	re := regexp.MustCompile(rawResult[1])
	for _, h := range hosts {
		if s := re.FindString(h.Name); s != "" {
			result = append(result, h)
		}
	}

	return result, nil
}
