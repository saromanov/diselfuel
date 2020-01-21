package app

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/saromanov/diselfuel/internal/models"
)

// filterNodes provides filtering of nodes by the query
// for example:
// node=test1
// node=test*
func filterNodes(query string, hosts []*models.Host) ([]*models.Host, error) {
	if query == "*" {
		return hosts, nil
	}
	query = prepareQuery(query)
	if strings.HasPrefix(query, "node") {
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
	if strings.HasPrefix(query, "address") {
		rawResult := strings.Split(query, "=")
		if len(rawResult) < 2 {
			return nil, fmt.Errorf("invalid expression")
		}
		result := []*models.Host{}
		re := regexp.MustCompile(rawResult[1])
		for _, h := range hosts {
			if s := re.FindString(h.Address); s != "" {
				result = append(result, h)
			}
		}
	}

	if strings.HasPrefix(query, "user") {
		rawResult := strings.Split(query, "=")
		if len(rawResult) < 2 {
			return nil, fmt.Errorf("invalid expression")
		}
		result := []*models.Host{}
		re := regexp.MustCompile(rawResult[1])
		for _, h := range hosts {
			if s := re.FindString(h.User); s != "" {
				result = append(result, h)
			}
		}
	}

	if strings.HasPrefix(query, "tags") {
		rawResult := strings.Split(query, "=")
		if len(rawResult) < 2 {
			return nil, fmt.Errorf("invalid expression")
		}
		result := []*models.Host{}
		re := regexp.MustCompile(rawResult[1])
		for _, h := range hosts {
			for _, t := range h.Tags {
				if s := re.FindString(t); s != "" {
					result = append(result, h)
					break
				}
			}
		}
	}
	return nil, nil
}

func prepareQuery(s string) string {
	if strings.Contains(s, "=*") {
		return strings.Replace(s, "=*", "=.", 1)
	}
	return s
}
