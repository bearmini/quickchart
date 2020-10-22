package quickchart

type ChartData struct {
	Labels   []Label   `json:"labels,omitempty"`
	Datasets []Dataset `json:"datasets,omitempty"`
}
