package quickchart

type Padding interface{}

type PaddingNumber float64

type PaddingObject struct {
	Bottom *float64 `json:"bottom,omitempty"`
	Left   *float64 `json:"left,omitempty"`
	Right  *float64 `json:"right,omitempty"`
	Top    *float64 `json:"top,omitempty"`
}
