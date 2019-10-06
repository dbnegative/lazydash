package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//NewDashboard initialises a new dashboard
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

//AddPanel adds a panel to the dashboard
func (d *Dashboard) AddPanel(p Panel) {
	d.Panels = append(d.Panels, p)
}

//AddVariable adds a new templating variable
func (d *Dashboard) AddVariable(t TemplatingVar) {
	d.Templating.List = append(d.Templating.List, t)

}

//AddAnnotation adds a new annotation to the dashboard
func (d *Dashboard) AddAnnotation() {

}

func (d *Dashboard) DumpJSON(pretty bool) {

	var b []byte
	var err error

	if pretty {
		b, err = json.MarshalIndent(*d, "", "  ")
		if err != nil {
			log.Fatalf("could not unmarshal dashboard json: %v", err)
		}
	} else {
		b, err = json.Marshal(*d)
		if err != nil {
			log.Fatalf("could not unmarshal (noident) dashboard json: %v", err)
		}
	}

	fmt.Println(string(b))
}
