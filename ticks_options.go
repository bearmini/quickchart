package quickchart

type TicksOptions struct {
	AutoSkip        *bool    `json:"autoSkip,omitempty"`
	AutoSkipPadding *float64 `json:"autoSkipPadding,omitempty"`
	BeginAtZero     *bool    `json:"beginAtZero,omitempty"`
	LabelOffset     *float64 `json:"labelOffset,omitempty"`
	Labels          []string `json:"labels,omitempty"`
	Max             *float64 `json:"max,omitempty"`
	MaxRotation     *float64 `json:"maxRotation,omitempty"`
	MaxTicksLimit   *float64 `json:"maxTicksLimit,omitempty"`
	Min             *float64 `json:"min,omitempty"`
	MinRotation     *float64 `json:"minRotation,omitempty"`
	Mirror          *bool    `json:"mirror,omitempty"`
	Padding         *float64 `json:"padding,omitempty"`
	Precision       *float64 `json:"precision,omitempty"`
	SampleSize      *float64 `json:"sampleSize,omitempty"`
	StepSize        *float64 `json:"stepSize,omitempty"`
	SuggestedMax    *float64 `json:"suggestedMax,omitempty"`
	SuggestedMin    *float64 `json:"suggestedMin,omitempty"`
}
