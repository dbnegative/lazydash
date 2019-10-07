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

//Check that Annotations conform
func TestParseAnnotation(t *testing.T) {
	dash := Loadtestdata()

	a := Annotation{
		Datasource: "-- Grafana --",
		Enable:     true,
		Hide:       true,
		IconColor:  "rgba(0, 211, 255, 1)",
		Name:       "Annotations & Alerts",
		Type:       "dashboard",
	}

	if a.Datasource != dash.Annotations.List[0].Datasource {
		t.Errorf("Datasource do not match, expected %v got %v ", a.Datasource, dash.Annotations.List[0].Datasource)
	}
	if a.Enable != dash.Annotations.List[0].Enable {
		t.Errorf("Enabled does not match, expected %v got %v ", a.Enable, dash.Annotations.List[0].Enable)
	}
	if a.Hide != dash.Annotations.List[0].Hide {
		t.Errorf("Hide does not match, expected %v got %v ", a.Hide, dash.Annotations.List[0].Hide)
	}
	if a.IconColor != dash.Annotations.List[0].IconColor {
		t.Errorf("IconColor does not match, expected %v got %v ", a.IconColor, dash.Annotations.List[0].IconColor)
	}
	if a.Name != dash.Annotations.List[0].Name {
		t.Errorf("Name does not match, expected %v got %v ", a.Name, dash.Annotations.List[0].Name)
	}
	if a.Type != dash.Annotations.List[0].Type {
		t.Errorf("Type does not match, expected %v got %v ", a.Type, dash.Annotations.List[0].Type)
	}
}
