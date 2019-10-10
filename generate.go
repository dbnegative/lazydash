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

import (
	"strings"
)

var (
	counterTmpl = &MetricsTemplate{
		template:  "sum(rate(:METRIC: [1m]))",
		delimiter: ":METRIC:",
	}
	gaugeTmpl = &MetricsTemplate{
		template:  ":METRIC:",
		delimiter: ":METRIC:",
	}
)

//PanelGridLayout describes a panel grid layout
type PanelGridLayout struct {
	X     int
	Y     int
	Count int
}

//NewPanelGridLayout returns a new panel grid layout
func NewPanelGridLayout() *PanelGridLayout {
	return &PanelGridLayout{X: 0, Y: 0, Count: 0}
}

//UpdateY updates the last y grid posistion
func (p *PanelGridLayout) UpdateY(y int) {
	p.Y = p.Y + y
	p.Count = p.Count + 1
}

//UpdateX updates the last x grid posistion
func (p *PanelGridLayout) UpdateX(x int) {
	if x > 0 {
		p.X = p.X + x
		p.Count = p.Count + 1
	} else {
		p.X = 0
	}

}

//Generate a Dashboard based off an ingested prometheus metrics
func (d *Dashboard) Generate(metrics MetricMap, cfg Config) {

	//count, lastx, lasty := 0, 0, 0
	pgrid := NewPanelGridLayout()

	if len(cfg.CounterExprTmpl) > 0 {
		counterTmpl.SetTemplate(string(cfg.CounterExprTmpl))
		counterTmpl.SetDelimiter(cfg.Delimiter)
	}

	if len(cfg.GaugeExprTmpl) > 0 {
		gaugeTmpl.SetTemplate(string(cfg.GaugeExprTmpl))
		counterTmpl.SetDelimiter(cfg.Delimiter)
	}

	for _, v := range metrics.List() {

		p := NewPanel(strings.Replace(metrics.Get(v).Name(), "_", " ", -1))
		p.SetDescription(metrics.Get(v).Help())
		p.SetUnit(metrics.Get(v).Unit())
		p.SetGridPos(pgrid.X, pgrid.Y, 7, 12)

		if cfg.Table {
			p.Legend = PanelLegend{
				Show:         true,
				Current:      true,
				Values:       true,
				AlignAsTable: true,
			}
		}

		switch metrics.Get(v).Type() {

		case "counter":

			counterTmpl.SetMetric(metrics.Get(v).Name() + metrics.Get(v).Suffix())
			p.SetMetricExpr(counterTmpl.MetricTemplate())
			p.SetLegendFormat(
				CreateLegendFormat(
					metrics.Get(v).Labels(),
					cfg.CounterLegend))
			p.SetType("graph")
			d.AddPanel(*p)

		case "gauge":

			gaugeTmpl.SetMetric(metrics.Get(v).Name() + metrics.Get(v).Suffix())
			p.SetMetricExpr(gaugeTmpl.MetricTemplate())
			p.SetLegendFormat(
				CreateLegendFormat(
					metrics.Get(v).Labels(),
					cfg.GaugeLegend))

			if cfg.Gauges {
				p.SetType("gauge")
			}

			d.AddPanel(*p)
		}

		//add 2 Panels to a row
		if (pgrid.Count % 2) < 1 {
			pgrid.UpdateX(0)
			pgrid.UpdateY(9)
		} else {
			pgrid.UpdateX(12)
		}

	}
}
