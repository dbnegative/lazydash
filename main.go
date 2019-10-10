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
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {

	cfg := Config{}

	app := kingpin.New("lazydash", "generate grafana dashboard json from prometheus metrics data via file or by | pipe")
	app.Flag("file", "Parse metrics from file.").Short('f').StringVar(&cfg.File)
	app.Flag("title", "Dashboard title").Short('t').Default("Demo").StringVar(&cfg.Title)
	app.Flag("stdin", "Read from stdin").Default("true").BoolVar(&cfg.Stdin)
	app.Flag("pretty", "Print pretty indented JSON").Short('p').Default("false").BoolVar(&cfg.Pretty)
	app.Flag("gauges", "Render gauge values as gauge panel types instead of graph").Short('g').Default("false").BoolVar(&cfg.Gauges)
	app.Flag("set-counter-expr", "Set custom meterics query expression for counter type metric").Default("sum(rate(:METRIC: [1m]))").StringVar(&cfg.CounterExprTmpl)
	app.Flag("set-gauge-expr", "Set custom meterics query expression for gauge type metric").Default(":METRIC:").StringVar(&cfg.GaugeExprTmpl)
	app.Flag("set-delimiter", "Set custom meterics delimiter used to insert metric name into expression, only used if a custom expression is set").Default(":METRIC:").StringVar(&cfg.Delimiter)
	app.Flag("set-counter-legend", "Set the default counter panel legend format").Default("Job:[{{job}}]").StringVar(&cfg.CounterLegend)
	app.Flag("set-gauge-legend", "Set the default counter panel legend format").Default("Job:[{{job}}]").StringVar(&cfg.CounterLegend)

	app.Version("0.1.0")
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))

	var (
		data    []byte
		metrics MetricMap
	)

	switch {
	case cfg.File != "": //Read from file
		data, _ = LoadFromFile(cfg.File)
	case cfg.Stdin && cfg.File == "": //Read from STDIN
		data = LoadFromStdin()
	}

	//have we received input?
	if len(data) > 0 {
		metrics = ParseMetrics(data)
	}

	//do we have any parsed metrics?
	if len(metrics.List()) > 0 {
		dashboard := NewDashboard(cfg.Title)
		dashboard.Generate(metrics, cfg)
		dashboard.DumpJSON(cfg.Pretty)
	} else {
		os.Exit(1)
	}

}
