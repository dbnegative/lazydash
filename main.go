package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app    = kingpin.New("lazydash", "generate grafana dashboard json from prometheus metrics data via file or by | pipe")
	file   = app.Flag("file", "Parse metrics from file.").Short('f').String()
	title  = app.Flag("title", "Dashboard title").Short('t').String()
	stdin  = app.Flag("stdin", "Read from stdin").Default("true").Bool()
	pretty = app.Flag("pretty", "Print pretty indented JSON").Short('p').Default("false").Bool()
	gauges = app.Flag("gauges", "Render gauge values as gauge panel types instead of graph").Short('g').Default("false").Bool()
)

func main() {

	app.Version("0.0.1")
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))

	var b []byte

	//Read from file
	if *file != "" {
		b, _ = LoadFromFile(*file)
	}

	//Read from STDIN
	if *stdin && *file == "" {
		b = LoadFromStdin()
	}

	var metrics MetricMap

	if len(b) > 0 {
		metrics = ParseMetrics(b)
	}

	if len(metrics.List()) > 0 {
		dashboard := Generate(metrics, *gauges)
		dashboard.DumpJSON(*pretty)
	} else {
		os.Exit(1)
	}

}
