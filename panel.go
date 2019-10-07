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

//Panel contains all Grafana panel attributes
type panel struct {
	GridPos       panelGridPos  `json:"gridPos"`
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	ID            int           `json:"id"`
	Mode          string        `json:"mode"`
	Content       string        `json:"content,omitempty"`
	Targets       []panelTarget `json:"targets,omitempty"`
	Description   string        `json:"description,omitempty"`
	Legend        panelLegend   `json:"legend,omitempty"`
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
	YAxes       []panelYAxes `json:"yaxes,omitempty"`
	YAxis       panelYAxis   `json:"yaxis,omitempty"`
	XAxis       panelXAxis   `json:"xaxis,omitempty"`
	ToolTip     panelToolTip `json:"tooltip,omitempty"`
	Options     panelOptions `json:"options,omitempty"`
}

type panelFieldOptionsThershold struct {
	Value int    `json:"value,omitempty"`
	Color string `json:"color,omitempty"`
}

type panelFieldOptionsDefaults struct {
	Thresholds []panelFieldOptionsThershold `json:"thresholds,omitempty"`
	//Mappings[]
	Unit string `json:"unit,omitempty"`
}

type panelFieldOptions struct {
	Values   bool                      `json:"values,omitempty"`
	Calcs    []string                  `json:"calcs,omitempty"`
	Defaults panelFieldOptionsDefaults `json:"defaults,omitempty"`
	//Override": {}
}

type panelOptions struct {
	ShowThresholdMarkers bool              `json:"showThresholdMarkers,omitempty"`
	ShowThresholdLabels  bool              `json:"showThresholdLabels,omitempty"`
	FieldOptions         panelFieldOptions `json:"fieldOptions,omitempty"`
	Orientation          string            `json:"orientation,omitempty"`
}

//PanelLegend contains all Legend options
type panelLegend struct {
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

//PanelToolTip contains panel tooltip options
type panelToolTip struct {
	Shared    bool   `json:"shared,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	ValueType string `json:"value_type,omitempty"`
}

//PanelXAxis contains panel xaxis options
type panelXAxis struct {
	Buckets int    `json:"buckets,omitempty"`
	Mode    string `json:"mode,omitempty"`
	Name    string `json:"name,omitempty"`
	Show    bool   `json:"show,omitempty"`
}

//PanelYAxes contains Yaxes options
type panelYAxes struct {
	Decimals int    `json:"decimals,omitempty"`
	Format   string `json:"format,omitempty"`
	Label    string `json:"label,omitempty"`
	LogBase  int    `json:"logBase,omitempty"`
	Max      int    `json:"max,omitempty"`
	Min      int    `json:"min,omitempty"`
	Show     bool   `json:"show,omitempty"`
}

//PanelYAxis contains panel yaxis options
type panelYAxis struct {
	Align      bool `json:"align,omitempty"`
	AlignLevel int  `json:"alignLevel,omitempty"`
}

//PanelGridPos holds a panels grid posistion
type panelGridPos struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	H int `json:"h,omitempty"`
	W int `json:"w,omitempty"`
}

//PanelTarget contains a Panel metrics experssion and order
type panelTarget struct {
	Expr         string `json:"expr,ommitempty"`        // "sum(rate(process_cpu_seconds_total [1m]))",
	RefID        string `json:"refId,omitempty"`        //"A"
	LegendFormat string `json:"legendFormat,omitempty"` //: "{{state}}"
}

type metricsTemplate struct {
	template  string
	delimiter string
	metric    string
}

func NewPanel(title string) *panel {
	return &panel{
		Title:       title,
		Type:        "graph",
		Description: "Basic Panel",
		Targets: []panelTarget{
			{
				Expr:         "",
				RefID:        "A",
				LegendFormat: "",
			},
		}, //Targets
		GridPos: panelGridPos{
			X: 0,
			Y: 0,
			H: 0,
			W: 0,
		}, //GridPos
		YAxes: []panelYAxes{
			{
				Format:  "short",
				LogBase: 1,
				Show:    true,
			},
			{
				Show: false,
			},
		},
		XAxis: panelXAxis{
			Mode: "time",
			Name: "",
			Show: true,
		},
		Options: panelOptions{
			FieldOptions: panelFieldOptions{
				Defaults: panelFieldOptionsDefaults{
					Unit: "short",
				},
			},
		},
	}
}

func (p *panel) SetGridPos(x, y, h, w int) {
	p.GridPos.X = x
	p.GridPos.Y = y
	p.GridPos.H = h
	p.GridPos.W = w
}

func (p *panel) SetType(ptype string) {
	p.Type = ptype
}

func (p *panel) SetUnit(unit string) {
	//if type is graph
	p.YAxes[0].Format = unit
	//if type is guage
	p.Options.FieldOptions.Defaults.Unit = unit
}

func (p *panel) SetLegendFormat(format string) {
	for i := range p.Targets {
		p.Targets[i].LegendFormat = format
	}
}

func (p *panel) SetDescription(description string) {
	p.Description = description
}

func (p *panel) SetMetricExpr(expr string) {
	for i := range p.Targets {
		p.Targets[i].Expr = expr
	}
}

func (mt *metricsTemplate) MetricTemplate() string {
	return strings.Replace(mt.template, mt.delimiter, mt.metric, -1)
}

func (mt *metricsTemplate) SetMetric(metric string) {
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
