package quickchart

type Dataset struct {
	Label           *string  `json:"label,omitempty"`
	Data            []Data   `json:"data,omitempty"`
	BackgroundColor []string `json:"backgroundColor,omitempty"`
	BorderColor     []string `json:"borderColor,omitempty"`
	BorderWidth     *int     `json:"borderWidth,omitempty"`
}
