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
