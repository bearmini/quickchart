package quickchart

type Font struct {
	Color      *string  `json:"color,omitempty"`
	Family     *string  `json:"family,omitempty"`
	Size       *float64 `json:"size,omitempty"`
	Style      *string  `json:"style,omitempty"`
	LineHeight *float64 `json:"lineHeight,omitempty"`
}
