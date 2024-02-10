package path

import (
	"encoding/json"
	"net/http"
	"time"
)

type History struct {
	Host    string    `json:"host"`
	Reason  string    `json:"reason"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	PeakBPS PeakData  `json:"peak_bps"`
	PeakPPS PeakData  `json:"peak_pps"`
}
type PeakData struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func (c *APIClient) GetHistory() ([]*History, error) {
	raw, err := c.requestHandler("/attack_history", http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}
	type response struct {
		History []*History `json:"attack_history"`
	}
	data := &response{}
	json.Unmarshal(raw, &data)

	return data.History, nil
}
