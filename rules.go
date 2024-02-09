package path

import (
	"encoding/json"
	"net/http"
)

// firewall rule
type Rule struct {
	ID string `json:"id"`

	// source values
	Source     string `json:"source"`
	SourcePort int    `json:"src_port"`
	SourceASN  int    `json:"source_asn"`

	// destination values
	Destination     string `json:"destination"`
	DestinationPort int    `json:"dst_port"`

	Protocol      string `json:"protocol"`
	RateLimiterID string `json:"rate_limiter_id"`
	Whitelist     bool   `json:"whitelist"`
	Priority      bool   `json:"priority"`
	Comment       string `json:"comment"`
}

type CreateRule struct {
	// source values
	Source     string `json:"source,omitempty"`
	SourcePort int    `json:"src_port,omitempty"`
	SourceASN  int    `json:"source_asn,omitempty"`

	// destination values
	Destination     string `json:"destination"`
	DestinationPort int    `json:"dst_port"`

	Protocol      string `json:"protocol"`
	RateLimiterID string `json:"rate_limiter_id,omitempty"`
	Whitelist     bool   `json:"whitelist"`
	Priority      bool   `json:"priority"`
	Comment       string `json:"comment"`
}

func (c *APIClient) GetRules() (map[string]*Rule, error) {
	raw, err := c.requestHandler("/rules", http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}
	type response struct {
		Rules []*Rule
	}
	data := &response{}
	json.Unmarshal(raw, &data)

	rules := map[string]*Rule{}
	for _, r := range data.Rules {
		rules[r.ID] = r
	}
	return rules, nil
}

func (c *APIClient) GetRule(id string) *Rule {
	return nil
}

func (c *APIClient) CreateRule(rule *CreateRule) (*Rule, error) {
	raw, _ := json.Marshal(rule)
	raw, err := c.requestHandler("/rules", http.MethodPost, nil, raw)
	if err != nil {
		return nil, err
	}
	r := &Rule{}
	json.Unmarshal(raw, &r)
	return r, nil
}
