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
	"encoding/json"
	"fmt"
	"log"
)

//RefreshIntervals for quick reference
var RefreshIntervals = []string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"}

//TimeOptions for quick referenece
var TimeOptions = []string{"5m", "15m", "1h", "3h", "6h", "12h", "24h", "2d", "3d", "4d", "7d", "30d"}

//TimeRange contains a range of Time
type TimeRange struct {
	From string `json:"from"` //: "now-6h",
	To   string `json:"to"`   //: "now"
}

//Dashboard holds all other Grafana sub containers
type Dashboard struct {
	ID            int         `json:"id,omitempty"`
	UID           string      `json:"uid,ommitempty"`
	Title         string      `json:"title,omitempty"` // "New Dashboard",
	Tags          []string    `json:"tags,omitempty"`
	TimeZone      string      `json:"Timezone,omitempty"`
	Editable      bool        `json:"editable,omitempty"`
	HideControls  string      `json:"hideControls,omitempty"`
	GraphToolTip  int         `json:"graphTooltip,omitempty"`
	Panels        []Panel     `json:"Panels,omitempty"`
	Time          TimeRange   `json:"Time,omitempty"`
	TimePicker    TimePicker  `json:"Timepicker,omitempty"`
	Templating    Templating  `json:"templating,omitempty"`
	Annotations   Annotations `json:"annotations,omitempty"`
	Refresh       string      `json:"refresh,omitempty"`
	SchemaVersion int         `json:"schemaVersion,omitempty"`
	Version       int         `json:"version,omitempty"`
	Links         []string    `json:"links,omitempty"`
	Style         string      `json:"style,omitempty"`
}

//TimePicker comtains all attribuyes used to set Dashboard Time options
type TimePicker struct {
	Collapse         bool     `json:"collapse,omitempty"`
	Enable           bool     `json:"enable,omitempty"`
	Notice           bool     `json:"notice,omitempty"`
	Now              bool     `json:"now,omitempty"`
	RefreshIntervals []string `json:"refresh_intervals,omitempty"`
	Status           string   `json:"status,omitempty"`
	TimeOptions      []string `json:"Time_options"`
	Type             string   `json:"type,omitempty"`
	NowDelay         string   `json:"nowDelay,omitempty"`
}

//NewDashboard initialises a new Dashboard
func NewDashboard(title string) *Dashboard {

	return &Dashboard{
		Title:    title,
		TimeZone: "",
		ID:       1,
		Links:    []string{""},
		Version:  1,
		Time: TimeRange{
			From: "now-5m",
			To:   "now",
		},
		TimePicker: TimePicker{
			RefreshIntervals: RefreshIntervals,
		},
		//SchemaVersion: 20,
		//Style:         "dark",
		Tags: []string{""},
		Templating: Templating{
			List: []TemplatingVar{},
		},
		Panels: []Panel{},
	}
}

//AddPanel adds a Panel to the Dashboard
func (d *Dashboard) AddPanel(p Panel) {
	d.Panels = append(d.Panels, p)
}

//AddVariable adds a new templating variable
func (d *Dashboard) AddVariable(t TemplatingVar) {
	d.Templating.List = append(d.Templating.List, t)

}

//AddAnnotation adds a new annotation to the Dashboard
func (d *Dashboard) AddAnnotation() {

}

//DumpJSON dumps the dashboard as JSON
func (d *Dashboard) DumpJSON(pretty bool) {

	var b []byte
	var err error

	if pretty {
		b, err = json.MarshalIndent(*d, "", "  ")
		if err != nil {
			log.Fatalf("could not unmarshal Dashboard json: %v", err)
		}
	} else {
		b, err = json.Marshal(*d)
		if err != nil {
			log.Fatalf("could not unmarshal (noident) Dashboard json: %v", err)
		}
	}

	fmt.Println(string(b))
}
