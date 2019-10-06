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

//Generate a dashboard based off an ingested prometheus metrics
func Generate(metrics MetricMap, gauges bool) *Dashboard {
	dashboard := NewDashboard(*title)
	count, lastx, lasty := 0, 0, 0
	ptype := "graph"
	if gauges {
		ptype = "gauge"
	}

	for _, v := range metrics.List() {
		labels := ""

		if len(metrics.Get(v).Labels()) > 0 {
			for _, v := range metrics.Get(v).Labels() {
				labels = labels + v + ":" + "[{{" + v + "}}] "
			}
		} else {
			labels = "Job:[{{job}}]"
		}

		switch metrics.Get(v).Type() {
		case "counter":

			//add 2 panels to a row
			if (count % 2) < 1 {
				lastx = 0
				dashboard.AddPanel(*NewPanel(metrics.Get(v).Help(),
					"graph",
					metrics.Get(v).Help(),
					metrics.Get(v).Name()+metrics.Get(v).Suffix(),
					metrics.Get(v).Unit(),
					lastx,
					lasty,
					"sum(rate(:METRIC: [1m]))",
					":METRIC:",
					labels))
				lasty = lasty + 9
				count++
			} else {
				lastx = lastx + 12
				dashboard.AddPanel(*NewPanel(metrics.Get(v).Help(),
					"graph", metrics.Get(v).Help(),
					metrics.Get(v).Name()+metrics.Get(v).Suffix(),
					metrics.Get(v).Unit(),
					lastx,
					lasty,
					"sum(rate(:METRIC: [1m]))",
					":METRIC:",
					labels))
				count++
			}

		case "gauge":
			//add 2 panels to a row
			if (count % 2) < 1 {
				lastx = 0
				dashboard.AddPanel(*NewPanel(metrics.Get(v).Help(),
					ptype,
					metrics.Get(v).Help(),
					metrics.Get(v).Name()+metrics.Get(v).Suffix(),
					metrics.Get(v).Unit(),
					lastx,
					lasty,
					":METRIC:",
					":METRIC:",
					labels))
				lasty = lasty + 9
				count++
			} else {
				lastx = lastx + 12
				dashboard.AddPanel(*NewPanel(metrics.Get(v).Help(),
					ptype, metrics.Get(v).Help(),
					metrics.Get(v).Name()+metrics.Get(v).Suffix(),
					metrics.Get(v).Unit(),
					lastx,
					lasty,
					":METRIC:",
					":METRIC:",
					labels))
				count++
			}

		}
	}
	return dashboard
}
