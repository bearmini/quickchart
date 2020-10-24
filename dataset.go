package quickchart

type Dataset interface{}

type LineChartDataset struct {
	Data                      []Data    `json:"data,omitempty"`
	BackgroundColor           *string   `json:"backgroundColor,omitempty"`
	BorderCapStyle            *string   `json:"borderCapStyle,omitempty"`
	BorderColor               *string   `json:"borderColor,omitempty"`
	BorderDash                []float64 `json:"borderDash,omitempty"`
	BorderDashOffset          *float64  `json:"borderDashOffset,omitempty"`
	BorderJoinStyle           *float64  `json:"borderJoinStyle,omitempty"`
	BorderWidth               *float64  `json:"borderWidth,omitempty"`
	CubicInterpolationMode    *string   `json:"cubicInterpolationMode,omitempty"`
	Clip                      *float64  `json:"clip,omitempty"`
	Fill                      *bool     `json:"fill,omitempty"`
	HoverBackgroundColor      *string   `json:"hoverBackgroundColor,omitempty"`
	HoverBorderCapStyle       *string   `json:"hoverBorderCapStyle,omitempty"`
	HoverBorderColor          *string   `json:"hoverBorderColor,omitempty"`
	HoverBorderDash           []float64 `json:"hoverBorderDash,omitempty"`
	HoverBorderDashOffset     *float64  `json:"hoverBorderDashOffset,omitempty"`
	HoverBorderJoinStyle      *string   `json:"hoverBorderJoinStyle,omitempty"`
	HoverBorderWidth          *float64  `json:"hoverBorderWidth,omitempty"`
	Label                     *string   `json:"label,omitempty"`
	LineTension               *float64  `json:"lineTension,omitempty"`
	Order                     *float64  `json:"order,omitempty"`
	PointBackgroundColor      *string   `json:"pointBackgroundColor,omitempty"`
	PointBorderColor          *string   `json:"pointBorderColor,omitempty"`
	PointBorderWidth          *float64  `json:"pointBorderWidth,omitempty"`
	PointHitRadius            *float64  `json:"pointHitRadius,omitempty"`
	PointHoverBackgroundColor *string   `json:"pointHoverBackgroundColor,omitempty"`
	PointHoverBorderColor     *string   `json:"pointHoverBorderColor,omitempty"`
	PointHoverBorderWidth     *float64  `json:"pointHoverBorderWidth,omitempty"`
	PointHoverRadius          *float64  `json:"pointHoverRadius,omitempty"`
	PointRadius               *float64  `json:"pointRadius,omitempty"`
	PointRotation             *float64  `json:"pointRotation,omitempty"`
	PointStyle                *string   `json:"pointStyle,omitempty"`
	ShowLine                  *bool     `json:"showLine,omitempty"`
	SpanGaps                  *bool     `json:"spanGaps,omitempty"`
	SteppedLine               *bool     `json:"steppedLine,omitempty"`
	XAxisID                   *string   `json:"xAxisID,omitempty"`
	YAxisID                   *string   `json:"yAxisID,omitempty"`
}

type BarChartDataset struct {
	Data                 []Data   `json:"data,omitempty"`
	BackgroundColor      *string  `json:"backgroundColor,omitempty"`
	BarPercentage        *float64 `json:"barPercentage,omitempty"`
	BarThickness         *float64 `json:"barThickness,omitempty"`
	BorderColor          *string  `json:"borderColor,omitempty"`
	BorderSkipped        *string  `json:"borderSkipped,omitempty"`
	BorderWidth          *float64 `json:"borderWidth,omitempty"`
	CategoryPercentage   *float64 `json:"categoryPercentage,omitempty"`
	HoverBackgroundColor *string  `json:"hoverBackgroundColor,omitempty"`
	HoverBorderColor     *string  `json:"hoverBorderColor,omitempty"`
	HoverBorderWidth     *float64 `json:"hoverBorderWidth,omitempty"`
	Label                *string  `json:"label,omitempty"`
	MaxBarThickness      *float64 `json:"maxBarThickness,omitempty"`
	MinBarLength         *float64 `json:"minBarLength,omitempty"`
	Order                *float64 `json:"order,omitempty"`
	XAxisID              *string  `json:"xAxisID,omitempty"`
	YAxisID              *string  `json:"yAxisID,omitempty"`
}

