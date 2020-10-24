package quickchart

type ElementsOptions struct {
	Arc       *ArcElementOptions       `json:"arc,omitempty"`
	Line      *LineElementOptions      `json:"line,omitempty"`
	Point     *PointElementOptions     `json:"point,omitempty"`
	Rectangle *RectangleElementOptions `json:"rectangle,omitempty"`
}

type ArcElementOptions struct {
	Angle           *float64 `json:"angle,omitempty"`
	BackgroundColor Color    `json:"backgroundColor,omitempty"`
	BorderAlign     *string  `json:"borderAlign,omitempty"`
	BorderColor     Color    `json:"borderColor,omitempty"`
	BorderWidth     *float64 `json:"borderWidth,omitempty"`
}

type LineElementOptions struct {
	BackgroundColor        Color     `json:"backgroundColor,omitempty"`
	BorderColor            Color     `json:"borderColor,omitempty"`
	BorderCapStyle         *string   `json:"borderCapStyle,omitempty"`
	BorderDash             []float64 `json:"borderDash,omitempty"`
	BorderDashOffset       *float64  `json:"borderDashOffset,omitempty"`
	BorderJoinStyle        *string   `json:"borderJoinStyle,omitempty"`
	BorderWidth            *float64  `json:"borderWidth,omitempty"`
	CapBezierPoints        *bool     `json:"capBezierPoints,omitempty"`
	CubicInterpolationMode *string   `json:"cubicInterpolationMode,omitempty"`
	Fill                   Fill      `json:"fill,omitempty"`
	Stepped                *bool     `json:"stepped,omitempty"`
	Tension                *float64  `json:"tension,omitempty"`
}

type PointElementOptions struct {
	PointStyle *string `json:"pointStyle,omitempty"`
}

type RectangleElementOptions struct {
	BackgroundColor Color    `json:"backgroundColor,omitempty"`
	BorderWidth     *float64 `json:"borderWidth,omitempty"`
	BorderColor     Color    `json:"borderColor,omitempty"`
	BorderSkipped   *string  `json:"borderSkipped,omitempty"`
}
