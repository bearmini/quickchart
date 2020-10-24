package quickchart

type ScaleLabelOptions struct {
	Display     *bool   `json:"display,omitempty"`
	LabelString *string `json:"labelString,omitempty"`
	Font        *Font   `json:"font,omitempty"`
	Padding     Padding `json:"padding,omitempty"`
}
