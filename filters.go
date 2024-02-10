package path

import (
	"encoding/json"
	"net/http"
)

type Filter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Settings map[string]interface{}
}

func (c *APIClient) GetFilters() (map[string]*Filter, error) {
	raw, err := c.requestHandler("/filters", http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}
	type response struct {
		Filters []*Filter `json:"filters"`
	}
	data := &response{}
	json.Unmarshal(raw, &data)

	filters := map[string]*Filter{}
	for _, f := range data.Filters {
		filters[f.ID] = f
	}
	return filters, nil
}
