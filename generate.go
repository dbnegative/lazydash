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

//Generate a Dashboard based off an ingested prometheus metrics
func Generate(metrics MetricMap, gauges bool) *Dashboard {

	Dashboard := NewDashboard(*title)
	count, lastx, lasty := 0, 0, 0

	for _, v := range metrics.List() {

		p := NewPanel(strings.Replace(metrics.Get(v).Name(), "_", " ", -1))
		p.SetDescription(metrics.Get(v).Help())
		p.SetUnit(metrics.Get(v).Unit())
		p.SetLegendFormat(CreateLegendFormat(metrics.Get(v).Labels(), ""))
		p.SetGridPos(lastx, lasty, 7, 12)

		switch metrics.Get(v).Type() {

		case "counter":

			counterTmpl.SetMetric(metrics.Get(v).Name() + metrics.Get(v).Suffix())
			p.SetMetricExpr(counterTmpl.MetricTemplate())
			p.SetType("graph")

			//add 2 Panels to a row
			if (count % 2) < 1 {
				Dashboard.AddPanel(*p)
				lastx = 0
				lasty = lasty + 9
				count++
			} else {
				lastx = lastx + 12
				Dashboard.AddPanel(*p)
				count++
			}

		case "gauge":

			gaugeTmpl.SetMetric(metrics.Get(v).Name() + metrics.Get(v).Suffix())
			p.SetMetricExpr(gaugeTmpl.MetricTemplate())
			if gauges {
				p.SetType("gauges")
			}

			//add 2 Panels to a row
			if (count % 2) < 1 {
				lastx = 0
				Dashboard.AddPanel(*p)
				lasty = lasty + 9
				count++
			} else {
				lastx = lastx + 12
				Dashboard.AddPanel(*p)
				count++
			}

		}
	}
	//fmt.Printf("%+v\n", Dashboard)
	return Dashboard
}
