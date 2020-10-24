package quickchart

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
)

type Chart struct {
	width            *uint
	height           *uint
	devicePixelRatio *float64
	backgroundColor  *string
	format           *string
	encoding         *string
	version          *string
	Type             ChartType     `json:"type"`
	Data             ChartData     `json:"data"`
	Options          *ChartOptions `json:"options,omitempty"`
}

func NewChart(t ChartType, d ChartData, o *ChartOptions) *Chart {
	return &Chart{
		Type:    t,
		Data:    d,
		Options: o,
	}
}

func (c *Chart) Width(w uint) *Chart {
	c.width = &w
	return c
}

func (c *Chart) Height(h uint) *Chart {
	c.height = &h
	return c
}

func (c *Chart) DevicePixelRatio(dpr float64) *Chart {
	c.devicePixelRatio = &dpr
	return c
}

func (c *Chart) BackgroundColor(bc string) *Chart {
	c.backgroundColor = &bc
	return c
}

func (c *Chart) Format(f string) *Chart {
	c.format = &f
	return c
}

func (c *Chart) Encoding(e string) *Chart {
	c.encoding = &e
	return c
}

func (c *Chart) Version(v string) *Chart {
	c.version = &v
	return c
}

func (c *Chart) Encode() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	if c.encoding != nil && *c.encoding == "base64" {
		return base64.StdEncoding.EncodeToString(b), nil
	}
	return url.QueryEscape(string(b)), nil
}

func (c *Chart) GetURL() (string, error) {
	encodedChart, err := c.Encode()
	if err != nil {
		return "", err
	}

	q := url.Values{}
	if c.width != nil {
		q.Add("w", fmt.Sprintf("%d", *c.width))
	}
	if c.height != nil {
		q.Add("h", fmt.Sprintf("%d", *c.height))
	}
	if c.devicePixelRatio != nil {
		q.Add("devicePixelRatio", fmt.Sprintf("%f", *c.devicePixelRatio))
	}
	if c.backgroundColor != nil {
		q.Add("bkg", *c.backgroundColor)
	}
	if c.format != nil {
		q.Add("f", *c.format)
	}
	if c.encoding != nil {
		q.Add("encoding", *c.encoding)
	}

	return fmt.Sprintf("https://quickchart.io/chart?c=%s", encodedChart), nil
}
