package quickchart

type ElementsOptions struct {
	Point     *PointElementOptions     `json:"point,omitempty"`
	Rectangle *RectangleElementOptions `json:"rectangle,omitempty"`
}

type PointElementOptions struct {
	PointStyle *string `json:"pointStyle,omitempty"`
}

type RectangleElementOptions struct {
	BorderWidth *float64 `json:"borderWidth,omitempty"`
}
