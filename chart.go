package quickchart

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Chart struct {
	apiEndpoint      string
	width            *float64
	height           *float64
	devicePixelRatio *float64
	backgroundColor  *string
	format           *string
	encoding         *string
	version          *string
	Type             ChartType     `json:"type"`
	Data             ChartData     `json:"data"`
	Options          *ChartOptions `json:"options,omitempty"`
}

const (
	defaultAPIEndpoint = "https://quickchart.io/"
)

func NewChart(t ChartType, d ChartData, o *ChartOptions) *Chart {
	return &Chart{
		apiEndpoint: defaultAPIEndpoint,
		Type:        t,
		Data:        d,
		Options:     o,
	}
}

func (c *Chart) APIEndpoint(e string) *Chart {
	c.apiEndpoint = e
	return c
}

func (c *Chart) Width(w float64) *Chart {
	c.width = &w
	return c
}

func (c *Chart) Height(h float64) *Chart {
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

	// returns plain json. URL Encoding will be done when generating URL
	return string(b), nil
}

type PostRequest struct {
	BackgroundColor string  `json:"backgroundColor"`
	Width           float64 `json:"width"`
	Height          float64 `json:"height"`
	Format          string  `json:"format"`
	Chart           *Chart  `json:"chart"`
}

func (c *Chart) Download() ([]byte, error) {
	req := PostRequest{
		BackgroundColor: *c.backgroundColor,
		Width:           *c.width,
		Height:          *c.height,
		Format:          *c.format,
		Chart:           c,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/chart", c.apiEndpoint), "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("failed to download chart image: %s", string(respBody))
	}

	return respBody, nil
}

func (c *Chart) GetURL() (string, error) {
	q := url.Values{}
	if c.width != nil {
		q.Add("w", fmt.Sprintf("%f", *c.width))
	}
	if c.height != nil {
		q.Add("h", fmt.Sprintf("%f", *c.height))
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
	if c.version != nil {
		q.Add("version", *c.version)
	}

	encodedChart, err := c.Encode()
	if err != nil {
		return "", err
	}
	q.Add("c", encodedChart)

	return fmt.Sprintf("%s/chart?%s", c.apiEndpoint, q.Encode()), nil
}

type GenerateShortURLRequest struct {
	Chart            *Chart   `json:"chart,omitempty"`
	Width            *float64 `json:"width,omitempty"`
	Height           *float64 `json:"height,omitempty"`
	DevicePixelRatio *float64 `json:"devicePixelRatio,omitempty"`
	BackgroundColor  *string  `json:"backgroundColor,omitempty"`
	Format           *string  `json:"format,omitempty"`
	Encoding         *string  `json:"encoding,omitempty"`
	Version          *string  `json:"version,omitempty"`
}

type GenerateShortURLResponse struct {
	Success bool   `json:"success"`
	URL     string `json:"url"`
}

func GenerateShortURL(chart *Chart) (string, error) {
	req := GenerateShortURLRequest{
		Chart:            chart,
		Width:            chart.width,
		Height:           chart.height,
		DevicePixelRatio: chart.devicePixelRatio,
		BackgroundColor:  chart.backgroundColor,
		Format:           chart.format,
		Encoding:         chart.encoding,
		Version:          chart.version,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	respObj, err := http.Post(fmt.Sprintf("%s/chart/create", chart.apiEndpoint), "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	defer func() {
		respObj.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(respObj.Body)
	if err != nil {
		return "", err
	}

	var resp GenerateShortURLResponse

	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return "", err
	}

	if !resp.Success {
		return "", errors.New(string(respBody))
	}
	return resp.URL, nil
}
