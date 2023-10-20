package main

// Config contains configuration options
type Config struct {
	Title              string
	Gauges             bool
	File               string
	Pretty             bool
	Stdin              bool
	CounterExprTmpl    string
	GaugeExprTmpl      string
	SummaryExprTmpl    string
	Delimiter          string
	CounterLegend      string
	GaugeLegend        string
	SummaryLegend      string
	Table              bool
	URL                string
	Token              string
	GrafanaHost        string
	InsecureSkipVerify bool
}
