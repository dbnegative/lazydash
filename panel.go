//MIT License

//Copyright (c) 2019 Jason Witting

//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:

//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.

//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package main

import "strings"

//Panel contains all Grafana Panel attributes
type Panel struct {
	GridPos       PanelGridPos  `json:"gridPos"`
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	ID            int           `json:"id"`
	Mode          string        `json:"mode"`
	Content       string        `json:"content,omitempty"`
	Targets       []PanelTarget `json:"targets,omitempty"`
	Description   string        `json:"description,omitempty"`
	Legend        PanelLegend   `json:"legend,omitempty"`
	Bars          bool          `json:"bars,omitempty"`
	DashLength    int           `json:"dashLength"`
	Dashes        bool          `json:"dashes,omitempty"`
	Fill          int           `json:"fill,omitempty"`
	FillGradient  int           `json:"fillGradient,omitempty"`
	Lines         bool          `json:"lines,omitempty"`
	LinesWidth    int           `json:"linewidth,omitempty"`
	NullPointMode string        `json:"nullPointMode,omitempty"`
	Percentage    bool          `json:"percentage,omitempty"`
	PointRadius   int           `json:"pointradius,omitempty"`
	Points        bool          `json:"points,omitempty"`
	Render        string        `json:"renderer,omitempty"`
	//SeriesOverrides "seriesOverrides": [],
	SpaceLength int          `json:"spaceLength,omitempty"`
	Stack       bool         `json:"stack,omitempty"`
	SteppedLine bool         `json:"steppedLine,omitempty"`
	YAxes       []PanelYAxes `json:"yaxes,omitempty"`
	YAxis       PanelYAxis   `json:"yaxis,omitempty"`
	XAxis       PanelXAxis   `json:"xaxis,omitempty"`
	ToolTip     PanelToolTip `json:"tooltip,omitempty"`
	Options     PanelOptions `json:"options,omitempty"`
}

//PanelFieldOptionsThershold describes a panel option thershold
type PanelFieldOptionsThershold struct {
	Value int    `json:"value,omitempty"`
	Color string `json:"color,omitempty"`
}

//PanelFieldOptionsDefaults describes a panel option defaults
type PanelFieldOptionsDefaults struct {
	Thresholds []PanelFieldOptionsThershold `json:"thresholds,omitempty"`
	//Mappings[]
	Unit string `json:"unit,omitempty"`
}

//PanelFieldOptions describes panel field options
type PanelFieldOptions struct {
	Values   bool                      `json:"values,omitempty"`
	Calcs    []string                  `json:"calcs,omitempty"`
	Defaults PanelFieldOptionsDefaults `json:"defaults,omitempty"`
	//Override": {}
}

//PanelOptions describes panel options
type PanelOptions struct {
	ShowThresholdMarkers bool              `json:"showThresholdMarkers,omitempty"`
	ShowThresholdLabels  bool              `json:"showThresholdLabels,omitempty"`
	FieldOptions         PanelFieldOptions `json:"fieldOptions,omitempty"`
	Orientation          string            `json:"orientation,omitempty"`
}

//PanelLegend describes all Legend options
type PanelLegend struct {
	Avg          bool `json:"avg,omitempty"`
	Current      bool `json:"current,omitempty"`
	Max          bool `json:"max,omitempty"`
	Min          bool `json:"min,omitempty"`
	Show         bool `json:"show,omitempty"`
	Total        bool `json:"total,omitempty"`
	Values       bool `json:"values,omitempty"`
	HideEmpty    bool `json:"hideEmpty,omitempty"`
	HideZero     bool `json:"hideZero,omitempty"`
	RightSide    bool `json:"rightSide,omitempty"`
	SideWidth    int  `json:"sideWidth,omitempty"`
	AlignAsTable bool `json:"alignAsTable,omitempty"`
}

