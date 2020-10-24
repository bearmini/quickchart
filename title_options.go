package quickchart

type TitleOptions struct {
	Display    *bool    `json:"display,omitempty"`
	FontSize   *float64 `json:"fontSize,omitempty"`
	FontFamily *string  `json:"fontFamily,omitempty"`
	FontColor  Color    `json:"fontColor,omitempty"`
	FontStyle  *string  `json:"fontStyle,omitempty"`
	LineHeight *float64 `json:"lineHeight,omitempty"`
	Padding    *float64 `json:"padding,omitempty"`
	Position   *string  `json:"position,omitempty"`
	Text       *string  `json:"text,omitempty"`
}
