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

import "testing"

func TestDashboard(t *testing.T) {
	dash := LoadTestData()

	d := &Dashboard{
		Title:    "demo",
		TimeZone: "",
		ID:       1,
		Links:    []string{""},
		UID:      "QIbx6hhZz",
		Version:  4,
		Time: TimeRange{
			From: "now-5m",
			To:   "now",
		},
		TimePicker: TimePicker{
			RefreshIntervals: RefreshIntervals,
		},
		SchemaVersion: 20,
		Style:         "dark",
		Tags:          []string{""},
		Templating: Templating{
			List: []TemplatingVar{},
		},
	}

	//fmt.Printf("%v", d)

	if d.ID != dash.ID {
		t.Errorf("ID does not match expected %v got %v ", d.ID, dash.ID)
	}
	if d.Title != dash.Title {
		t.Errorf("Title does not match expected %v got %v ", d.Title, dash.Title)
	}

	if d.SchemaVersion != dash.SchemaVersion {
		t.Errorf("Schema Versions do not match, expected %v got %v ", d.SchemaVersion, d.SchemaVersion)
	}
	if d.Style != dash.Style {
		t.Errorf("Styles do not match, expected %v got %v ", d.Panels[0].Type, dash.Panels[0].Type)
	}

}
