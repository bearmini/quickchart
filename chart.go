package quickchart

import (
	"encoding/json"
	"net/url"
)

type Chart struct {
	Type    ChartType     `json:"type"`
	Data    ChartData     `json:"data"`
	Options *ChartOptions `json:"options,omitempty"`
}

func (c *Chart) Encode() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return url.QueryEscape(string(b)), nil
}
