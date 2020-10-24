package quickchart

type LegendOptions struct {
	Align         *string              `json:"align,omitempty"`
	Display       *bool                `json:"display,omitempty"`
	FullWidth     *bool                `json:"fullWidth,omitempty"`
	Labels        *LegendLabelsOptions `json:"labels,omitempty"`
	Position      *string              `json:"position,omitempty"`
	Reverse       *bool                `json:"reverse,omitempty"`
	Rtl           *bool                `json:"rtl,omitempty"`
	TextDirection *string              `json:"textDirection,omitempty"`
}
