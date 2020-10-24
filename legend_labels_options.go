package quickchart

type LegendLabelsOptions struct {
	BoxWidth      *float64 `json:"boxWidth,omitempty"`
	FontSize      *float64 `json:"fontSize,omitempty"`
	FontColor     *string  `json:"fontColor,omitempty"`
	FontFamily    *string  `json:"fontFamily,omitempty"`
	Padding       *float64 `json:"padding,omitempty"`
	UsePointStyle *bool    `json:"usePointStyle,omitempty"`
}
