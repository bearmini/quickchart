package quickchart

func String(s string) *string {
	return &s
}

func Number(f float64) *float64 {
	return &f
}

func Bool(b bool) *bool {
	return &b
}
