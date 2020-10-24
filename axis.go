package quickchart

type Axis struct {
	Display    *bool              `json:"display,omitempty"`
	GridLines  *GridLinesOptions  `json:"gridLines,omitempty"`
	ID         *string            `json:"id,omitempty"`
	Max        *float64           `json:"max,omitempty"`
	Min        *float64           `json:"min,omitempty"`
	Offset     *bool              `json:"offset,omitempty"`
	Position   *string            `json:"position,omitempty"`
	ScaleLabel *ScaleLabelOptions `json:"scaleLabel,omitempty"`
	Stacked    *bool              `json:"stacked,omitempty"`
	Ticks      *TicksOptions      `json:"ticks,omitempty"`
	Type       *string            `json:"type,omitempty"`
}
