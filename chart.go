package quickchart

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Chart struct {
	Type    ChartType     `json:"type"`
	Data    ChartData     `json:"data"`
	Options *ChartOptions `json:"options,omitempty"`
}

func NewChart(t ChartType, d ChartData, o *ChartOptions) *Chart {
	return &Chart{
		Type:    t,
		Data:    d,
		Options: o,
	}
}

func (c *Chart) Encode() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return url.QueryEscape(string(b)), nil
}

func (c *Chart) GetURL() (string, error) {
	encodedChart, err := c.Encode()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://quickchart.io/chart?c=%s", encodedChart), nil
}
