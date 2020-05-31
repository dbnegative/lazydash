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
	"io"
	"strings"

	PromLabel "github.com/prometheus/prometheus/pkg/labels"
	PromParse "github.com/prometheus/prometheus/pkg/textparse"
)

//ParseMetrics parses metrics as a []byte and returns a copy of MetricsMap
func ParseMetrics(metrics []byte) MetricMap {

	p := PromParse.NewPromParser(metrics)
	mm := MetricMap{}
	mm.New()

	for {
		et, err := p.Next()
		if err == io.EOF {
			break
		}

		//May be parsed out of order
		switch et {

		case PromParse.EntryHelp:
			m, h := p.Help()
			mm.Set(string(m), NewMetric(string(m), string(h), nil, "", "", "short"))

		case PromParse.EntryType:
			m, typ := p.Type()
			mm.Get(string(m)).SetType(string(typ))

		case PromParse.EntrySeries:
			labels := &PromLabel.Labels{}
			p.Metric(labels)

			labelmap := labels.Map()
			name := labelmap["__name__"]

			//Unify metrics key for simple access
			if strings.Contains(name, "_bucket") {
				name = strings.TrimSuffix(name, "_bucket")
				mm.Get(name).SetSuffix("_bucket")
			} else if strings.Contains(name, "_sum") {
				name = strings.TrimSuffix(name, "_sum")
				mm.Get(name).SetSuffix("_sum")
			} else if strings.Contains(name, "_count") {
				name = strings.TrimSuffix(name, "_count")
				mm.Get(name).SetSuffix("_count")
			}

			//Try guess the metrics unit type
			if strings.Contains(name, "_seconds") {
				mm.Get(name).SetUnit("s")
			} else if strings.Contains(name, "_milliseconds") {
				mm.Get(name).SetUnit("ms")
			} else if strings.Contains(name, "_bytes") {
				mm.Get(name).SetUnit("decbytes")
			} else {
				mm.Get(name).SetUnit("short")
			}

			//Add all labels
			for k := range labelmap {
				if k != "__name__" && k != "" {
					mm.Get(name).AddLabel(k)
				}
			}

		case PromParse.EntryComment:
			//fmt.Println(string(p.Comment()))
		}
	}

	return mm
}