//PanelToolTip describes Panel tooltip options
type PanelToolTip struct {
	Shared    bool   `json:"shared,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	ValueType string `json:"value_type,omitempty"`
}

//PanelXAxis describes Panel xaxis options
type PanelXAxis struct {
	Buckets int    `json:"buckets,omitempty"`
	Mode    string `json:"mode,omitempty"`
	Name    string `json:"name,omitempty"`
	Show    bool   `json:"show,omitempty"`
}

//PanelYAxes describes Yaxes options
type PanelYAxes struct {
	Decimals int    `json:"decimals,omitempty"`
	Format   string `json:"format,omitempty"`
	Label    string `json:"label,omitempty"`
	LogBase  int    `json:"logBase,omitempty"`
	Max      int    `json:"max,omitempty"`
	Min      int    `json:"min,omitempty"`
	Show     bool   `json:"show,omitempty"`
}

//PanelYAxis describes Panel yaxis options
type PanelYAxis struct {
	Align      bool `json:"align,omitempty"`
	AlignLevel int  `json:"alignLevel,omitempty"`
}

//PanelGridPos holds a Panels grid posistion
type PanelGridPos struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	H int `json:"h,omitempty"`
	W int `json:"w,omitempty"`
}

//PanelTarget contains a Panel metrics experssion and order
type PanelTarget struct {
	Expr         string `json:"expr,ommitempty"`        // "sum(rate(process_cpu_seconds_total [1m]))",
	RefID        string `json:"refId,omitempty"`        //"A"
	LegendFormat string `json:"legendFormat,omitempty"` //: "{{state}}"
}

//MetricsTemplate describes an metrics expression template
type MetricsTemplate struct {
	template  string
	delimiter string
	metric    string
}

//NewPanel returns a new initialised Panel
func NewPanel(title string) *Panel {
	return &Panel{
		Title:       title,
		Type:        "graph",
		Description: "Basic Panel",
		Targets: []PanelTarget{
			{
				Expr:         "",
				RefID:        "A",
				LegendFormat: "",
			},
		}, //Targets
		GridPos: PanelGridPos{
			X: 0,
			Y: 0,
			H: 0,
			W: 0,
		}, //GridPos
		YAxes: []PanelYAxes{
			{
				Format:  "short",
				LogBase: 1,
				Show:    true,
			},
			{
				Show: false,
			},
		},
		XAxis: PanelXAxis{
			Mode: "time",
			Name: "",
			Show: true,
		},
		Options: PanelOptions{
			FieldOptions: PanelFieldOptions{
				Defaults: PanelFieldOptionsDefaults{
					Unit: "short",
				},
			},
		},
	}
}

//SetGridPos sets the panel grid position
func (p *Panel) SetGridPos(x, y, h, w int) {
	p.GridPos.X = x
	p.GridPos.Y = y
	p.GridPos.H = h
	p.GridPos.W = w
}

//SetType sets the panel type
func (p *Panel) SetType(ptype string) {
	p.Type = ptype
}

//SetUnit sets the panel y axis unit type
func (p *Panel) SetUnit(unit string) {
	//if type is graph
	p.YAxes[0].Format = unit
	//if type is guage
	p.Options.FieldOptions.Defaults.Unit = unit
}

//SetLegendFormat sets the panel legend format
func (p *Panel) SetLegendFormat(format string) {
	for i := range p.Targets {
		p.Targets[i].LegendFormat = format
	}
}

//SetDescription sets the panel description
func (p *Panel) SetDescription(description string) {
	p.Description = description
}

//SetMetricExpr sets the panels metrics expression
func (p *Panel) SetMetricExpr(expr string) {
	for i := range p.Targets {
		p.Targets[i].Expr = expr
	}
}

//MetricTemplate returns a valid panel metric expression
func (mt *MetricsTemplate) MetricTemplate() string {
	return strings.Replace(mt.template, mt.delimiter, mt.metric, -1)
}

//SetMetric sets a metric
func (mt *MetricsTemplate) SetMetric(metric string) {
	mt.metric = metric
}

//CreateLegendFormat returns a foramtted LegendFormat string based on metric labels
func CreateLegendFormat(labels []string, fallback string) string {

	if len(labels) < 1 {
		if fallback != "" {
			return fallback + ":[{{" + fallback + "}}]"
		}
		return "Job:[{{job}}]"
	}

	formattedLabels := ""

	for _, v := range labels {
		formattedLabels = formattedLabels + v + ":" + "[{{" + v + "}}] "
	}

	return formattedLabels
}
