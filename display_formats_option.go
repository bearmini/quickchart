package quickchart

type DisplayFormatsOption struct {
	Millisecond *string `json:"millisecond,omitempty"`
	Second      *string `json:"second,omitempty"`
	Minute      *string `json:"minute,omitempty"`
	Hour        *string `json:"hour,omitempty"`
	Day         *string `json:"day,omitempty"`
	Week        *string `json:"week,omitempty"`
	Month       *string `json:"month,omitempty"`
	Quarter     *string `json:"quarter,omitempty"`
	Year        *string `json:"year,omitempty"`
}
