package quickchart

type TimeOptions struct {
	DisplayFormats *DisplayFormatsOption `json:"displayFormats,omitempty"`
	ISOWeekday     *bool                 `json:"isoWeekday,omitempty"`
	Parser         *string               `json:"parser,omitempty"`
	Round          *string               `json:"round,omitempty"`
	TooltipFormat  *string               `json:"tooltipFormat,omitempty"`
	Unit           *string               `json:"unit,omitempty"`
	StepSize       *float64              `json:"stepSize,omitempty"`
	MinUnit        *string               `json:"minUnit,omitempty"`
}
