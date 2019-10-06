package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	file   = kingpin.Flag("file", "Parse metrics from file.").String()
	title  = kingpin.Flag("title", "Dashboard title").String()
	stdin  = kingpin.Flag("stdin", "Read from stdin").Default("true").Bool()
	pretty = kingpin.Flag("pretty", "Print pretty indented JSON").Default("false").Bool()
)

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()

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
		dashboard := Generate(metrics)
		dashboard.DumpJSON(*pretty)
	}

}
