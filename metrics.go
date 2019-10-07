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
	"fmt"
	"sort"
)

type Metric struct {
	help   string
	mtype  string
	name   string
	suffix string
	labels map[string]bool
	unit   string
}

type MetricMap struct {
	metrics map[string]*Metric
}

func NewMetric(name string, help string, labels map[string]bool, mtype string, suffix string, unit string) *Metric {
	m := &Metric{}
	m.SetName(name)
	m.SetHelp(help)
	m.SetSuffix(suffix)
	m.SetType(mtype)
	m.SetUnit(unit)

	if len(labels) < 0 || labels == nil {
		m.SetLabels(make(map[string]bool))
	}

	return m
}

func (m *Metric) SetHelp(help string) {
	m.help = help
}

func (m *Metric) Help() string {
	return m.help
}

func (m *Metric) SetType(mtype string) {
	m.mtype = mtype
}

func (m *Metric) Type() string {
	return m.mtype
}

func (m *Metric) SetName(name string) {
	m.name = name
}

func (m *Metric) Name() string {
	return m.name
}

func (m *Metric) SetUnit(unit string) {
	m.unit = unit
}

func (m *Metric) Unit() string {
	return m.unit
}

func (m *Metric) Suffix() string {
	return m.suffix
}

func (m *Metric) SetSuffix(suffix string) {
	m.suffix = suffix
}

func (m *Metric) SetLabels(labels map[string]bool) {
	m.labels = labels
}

func (m *Metric) AddLabel(label string) {
	m.labels[label] = true
}

func (m *Metric) Labels() []string {
	labels := []string{}
	if len(m.labels) > 0 {
		for label := range m.labels {
			labels = append(labels, label)
		}
		return labels
	}

	return nil
}

//Print out all metrics in the map
func (m *MetricMap) Print() {
	for _, v := range m.metrics {
		fmt.Printf("Name: %s Suffix: %s Type: %s  Help: %s Labels %v \n", v.Name(), v.Suffix(), v.Type(), v.Help(), v.Labels())
	}
}

//Set a metric to the map
func (m *MetricMap) Set(name string, metric *Metric) {
	m.metrics[name] = metric
}

//Get a metric out of the map
func (m *MetricMap) Get(name string) *Metric {
	return m.metrics[name]
}

func (m *MetricMap) List() []string {
	list := []string{}
	for k := range m.metrics {
		list = append(list, k)
	}
	sort.Strings(list)

	return list
}
func (m *MetricMap) Init() {
	m.metrics = make(map[string]*Metric)
}
