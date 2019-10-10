package main

type Config struct {
	Title             string
	Gauges            bool
	File              string
	Pretty            bool
	Stdin             bool
	CounterExprTmpl   string
	GaugeExprTmpl     string
	HistogramExprTmpl string
	Delimiter         string
	CounterLegend     string
	GaugeLegend       string
}
