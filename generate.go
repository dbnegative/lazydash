package main

//Generate a dashboard based off an ingested prometheus metrics
func Generate(metrics MetricMap) *Dashboard {
	dashboard := NewDashboard(*title)
	count, lastx, lasty := 0, 0, 0

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
					"graph",
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
					"graph", metrics.Get(v).Help(),
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
