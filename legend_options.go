package quickchart

type LegendOptions struct {
	Display  *bool                `json:"display,omitempty"`
	Labels   *LegendLabelsOptions `json:"labels,omitempty"`
	Position *string              `json:"position,omitempty"`
}
