package quickchart

type ScaleLabelOptions struct {
	Display     *bool      `json:"display,omitempty"`
	FontSize    *float64   `json:"fontSize,omitempty"`
	FontFamily  *string    `json:"fontFamily,omitempty"`
	FontColor   Color      `json:"fontColor,omitempty"`
	FontStyle   *string    `json:"fontStyle,omitempty"`
	LabelString *string    `json:"labelString,omitempty"`
	LineHeight  LineHeight `json:"lineHeight,omitempty"`
	Padding     Padding    `json:"padding,omitempty"`
}