type RadarChartDataset struct {
	Data                  []Data    `json:"data,omitempty"`
	BackgroundColor       *string   `json:"backgroundColor,omitempty"`
	BorderCapStyle        *string   `json:"borderCapStyle,omitempty"`
	BorderColor           *string   `json:"borderColor,omitempty"`
	BorderDash            []float64 `json:"borderDash,omitempty"`
	BorderDashOffset      *float64  `json:"borderDashOffset,omitempty"`
	BorderJoinStyle       *float64  `json:"borderJoinStyle,omitempty"`
	BorderWidth           *float64  `json:"borderWidth,omitempty"`
	HoverBackgroundColor  *string   `json:"hoverBackgroundColor,omitempty"`
	HoverBorderCapStyle   *string   `json:"hoverBorderCapStyle,omitempty"`
	HoverBorderColor      *string   `json:"hoverBorderColor,omitempty"`
	HoverBorderDash       []float64 `json:"hoverBorderDash,omitempty"`
	HoverBorderDashOffset *float64  `json:"hoverBorderDashOffset,omitempty"`
	HoverBorderJoinStyle  *string   `json:"hoverBorderJoinStyle,omitempty"`
	HoverBorderWidth      *float64  `json:"hoverBorderWidth,omitempty"`
	Fill                  *bool     `json:"fill,omitempty"`
	Label                 *string   `json:"label,omitempty"`
	LineTension           *float64  `json:"lineTension,omitempty"`
	Order                 *float64  `json:"order,omitempty"`
	PointBackgroundColor  *string   `json:"pointBackgroundColor,omitempty"`
	PointBorderColor      *string   `json:"pointBorderColor,omitempty"`
	PointBorderWidth      *float64  `json:"pointBorderWidth,omitempty"`
	PointHoverRadius      *float64  `json:"pointHoverRadius,omitempty"`
	PointRadius           *float64  `json:"pointRadius,omitempty"`
	PointRotation         *float64  `json:"pointRotation,omitempty"`
	PointStyle            *string   `json:"pointStyle,omitempty"`
	SpanGaps              *bool     `json:"spanGaps,omitempty"`
}

type DoughnutChartDataset struct {
	Data                 []Data   `json:"data,omitempty"`
	BackgroundColor      *string  `json:"backgroundColor,omitempty"`
	BorderAlign          *string  `json:"borderAlign,omitempty"`
	BorderColor          *string  `json:"borderColor,omitempty"`
	BorderWidth          *float64 `json:"borderWidth,omitempty"`
	HoverBackgroundColor *string  `json:"hoverBackgroundColor,omitempty"`
	HoverBorderColor     *string  `json:"hoverBorderColor,omitempty"`
	HoverBorderWidth     *float64 `json:"hoverBorderWidth,omitempty"`
	Weight               *float64 `json:"weight,omitempty"`
}

type PolarAreaChartDataset struct {
	Data                 []Data   `json:"data,omitempty"`
	BackgroundColor      *string  `json:"backgroundColor,omitempty"`
	BorderAlign          *string  `json:"borderAlign,omitempty"`
	BorderColor          *string  `json:"borderColor,omitempty"`
	BorderWidth          *float64 `json:"borderWidth,omitempty"`
	HoverBackgroundColor *string  `json:"hoverBackgroundColor,omitempty"`
	HoverBorderColor     *string  `json:"hoverBorderColor,omitempty"`
	HoverBorderWidth     *float64 `json:"hoverBorderWidth,omitempty"`
}

type BubbleChartDataset struct {
	Data                 []Data   `json:"data,omitempty"`
	BackgroundColor      *string  `json:"backgroundColor,omitempty"`
	BorderColor          *string  `json:"borderColor,omitempty"`
	BorderWidth          *float64 `json:"borderWidth,omitempty"`
	HoverBackgroundColor *string  `json:"hoverBackgroundColor,omitempty"`
	HoverBorderColor     *string  `json:"hoverBorderColor,omitempty"`
	HoverBorderWidth     *float64 `json:"hoverBorderWidth,omitempty"`
	HoverRadius          *float64 `json:"hoverRadius,omitempty"`
	HitRadius            *float64 `json:"hitRadius,omitempty"`
	Label                *string  `json:"label,omitempty"`
	Order                *float64 `json:"order,omitempty"`
	PointStyle           *string  `json:"pointStyle,omitempty"`
	Radius               *float64 `json:"radius,omitempty"`
	Rotation             *float64 `json:"rotation,omitempty"`
}
