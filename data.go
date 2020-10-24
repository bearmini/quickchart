package quickchart

type Data interface{}

type DataNumber float64

type DataPoint struct {
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
}

type DataTime struct {
	T *float64 `json:"t,omitempty"`
	Y *float64 `json:"y,omitempty"`
}

type DataBubble struct {
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	R *float64 `json:"r,omitempty"`
}
