package quickchart

type ChartOptions struct {
	Elements *ElementsOptions `json:"elements,omitempty"`
	Hover    *HoverOptions    `json:"hover,omitempty"`
	Legend   *LegendOptions   `json:"legend,omitempty"`
	Scales   *ScalesOptions   `json:"scales,omitempty"`
	Title    *TitleOptions    `json:"title,omitempty"`
	Tooltips *TooltipsOptions `json:"tooltips,omitempty"`
}
